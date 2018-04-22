package opensvc

import (
	"context"

	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/ipc"
	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/payload"
)

type Endpoints struct {
	SnsEndpoint endpoint.Endpoint
}

func MakeServerEndpoints() Endpoints {
	return Endpoints{
		SnsEndpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return ipc.WatcherInvoking(_example.ServiceNameSns, ctx, request.(*payload.MossPacket))
		},
	}
}
