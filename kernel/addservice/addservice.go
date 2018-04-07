package addservice

import (
	"errors"

	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/metrics"

	"golang.org/x/net/context"
)

type Service interface {
	Invoking(ctx context.Context, input *payload.MossPacket) (*payload.MossPacket, error)
}

func NewService(counter metrics.Counter, histogram metrics.Histogram) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware()(svc)
		svc = InstrumentingMiddleware(counter, histogram, histogram)(svc)
	}
	return svc
}

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (s basicService) Invoking(_ context.Context, input *payload.MossPacket) (*payload.MossPacket, error) {
	if input.ServiceCode < 1 {
		return nil, errors.New("no protocol code")
	}
	return input, nil
}
