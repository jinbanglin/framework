package kernel

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jinbanglin/moss/kernel/addendpoint"
	"github.com/jinbanglin/moss/kernel/addtransport"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

type HttpGateway struct {
	cache       map[ServiceName]string //service name:route url array
	proxyServer addtransport.MutilEndpoints
	prefix      string
}

func NewHTTPGateway() *HttpGateway {
	return &HttpGateway{
		cache: make(map[ServiceName]string),
		proxyServer: addtransport.MutilEndpoints{
			Endpoints: make(map[string]addendpoint.Set),
		},
		prefix: viper.GetString("server.http_prefix"),
	}
}

func (h *HttpGateway) loadBalancing(watcher *Watcher) {
	for k, v := range WatcherInstance().watchers {
		h.proxyServer.Endpoints[h.getHostHeader(k)] = v.endpoint
	}
}

func (h *HttpGateway) getHostHeader(name ServiceName) string {
	if strings.HasPrefix(h.prefix, "/") {
		h.prefix = strings.TrimPrefix(h.prefix, "/")
	}
	format := "/%s/%s/{protocol}/"
	return fmt.Sprintf(format, h.prefix, name)
}

func (h *HttpGateway) GetServiceTpl(name ServiceName) string {
	return h.cache[name]
}

func (h *HttpGateway) MakeHttpHandle(r *mux.Router, serverId string) http.Handler {
	return addtransport.MakeMutilHTTPHandler(r, h.proxyServer, opentracing.GlobalTracer(), serverId)
}
