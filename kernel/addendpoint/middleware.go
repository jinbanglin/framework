package addendpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/metrics"
)

func InstrumentingMiddleware(duration metrics.Histogram) moss.Middleware {
	return func(next moss.Endpoint) moss.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("error", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func LoggingMiddleware() moss.Middleware {
	return func(next moss.Endpoint) moss.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.Errorf("transport_error |err=%v |took=%d", err, time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
