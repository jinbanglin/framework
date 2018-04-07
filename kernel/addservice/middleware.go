package addservice

import (
	"context"
	"fmt"
	"time"

	"moss/kernel/payload"
	"moss/log"
	"moss/metrics"
)

type Middleware func(Service) Service

func LoggingMiddleware() Middleware {
	return func(next Service) Service {
		return loggingMiddleware{next}
	}
}

type loggingMiddleware struct {
	next   Service
}

func (mw loggingMiddleware) Invoking(ctx context.Context, input *payload.MossPacket) (ouput *payload.MossPacket, err error) {
	defer func(begin time.Time) {
		log.Info(
			"method", "Invoking",
			"input", input,
			"output", input,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return input, nil
}

func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency, countResult metrics.Histogram) Middleware {
	return func(next Service) Service {
		return instrumentingMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			countResult:    countResult,
			next:           next,
		}
	}
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           Service
}

func (mw instrumentingMiddleware) Invoking(ctx context.Context, input *payload.MossPacket) (output *payload.MossPacket, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Invoking", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.countResult.With(lvs...).Observe(float64(input.ServiceCode))
	}(time.Now())
	return input, nil
}
