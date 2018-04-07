package etcdv3

import (
	"moss/discovery"
	"moss/log"
)

// Instancer yields instances stored in a certain etcd keyspace. Any kind of
// change in that keyspace is watched and will update the Instancer's Instancers.
type Instancer struct {
	cache  *discovery.Cache
	client Client
	prefix string
	quitc  chan struct{}
}

// NewInstancer returns an etcd instancer. It will start watching the given
// prefix for changes, and update the subscribers.
func NewInstancer(c Client, prefix string) (*Instancer, error) {
	s := &Instancer{
		client: c,
		prefix: prefix,
		cache:  discovery.NewCache(),
		quitc:  make(chan struct{}),
	}

	instances, err := s.client.GetEntries(s.prefix)
	if err == nil {
		log.Infof("NewInstancer |prefix=%s |instances=%d", s.prefix, len(instances))
	} else {
		log.Infof("NewInstancer |prefix=%s |err=%v", s.prefix, err)
	}
	s.cache.Update(discovery.Event{Instances: instances, Err: err})

	go s.loop()
	return s, nil
}

func (s *Instancer) loop() {
	ch := make(chan struct{})
	go s.client.WatchPrefix(s.prefix, ch)

	for {
		select {
		case <-ch:
			instances, err := s.client.GetEntries(s.prefix)
			if err != nil {
				log.Errorf("Etcdv3 |loop |msg=%s |err=%v","failed to retrieve entries",err)
				s.cache.Update(discovery.Event{Err: err})
				continue
			}
			s.cache.Update(discovery.Event{Instances: instances})

		case <-s.quitc:
			return
		}
	}
}

// Stop terminates the Instancer.
func (s *Instancer) Stop() {
	close(s.quitc)
}

// Register implements Instancer.
func (s *Instancer) Register(ch chan<- discovery.Event) {
	s.cache.Register(ch)
}

// Deregister implements Instancer.
func (s *Instancer) Deregister(ch chan<- discovery.Event) {
	s.cache.Deregister(ch)
}
