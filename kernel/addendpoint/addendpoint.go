package addendpoint

import (
	"time"

	"context"
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/kernel/addservice"
	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/metrics"
	"github.com/jinbanglin/moss/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"github.com/jinbanglin/moss/circuitbreaker"
	"github.com/jinbanglin/moss/log"
	"golang.org/x/time/rate"
	"github.com/jinbanglin/moss/limiter"
)

type Set struct {
	InvokeEndpoint moss.Endpoint
}

func NewEndpoint(svc addservice.Service, duration metrics.Histogram, trace opentracing.Tracer) (Set) {
	invokingEndpoint := MakeInvokeEndpoint(svc)
	invokingEndpoint = limiter.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1024))(invokingEndpoint)
	invokingEndpoint = limiter.NewDelayingLimiter(rate.NewLimiter(rate.Every(time.Second*10), 100))(invokingEndpoint)
	invokingEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "GRPC_CLIENT",
		MaxRequests: 10,
		Interval:    10,
		Timeout:     10,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Infof("New |gobreaker |state |from=%s |to=%s", from.String(), to.String())
		},
	}))(invokingEndpoint)
	invokingEndpoint = tracing.TraceServer(trace, "Invoking")(invokingEndpoint)
	invokingEndpoint = LoggingMiddleware()(invokingEndpoint)
	invokingEndpoint = InstrumentingMiddleware(duration.With("method", "Invoking"))(invokingEndpoint)
	return Set{
		InvokeEndpoint: invokingEndpoint,
	}
}

func (s Set) Invoking(ctx context.Context, a *payload.MossPacket) (*payload.MossPacket, error) {
	resp, err := s.InvokeEndpoint(ctx, a)
	if err != nil {
		return nil, err
	}
	return resp.(*payload.MossPacket), err
}

func MakeInvokeEndpoint(s addservice.Service) moss.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.Invoking(ctx, request.(*payload.MossPacket))
	}
}
