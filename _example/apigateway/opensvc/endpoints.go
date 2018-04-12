package opensvc

import (
	"context"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/payload"
	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/distributor"
)

type Endpoints struct {
	SnsEndpoint moss.Endpoint
}

func MakeServerEndpoints() Endpoints {
	return Endpoints{
		SnsEndpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return distributor.WatcherInvoking(_example.ServiceNameSns, ctx, request.(*payload.MossPacket))
		},
	}
}
