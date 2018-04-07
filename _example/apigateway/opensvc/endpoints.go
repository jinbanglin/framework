package opensvc

import (
	"context"

	"moss"
	"moss/kernel/payload"
)

type Endpoints struct {
	RegisterEndpoint moss.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		RegisterEndpoint: MakeRegisterEndpoint(s),
	}
}

func MakeRegisterEndpoint(s Service) moss.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.Register(ctx, request.(*payload.MossPacket))
	}
}
