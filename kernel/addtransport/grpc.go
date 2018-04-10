package addtransport

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/circuitbreaker"
	"github.com/jinbanglin/moss/kernel/addendpoint"
	"github.com/jinbanglin/moss/kernel/addservice"
	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/limiter"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/tracing"
	transportgrpc "github.com/jinbanglin/moss/transport/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	context2 "golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

const CONTEXT_KEY_SERVICE_CODE = "service_code"

type SchedulerHandler struct {
	RequestType reflect.Type
	handler     transportgrpc.Handler
}

type Scheduler struct {
	scheduler map[uint32]*SchedulerHandler
	tracer    opentracing.Tracer
	ops       []transportgrpc.ServerOption
	metrics   *Metrics
	sync.RWMutex
}

func (s *Scheduler) Invoking(ctx context2.Context, request *payload.MossPacket) (response *payload.MossPacket, err error) {
	response = &payload.MossPacket{
		Message: &payload.Message{Code: 50001, Msg: "invoking error"},
	}
	ctx = context.WithValue(ctx, CONTEXT_KEY_SERVICE_CODE, request.ServiceCode)
	schedulerHandler, err := gScheduler.GetHandler(request.ServiceCode)
	if err != nil {
		log.Error(err)
		return response, err
	}
	req := reflect.New(schedulerHandler.RequestType).Interface().(proto.Message)
	if err = GetCodecerByServiceCode(request.ServiceCode).Unmarshal(request.Payload, req); err != nil {
		log.Error(err)
		return response, err
	}
	_, res, err := schedulerHandler.handler.ServeGRPC(ctx, request)
	if err != nil {
		log.Errorf("Invoking |res=%v |err=%v", res, err)
		return response, err
	}
	loader, err := GetCodecerByServiceCode(request.ServiceCode).Marshal(res.(proto.Message))
	if err != nil {
		return response, err
	}
	response.Payload = loader
	response.Message = &payload.Message{Code: 20000, Msg: "SUCCESS"}
	return response, nil
}

var gScheduler *Scheduler

func NewGrpcServer(tracer opentracing.Tracer, ops []transportgrpc.ServerOption) *Scheduler {
	gScheduler = &Scheduler{
		scheduler: make(map[uint32]*SchedulerHandler),
		tracer:    tracer,
		ops:       ops,
		metrics:   NewMetrics(),
	}
	return gScheduler
}

func NewHandler(endpoint moss.Endpoint, tracer opentracing.Tracer, ops []transportgrpc.ServerOption) transportgrpc.Handler {
	service := addservice.NewService(gScheduler.metrics.Counters, gScheduler.metrics.SummaryError)
	endpoint = addendpoint.NewEndpoint(service, gScheduler.metrics.SummarySuccess, opentracing.GlobalTracer()).InvokeEndpoint
	return transportgrpc.NewServer(
		endpoint,
		decodeRequest,
		encodeResponse,
		append(ops, transportgrpc.ServerBefore(tracing.GRPCToContext(tracer, "Invoking")))...,
	)
}

func (s *Scheduler) RegisterHandler(serviceCode uint32, request proto.Message, endpoint moss.Endpoint) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.scheduler[serviceCode]; ok {
		panic("handler is already register")
	}
	s.scheduler[serviceCode] = &SchedulerHandler{RequestType: reflect.TypeOf(request).Elem(), handler: NewHandler(
		endpoint,
		s.tracer,
		s.ops,
	)}
}

func (s *Scheduler) GetHandler(serviceCode uint32) (handler *SchedulerHandler, err error) {
	s.RLock()
	var ok bool
	handler, ok = s.scheduler[serviceCode]
	if !ok {
		err = errors.New(fmt.Sprintf("cannot find service code: %d", serviceCode))
		s.RUnlock()
		return
	}
	s.RUnlock()
	return
}

func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
	}
	return request, nil
}

func encodeResponse(ctx context.Context, request interface{}) (res interface{}, err error) {
	serviceCode := ctx.Value(CONTEXT_KEY_SERVICE_CODE).(uint32)
	response := &payload.MossPacket{ServiceCode: serviceCode}
	response.Payload, err = GetCodecerByServiceCode(serviceCode).Marshal(request.(proto.Message))
	return response, err
}

func NewGRPCClient(conn *grpc.ClientConn, tracer opentracing.Tracer) addservice.Service {
	limiterErr := limiter.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1024))
	limiterDelay := limiter.NewDelayingLimiter(rate.NewLimiter(rate.Every(time.Second), 1024))
	var invokeEndpoint moss.Endpoint
	{
		invokeEndpoint = transportgrpc.NewClient(
			conn,
			"payload.Invoking",
			"Invoking",
			encodeGRPCInvokeRequest,
			decodeGRPCInvokeResponse,
			payload.MossPacket{},
			transportgrpc.ClientBefore(tracing.ContextToGRPC(tracer)),
		).Endpoint()
		invokeEndpoint = tracing.TraceClient(tracer, "Invoking")(invokeEndpoint)
		invokeEndpoint = limiterErr(invokeEndpoint)
		invokeEndpoint = limiterDelay(invokeEndpoint)
		invokeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "GRPC_CLIENT",
			MaxRequests: 10,
			Interval:    10,
			Timeout:     10,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failureRatio >= 0.6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				log.Infof("NewGRPCClient |gobreaker |state |from=%s |to=%s", from.String(), to.String())
			},
		}))(invokeEndpoint)
	}
	return addendpoint.Set{InvokeEndpoint: invokeEndpoint}
}

func decodeGRPCInvokeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeGRPCInvokeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}
