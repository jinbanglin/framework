package opensvc

import (
	"golang.org/x/net/context"
	"time"

	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/log"
)

type Middleware func(Service) Service

func LoggingMiddleware() Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	next   Service
}

func (l *loggingMiddleware)Register(ctx context.Context, request *payload.MossPacket) (*payload.MossPacket, error) {
	defer func(begin time.Time) {
		log.Info("opensvc", "Login", "took", time.Since(begin))
	}(time.Now())
	return l.next.Register(ctx, request)
}

