package dnssrv

import (
	"fmt"
	"net"
	"time"

	"github.com/jinbanglin/moss/discovery"
	"github.com/jinbanglin/moss/log"
)

// Instancer yields instances from the named DNS SRV record. The name is
// resolved on a fixed schedule. Priorities and weights are ignored.
type Instancer struct {
	cache *discovery.Cache
	name  string
	quit  chan struct{}
}

// NewInstancer returns a DNS SRV instancer.
func NewInstancer(
	name string,
	ttl time.Duration,
) *Instancer {
	return NewInstancerDetailed(name, time.NewTicker(ttl), net.LookupSRV, logger)
}

// NewInstancerDetailed is the same as NewInstancer, but allows users to
// provide an explicit lookup refresh ticker instead of a TTL, and specify the
// lookup function instead of using net.LookupSRV.
func NewInstancerDetailed(
	name string,
	refresh *time.Ticker,
	lookup Lookup,
) *Instancer {
	p := &Instancer{
		cache: discovery.NewCache(),
		name:  name,
		quit:  make(chan struct{}),
	}

	instances, err := p.resolve(lookup)
	if err == nil {
		log.Infof("MOSS |dnssrv |name=%s |instances=%v", name, len(instances))
	} else {
		log.Error("MOSS |", name, err)
	}
	p.cache.Update(discovery.Event{Instances: instances, Err: err})

	go p.loop(refresh, lookup)
	return p
}

// Stop terminates the Instancer.
func (s *Instancer) Stop() {
	close(s.quit)
}

func (s *Instancer) loop(t *time.Ticker, lookup Lookup) {
	defer t.Stop()
	for {
		select {
		case <-t.C:
			instances, err := s.resolve(lookup)
			if err != nil {
				s.cache.Update(discovery.Event{Err: err})
				continue // don't replace potentially-good with bad
			}
			s.cache.Update(discovery.Event{Instances: instances})

		case <-s.quit:
			return
		}
	}
}

func (s *Instancer) resolve(lookup Lookup) ([]string, error) {
	_, addrs, err := lookup("", "", s.name)
	if err != nil {
		return nil, err
	}
	instances := make([]string, len(addrs))
	for i, addr := range addrs {
		instances[i] = net.JoinHostPort(addr.Target, fmt.Sprint(addr.Port))
	}
	return instances, nil
}

// Register implements Instancer.
func (s *Instancer) Register(ch chan<- discovery.Event) {
	s.cache.Register(ch)
}

// Deregister implements Instancer.
func (s *Instancer) Deregister(ch chan<- discovery.Event) {
	s.cache.Deregister(ch)
}
