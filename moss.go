package moss

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"syscall"
	"time"

	"github.com/jinbanglin/moss/discovery/etcdv3"
	"github.com/jinbanglin/moss/distributor"
	"github.com/jinbanglin/moss/log"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/reflection"

	"github.com/golang/crypto/acme/autocert"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/jinbanglin/moss/payload"
)

var AppServer = &appServer{}

type appServer struct {
	*ConfigManager
	ServiceName
	wait sync.WaitGroup
}

func (a *appServer) SetupConfig(name ServiceName, f ...func()) {
	AppServer = &appServer{ConfigManager: &ConfigManager{EtcdEndPoints: &EtcdV3{}}, ServiceName: name}
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

func (a *appServer) GRPCServerStart() {
	addr := a.getServerAddr(CONNECTION_TYPE_GRPC)
	a.registerEtcdV3(addr,a.EtcdEndPoints.EtcdEndpoints)
	if len(a.Watchers) > 0 {
		distributor.WatcherInstance().Watch(a.GetWatchNames(),a.EtcdEndPoints.EtcdEndpoints)
	}
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Info("start at:", addr)
	baseServer := grpc.NewServer()
	payload.RegisterInvokingServer(baseServer, distributor.GGRPCServer.Scheduler)
	reflection.Register(baseServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		panic(err)
	}
}

func (a *appServer) HTTPTLSGatewayStart(r *mux.Router) {
	distributor.WatcherInstance().Watch(a.GetWatchNames(), a.EtcdEndPoints.EtcdEndpoints)
	gateway := distributor.NewHTTPGateway()
	gateway.LoadBalancing(distributor.WatcherInstance())
	log.Info("start at:", a.getServerAddr(CONNECTION_TYPE_HTTP))
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Secure World")
	})
	//r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("/root/data/view/"))))
	r.PathPrefix(viper.GetString("server.static")).Handler(http.StripPrefix(viper.GetString("server.static"), http.FileServer(http.Dir("/root/data/view/"))))
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	server := &http.Server{
		Addr:    ":443",
		Handler: gateway.MakeHttpHandle(r, a.EtcdEndPoints.ServerId),
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), gateway.MakeHttpHandle(r, a.EtcdEndPoints.ServerId))
	server.ListenAndServeTLS("", "")
}

func (a *appServer) Run() {
	a.wait.Wait()
}

func (a *appServer) registerEtcdV3(serverAddr string, etcdAddress []string) {
	etcdv3.DefaultEtcdV3Client(etcdAddress).Register(etcdv3.Service{
		Key:   "/" + string(a.ServiceName) + "/" + a.ConfigManager.EtcdEndPoints.ServerId,
		Value: serverAddr,
		TTL:   etcdv3.NewTTLOption(0, 0),
	})
	log.Infof("register etcd key=%s", "/"+string(a.ServiceName)+"/"+a.ConfigManager.EtcdEndPoints.ServerId)
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
	distributor.RegisterContinueSignal(syscall.SIGTERM, process)
	distributor.RegisterContinueSignal(syscall.SIGINT, process)
}
