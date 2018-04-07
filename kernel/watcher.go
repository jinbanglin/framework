package kernel

import (
	"context"
	"io"
	"time"

	"moss"
	"moss/discovery"
	"moss/discovery/etcdv3"
	"moss/discovery/lb"
	"moss/kernel/addendpoint"
	"moss/kernel/addtransport"
	"moss/kernel/payload"
	"moss/log"

	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
)

var gWatcher *Watcher

type Watcher struct {
	watchers map[ServiceName]*WatcherEndpoint
}

type WatcherEndpoint struct {
	etcdClient    etcdv3.Client
	etcdInstancer *etcdv3.Instancer
	factory       discovery.Factory
	sdEndPointer  discovery.Endpointer
	lbRoundRobin  lb.Balancer
	retry         moss.Endpoint
	endpoint      addendpoint.Set
}

func WatcherInstance() *Watcher {
	if gWatcher == nil {
		gWatcher = &Watcher{watchers: make(map[ServiceName]*WatcherEndpoint)}
	}
	return gWatcher
}

func (w *Watcher) Watch(config []*Watch) {
	for _, v := range config {
		log.Info("watch service name", v.ServiceName)
		w.watchers[v.ServiceName], _ = newWatchEndpoint(v.ServiceName)
	}
}

func (r *WatcherEndpoint) InvokeEndpoint() moss.Endpoint {
	return r.endpoint.InvokeEndpoint
}

func newWatchEndpoint(serviceName ServiceName) (watcherEndPoint *WatcherEndpoint, err error) {
	watcherEndPoint = &WatcherEndpoint{}
	watcherEndPoint.etcdClient = NewEtcdV3Client()
	watcherEndPoint.etcdInstancer, err = etcdv3.NewInstancer(
		watcherEndPoint.etcdClient,
		"/"+serviceName,
	)
	if err != nil {
		panic(err)
	}
	watcherEndPoint.factory = func(instance string) (moss.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(1*time.Second),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(64<<20),
				grpc.MaxCallSendMsgSize(64<<20),
			))
		if err != nil {
			return nil, nil, err
		}
		return addendpoint.MakeInvokeEndpoint(
			addtransport.NewGRPCClient(conn, opentracing.GlobalTracer())), conn, nil
	}
	watcherEndPoint.sdEndPointer = discovery.NewEndpointer(
		watcherEndPoint.etcdInstancer,
		watcherEndPoint.factory,
	)
	watcherEndPoint.lbRoundRobin = lb.NewRoundRobin(watcherEndPoint.sdEndPointer)
	watcherEndPoint.retry = lb.Retry(3, time.Second*1, watcherEndPoint.lbRoundRobin)
	watcherEndPoint.endpoint.InvokeEndpoint = watcherEndPoint.retry
	return
}

func WatcherInvoking(serviceName ServiceName, ctx context.Context, request *payload.MossPacket) (response *payload.MossPacket, error error) {
	return WatcherInstance().watchers[serviceName].endpoint.Invoking(ctx, request)
}
