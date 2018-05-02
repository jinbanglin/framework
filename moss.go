package moss

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/jinbanglin/moss/ipc"
	"github.com/jinbanglin/log"
	"github.com/jinbanglin/moss/sd/etcdv3"
	mosshttp "github.com/jinbanglin/moss/transport/http"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"golang.org/x/net/http2"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/reflection"

	"github.com/golang/crypto/acme/autocert"
	"google.golang.org/grpc"

	"github.com/jinbanglin/moss/payload"
)

var mode = flag.String("m", "", "project mode prd or dev")
var AppServer = &appServer{}

type appServer struct {
	*ConfigManager
	ServiceName
	wait sync.WaitGroup
}

func (a *appServer) SetupConfig(name ServiceName, f ...func()) {
	flag.Parse()
	AppServer = &appServer{ConfigManager: &ConfigManager{EtcdEndPoints: &EtcdV3{}}, ServiceName: name}
	AppServer.setupConfig(name, *mode, f...)
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
		ipc.WatcherInstance().Watch(a.GetWatchNames(), a.EtcdEndPoints.EtcdEndpoints)
	}
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Info("MOSS |start at", addr)
	baseServer := grpc.NewServer()
	payload.RegisterInvokingServer(baseServer, ipc.GGRPCServer.Scheduler)
	reflection.Register(baseServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		panic(err)
	}
}

func (a *appServer) AddFileSvc(r *mux.Router) {
	//r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("/root/data/view/"))))
	if fileSvc := viper.GetString("server.vf_dir"); strings.Count(fileSvc, "") > 0 {
		fs := afero.NewOsFs()
		stat, err := fs.Stat(fileSvc)
		if err != nil || !stat.IsDir() {
			log.Info(" |Stat |err=", err)
			if err = fs.MkdirAll(fileSvc, os.ModePerm); err != nil {
				panic(err)
			}
		}
		log.Infof("MOSS |file service route at=%s", "/web/")
		log.Infof("MOSS |file service file dir is=%s", fileSvc)
		r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir(fileSvc))))
	}
}

func (a *appServer) MakeGateway(r *mux.Router) {
	ipc.WatcherInstance().Watch(a.GetWatchNames(), a.EtcdEndPoints.EtcdEndpoints)
	gateway := ipc.NewHTTPGateway()
	gateway.LoadBalancing(ipc.WatcherInstance())
	a.AddFileSvc(r)
	r.HandleFunc("/moss", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "🔥 MOSS 🔥")
	})
	go a.AddHTTPServer(r, gateway)
	go a.AddTLSServer(r, gateway)
}

func (a *appServer) AddHTTPServer(r *mux.Router, gateway *ipc.HTTPGateway) {
	log.Info("MOSS |http start at", a.getServerAddr(CONNECTION_TYPE_HTTP))
	//log.Debug(http.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), gateway.MakeHttpHandle(r)))
	mosshttp.ListenAndServe(a.getServerAddr(CONNECTION_TYPE_HTTP), fasthttpadaptor.NewFastHTTPHandler(gateway.MakeHttpHandle(r)))
}

func (a *appServer) AddTLSServer(r *mux.Router, gateway *ipc.HTTPGateway) {
	log.Info("MOSS |https start at", ":443")
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	server := &http.Server{
		Addr:    ":443",
		Handler: gateway.MakeHttpHandle(r),
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			Time:           time.Now,
			NextProtos:     []string{http2.NextProtoTLS, "http/1.1"},
			MinVersion:     tls.VersionTLS12,
		},
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	log.Debug(server.ListenAndServeTLS("", ""))
}

func (a *appServer) Run() {
	a.wait.Wait()
}

func (a *appServer) registerEtcdV3(serverAddr string, etcdAddress []string) {
	etcdv3.DefaultEtcdV3Client(etcdAddress).Register(etcdv3.Service{
		Key:   "/" + string(a.ServiceName) + "/" + a.ConfigManager.EtcdEndPoints.ServerId,
		Value: serverAddr,
		TTL:   etcdv3.NewTTLOption(3, 10),
	})
	log.Infof("MOSS |register etcd key=%s", "/"+string(a.ServiceName)+"/"+a.ConfigManager.EtcdEndPoints.ServerId)
	log.Info("MOSS |register etcd value", serverAddr)
}

func (a *appServer) Stop(timeout time.Duration, f ...func()) {
	a.wait.Add(1)
	process := func() {
		defer a.wait.Done()
		time.AfterFunc(timeout, func() {
			log.Info("MOSS |server stop", "server_id", a.EtcdEndPoints.ServerId, "server_name", a.ServiceName)
		})
		for _, v := range f {
			v()
		}
		//todo free or else
	}
	ipc.RegisterContinueSignal(syscall.SIGTERM, process)
	ipc.RegisterContinueSignal(syscall.SIGINT, process)
}
