package endpoint

import (
	"context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type EndpointAdapter func(Endpoint) Endpoint



