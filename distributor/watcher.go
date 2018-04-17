package distributor

import (
	"context"
	"io"
	"time"

	"github.com/jinbanglin/moss/discovery"
	"github.com/jinbanglin/moss/discovery/etcdv3"
	"github.com/jinbanglin/moss/discovery/lb"
	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/payload"
	"google.golang.org/grpc"
)

var gWatcher *Watcher

type Watcher struct {
	watchers map[string]*WatcherEndpoint
}

type WatcherEndpoint struct {
	etcdInstancer *etcdv3.Instancer
	factory       discovery.Factory
	sdEndPointer  discovery.Endpointer
	lbRoundRobin  lb.Balancer
	retry         endpoint.Endpoint
	endpoint      endpoint.Endpoint
}

func WatcherInstance() *Watcher {
	if gWatcher == nil {
		gWatcher = &Watcher{watchers: make(map[string]*WatcherEndpoint)}
	}
	return gWatcher
}

func (w *Watcher) Watch(services, etcdAddress []string) {
	for _, v := range services {
		log.Info("MOSS |watch ", v)
		w.watchers[v] = newWatchEndpoint(v, etcdAddress)
	}
}

func newWatchEndpoint(serviceName string, etcdAddress []string) (watcher *WatcherEndpoint) {
	watcher = &WatcherEndpoint{}
	watcher.etcdInstancer = etcdv3.NewInstancer(etcdv3.DefaultEtcdV3Client(etcdAddress), "/"+string(serviceName))
	watcher.factory = func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if conn, err := grpc.Dial(instance, grpc.WithInsecure()); err != nil {
			log.Error("MOSS |err=", err)
			return nil, nil, err
		} else {
			return NewGRPCClient(conn), conn, nil
		}
	}
	watcher.sdEndPointer = discovery.NewEndpointer(watcher.etcdInstancer, watcher.factory)
	watcher.lbRoundRobin = lb.NewRoundRobin(watcher.sdEndPointer)
	watcher.retry = lb.Retry(3, time.Second*10, watcher.lbRoundRobin)
	watcher.endpoint = watcher.retry
	return
}

func WatcherInvoking(serviceName string, ctx context.Context, request *payload.MossPacket) (*payload.MossPacket, error) {
	response, err := WatcherInstance().watchers[serviceName].endpoint(ctx, request)
	if response == nil {
		response = &payload.MossPacket{MossMessage: payload.StatusText(payload.StatusBadRequest)}
	}
	return response.(*payload.MossPacket), err
}
