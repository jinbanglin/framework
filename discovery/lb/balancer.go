package lb

import (
	"errors"

	"github.com/jinbanglin/moss/endpoint"
)

// Balancer yields endpoints according to some heuristic.
type Balancer interface {
	Endpoint() (endpoint.Endpoint, error)
}

// ErrNoEndpoints is returned when no qualifying endpoints are available.
var ErrNoEndpoints = errors.New("no endpoints available")
