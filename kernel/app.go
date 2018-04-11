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
	"github.com/golang/crypto/acme/autocert"
	"crypto/tls"
	"fmt"
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

func (a *appServer) HttpsGatewayStart(r *mux.Router) {
	WatcherInstance().Watch(a.Watchers)
	gateway := NewHTTPGateway()
	gateway.loadBalancing(WatcherInstance())
	log.Info("HttpsGateway start at:", a.getServerAddr(CONNECTION_TYPE_HTTP))
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Secure World")
	})
	server := &http.Server{
		Addr:    ":443",
		Handler: gateway.MakeHttpHandle(r, a.EtcdEndPoints.ServerId),
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	//go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	//go http.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), gateway.MakeHttpHandle(r, a.EtcdEndPoints.ServerId))
	server.ListenAndServeTLS("", "")
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
		//todo free or else
	}
	RegisterContinueSignal(syscall.SIGTERM, process)
	RegisterContinueSignal(syscall.SIGINT, process)
}
