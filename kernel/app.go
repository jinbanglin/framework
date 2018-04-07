package kernel

import (
	"net"
	"net/http"
	"sync"
	"syscall"
	"time"

	"github.com/jinbanglin/moss/discovery/etcdv3"
	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/log"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var AppServer = &appServer{}

type appServer struct {
	*ConfigManager
	ServiceName
	wait sync.WaitGroup
}

func (a *appServer) SetupConfig(name ServiceName, f ...func()) {
	AppServer = &appServer{ConfigManager: &ConfigManager{
		EtcdEndPoints: &EtcdV3{},
	}, ServiceName: name}
	AppServer.setupConfig(name, f...)
}

func (a *appServer) getServerAddr(connectionType ConnectionType) string {
	for _, v := range a.Connections {
		if v.ConnType == connectionType {
			return v.Addr
		}
	}
	panic("no config connection type : " + connectionType)
}

func (a *appServer) GrpcServerStart() {
	addr := a.getServerAddr(CONNECTION_TYPE_GRPC)
	a.registerEtcdV3(addr)
	if len(a.Watchers) > 0 {
		WatcherInstance().Watch(a.Watchers)
	}
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Info("GrpcServer |start at:", addr)
	baseServer := grpc.NewServer()
	payload.RegisterInvokingServer(baseServer, GrpcInstance().gScheduler)
	reflection.Register(baseServer)
	if baseServer.Serve(grpcListener) != nil {
		panic("GrpcServer Work error")
	}
}

func (a *appServer) HttpGatewayStart(r *mux.Router) {
	WatcherInstance().Watch(a.Watchers)
	gateway := NewHTTPGateway()
	gateway.loadBalancing(WatcherInstance())
	log.Info("HttpGateway start at:", a.getServerAddr(CONNECTION_TYPE_HTTP))
	http.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), gateway.MakeHttpHandle(r, a.EtcdEndPoints.ServerId))
}

func (a *appServer) Run() {
	a.wait.Wait()
}

func (a *appServer) registerEtcdV3(serverAddr string) {
	NewEtcdV3Client().Register(etcdv3.Service{
		Key:   "/" + a.ServiceName + "/" + a.ConfigManager.EtcdEndPoints.ServerId,
		Value: serverAddr,
		TTL:   etcdv3.NewTTLOption(0, 0),
	})
	log.Infof("register etcd key=%s", "/"+a.ServiceName+"/"+a.ConfigManager.EtcdEndPoints.ServerId)
	log.Info("register etcd value", serverAddr)
}

func (a *appServer) Stop(timeout time.Duration, f ...func()) {
	a.wait.Add(1)
	process := func() {
		defer a.wait.Done()
		time.AfterFunc(timeout, func() {
			log.Info("server stop", "server_id", a.EtcdEndPoints.ServerId, "server_name", a.ServiceName)
		})
		for _, v := range f {
			v()
		}
		//todo 系统资源释放
	}
	RegisterContinueSignal(syscall.SIGTERM, process)
	RegisterContinueSignal(syscall.SIGINT, process)
}
