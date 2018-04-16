package etcdv3

import (
	"context"
	"time"

	"github.com/jinbanglin/moss/discovery"
	"github.com/jinbanglin/moss/log"
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
func NewInstancer(c Client, prefix string) (*Instancer) {
	s := &Instancer{
		client: c,
		prefix: prefix,
		cache:  discovery.NewCache(),
		quitc:  make(chan struct{}),
	}

	instances, err := s.client.GetEntries(s.prefix)
	if err != nil {
		panic(err)
	}
	log.Infof("✨MOSS✨ |prefix=%s |instances=%d", s.prefix, len(instances))
	s.cache.Update(discovery.Event{Instances: instances, Err: err})
	go s.loop()
	return s
}

func (s *Instancer) loop() {
	ch := make(chan struct{})
	go s.client.WatchPrefix(s.prefix, ch)
	for {
		select {
		case <-ch:
			instances, err := s.client.GetEntries(s.prefix)
			if err != nil {
				log.Errorf("✨MOSS✨ |Etcdv3 |loop |msg=%s |err=%v", "failed to retrieve entries", err)
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

func DefaultEtcdV3Client(address []string) Client {
	client, err := NewClient(context.Background(), address, ClientOptions{
		CACert:        "",
		Cert:          "",
		Key:           "",
		Username:      "",
		Password:      "",
		DialTimeout:   time.Second * 10,
		DialKeepAlive: time.Second * 3,
	})
	if err != nil {
		panic(err)
	}
	return client
}
