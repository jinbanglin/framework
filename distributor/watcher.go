package distributor

import (
	"context"
	"io"
	"time"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/discovery"
	"github.com/jinbanglin/moss/discovery/etcdv3"
	"github.com/jinbanglin/moss/discovery/lb"
	"github.com/jinbanglin/moss/log"
	"google.golang.org/grpc"
	distributorgrpc "github.com/jinbanglin/moss/distributor/grpc"
	"github.com/jinbanglin/moss/payload"
)

var gWatcher *Watcher

type Watcher struct {
	watchers map[moss.ServiceName]*WatcherEndpoint
}

type WatcherEndpoint struct {
	etcdInstancer *etcdv3.Instancer
	factory       discovery.Factory
	sdEndPointer  discovery.Endpointer
	lbRoundRobin  lb.Balancer
	retry         moss.Endpoint
	endpoint      moss.Endpoint
}

func WatcherInstance() *Watcher {
	if gWatcher == nil {
		gWatcher = &Watcher{watchers: make(map[moss.ServiceName]*WatcherEndpoint)}
	}
	return gWatcher
}

func (w *Watcher) Watch(config []*Watch) {
	for _, v := range config {
		log.Info("watch service name", v.ServiceName)
		w.watchers[v.ServiceName] = newWatchEndpoint(v.ServiceName)
	}
}

func newWatchEndpoint(serviceName moss.ServiceName) (watcher *WatcherEndpoint) {
	watcher = &WatcherEndpoint{}
	watcher.etcdInstancer = etcdv3.NewInstancer(defaultEtcdV3Client(), "/"+string(serviceName))
	watcher.factory = func(instance string) (moss.Endpoint, io.Closer, error) {
		if conn, err := grpc.Dial(instance, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(1*time.Second),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(64<<20),
				grpc.MaxCallSendMsgSize(64<<20),
			)); err != nil {
			log.Error(err)
			return nil, nil, err
		} else {
			return distributorgrpc.NewGRPCClient(conn), conn, nil
		}
	}
	watcher.sdEndPointer = discovery.NewEndpointer(watcher.etcdInstancer, watcher.factory)
	watcher.lbRoundRobin = lb.NewRoundRobin(watcher.sdEndPointer)
	watcher.retry = lb.Retry(1, time.Millisecond*100, watcher.lbRoundRobin)
	watcher.endpoint = watcher.retry
	return
}

func WatcherInvoking(serviceName moss.ServiceName, ctx context.Context, request *payload.MossPacket) (response *payload.MossPacket, error error) {
	if ret, err := WatcherInstance().watchers[serviceName].endpoint(ctx, request); err != nil {
		return nil, err
	} else {
		return ret.(*payload.MossPacket), nil
	}
}
