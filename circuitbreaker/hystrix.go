package circuitbreaker

import (
	"context"

	"moss"

	"github.com/afex/hystrix-go/hystrix"
)

// Hystrix returns an endpoint.Middleware that implements the circuit
// breaker pattern using the afex/hystrix-go package.
//
// When using this circuit breaker, please configure your commands separately.
//
// See https://godoc.org/github.com/afex/hystrix-go/hystrix for more
// information.
func Hystrix(commandName string) moss.Middleware {
	return func(next moss.Endpoint) moss.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err := hystrix.Do(commandName, func() (err error) {
				response, err = next(ctx, request)
				return err
			}, nil); err != nil {
				return nil, err
			}
			return response, nil
		}
	}
}
