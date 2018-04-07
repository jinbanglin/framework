package moss

import (
	"golang.org/x/net/context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Middleware func(Endpoint) Endpoint
