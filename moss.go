package moss

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/jinbanglin/moss/discovery/etcdv3"
	"github.com/jinbanglin/moss/distributor"
	"github.com/jinbanglin/moss/log"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/reflection"

	"github.com/golang/crypto/acme/autocert"
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
	a.registerEtcdV3(addr, a.EtcdEndPoints.EtcdEndpoints)
	if len(a.Watchers) > 0 {
		distributor.WatcherInstance().Watch(a.GetWatchNames(), a.EtcdEndPoints.EtcdEndpoints)
	}
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Info("✨MOSS✨ |start at:", addr)
	baseServer := grpc.NewServer()
	payload.RegisterInvokingServer(baseServer, distributor.GGRPCServer.Scheduler)
	reflection.Register(baseServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		panic(err)
	}
}

func (a *appServer) AddFileSvc(r *mux.Router) {
	//r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("/root/data/view/"))))
	if fileSvc := viper.GetString("server.static"); strings.Count(fileSvc, "") > 0 {
		log.Infof("✨MOSS✨ |file service route at=%s", fileSvc)
		r.PathPrefix(fileSvc).Handler(http.StripPrefix(fileSvc, http.FileServer(http.Dir(viper.GetString("server.dir")))))
	}
}

func (a *appServer) MakeGateway(r *mux.Router) {
	distributor.WatcherInstance().Watch(a.GetWatchNames(), a.EtcdEndPoints.EtcdEndpoints)
	gateway := distributor.NewHTTPGateway()
	gateway.LoadBalancing(distributor.WatcherInstance())
	a.AddFileSvc(r)
	go a.AddHTTPServer(r, gateway)
	go a.AddTLSServer(r, gateway)
}

func (a *appServer) AddHTTPServer(r *mux.Router, gateway *distributor.HTTPGateway) {
	r.HandleFunc("/moss", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "🔥 MOSS 🔥")
	})
	log.Info("✨MOSS✨ |http start at:", a.getServerAddr(CONNECTION_TYPE_HTTP))
	go log.Debug(http.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), gateway.MakeHttpHandle(r)))
}

func (a *appServer) AddTLSServer(r *mux.Router, gateway *distributor.HTTPGateway) {
	log.Info("✨MOSS✨ |https start at:", ":443")
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	server := &http.Server{
		Addr:    ":443",
		Handler: gateway.MakeHttpHandle(r),
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go log.Debug(server.ListenAndServeTLS("", ""))
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
	log.Infof("✨MOSS✨ |register etcd key=%s", "/"+string(a.ServiceName)+"/"+a.ConfigManager.EtcdEndPoints.ServerId)
	log.Info("✨MOSS✨ |register etcd value", serverAddr)
}

func (a *appServer) Stop(timeout time.Duration, f ...func()) {
	a.wait.Add(1)
	process := func() {
		defer a.wait.Done()
		time.AfterFunc(timeout, func() {
			log.Info("✨MOSS✨ |server stop", "server_id", a.EtcdEndPoints.ServerId, "server_name", a.ServiceName)
		})
		for _, v := range f {
			v()
		}
		//todo free or else
	}
	distributor.RegisterContinueSignal(syscall.SIGTERM, process)
	distributor.RegisterContinueSignal(syscall.SIGINT, process)
}
