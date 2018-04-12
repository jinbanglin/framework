package distributor

import (
	"reflect"
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/log"
	transportgrpc "github.com/jinbanglin/moss/transport/grpc"

	"github.com/golang/protobuf/proto"
	context2 "golang.org/x/net/context"
	"errors"
)

type SchedulerHandler struct {
	RequestType reflect.Type
	handler     transportgrpc.Handler
}

type GPRCInvoking struct {
	scheduler map[uint32]*SchedulerHandler
}

func (s *GPRCInvoking) Invoking(ctx context2.Context, request *payload.MossPacket) (response *payload.MossPacket, err error) {
	response = &payload.MossPacket{
		Message: &payload.Message{Code: 50001, Msg: "invoking error"},
	}
	schedulerHandler, err := s.GetHandler(request.ServiceCode)
	if err != nil {
		log.Error(err)
		return response, err
	}
	req := reflect.New(schedulerHandler.RequestType).Interface().(proto.Message)
	if err = moss.GetCodec(request.ServiceCode).Unmarshal(request.Payload, req); err != nil {
		log.Errorf("Invoking |request=%v |err=%v", request, err)
		return response, err
	}
	ctx = context2.WithValue(ctx, "user_id", request.UserId)
	ctx = context2.WithValue(ctx, "client_ip", request.ClientIp)
	res, err := schedulerHandler.handler.ServeGRPC(ctx, req)
	if err != nil {
		log.Errorf("Invoking |res=%v |err=%v", res, err)
		return response, err
	}
	loader, err := moss.GetCodec(request.ServiceCode).Marshal(res.(proto.Message))
	if err != nil {
		log.Errorf("Invoking |res=%v |err=%v", res, err)
		return response, err
	}
	response.Payload = loader
	response.Message = &payload.Message{Code: 20000, Msg: "SUCCESS"}
	response.ServiceCode = request.ServiceCode
	return response, nil
}

func AddEndpoint(endpoint moss.Endpoint) transportgrpc.Handler { return transportgrpc.NewServer(endpoint) }

func (s *GPRCInvoking) RegisterHandler(serviceCode uint32, request proto.Message, endpoint moss.Endpoint) {
	if _, ok := s.scheduler[serviceCode]; ok {
		panic("handler is already register")
	}
	s.scheduler[serviceCode] = &SchedulerHandler{RequestType: reflect.TypeOf(request).Elem(), handler: AddEndpoint(
		endpoint,
	)}
}

func (s *GPRCInvoking) GetHandler(serviceCode uint32) (handler *SchedulerHandler, err error) {
	var ok bool
	if handler, ok = s.scheduler[serviceCode]; !ok {
		log.Error(serviceCode)
		return nil, errors.New("no service")
	}
	return
}
