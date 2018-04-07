package lb

import (
	"errors"

	"github.com/jinbanglin/moss"
)

// Balancer yields endpoints according to some heuristic.
type Balancer interface {
	Endpoint() (moss.Endpoint, error)
}

// ErrNoEndpoints is returned when no qualifying endpoints are available.
var ErrNoEndpoints = errors.New("no endpoints available")
