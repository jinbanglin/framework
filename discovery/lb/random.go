package lb

import (
	"math/rand"

	"moss"
	"github.com/jinbanglin/moss/discovery"
)

// NewRandom returns a load balancer that selects services randomly.
func NewRandom(s discovery.Endpointer, seed int64) Balancer {
	return &random{
		s: s,
		r: rand.New(rand.NewSource(seed)),
	}
}

type random struct {
	s discovery.Endpointer
	r *rand.Rand
}

func (r *random) Endpoint() (moss.Endpoint, error) {
	endpoints, err := r.s.Endpoints()
	if err != nil {
		return nil, err
	}
	if len(endpoints) <= 0 {
		return nil, ErrNoEndpoints
	}
	return endpoints[r.r.Intn(len(endpoints))], nil
}
