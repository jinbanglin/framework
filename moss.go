// moss Endpoint is adapter to all handler middleware filter
package moss

import (
	"context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Middleware func(Endpoint) Endpoint
