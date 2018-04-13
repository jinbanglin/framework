package distributor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss/endpoint"
	"github.com/spf13/viper"
)

type HttpGateway struct {
	cache       map[string]string
	proxyServer MutilEndpoints
	prefix      string
}

func NewHTTPGateway() *HttpGateway {
	return &HttpGateway{
		cache: make(map[string]string),
		proxyServer: MutilEndpoints{
			Endpoints: make(map[string]endpoint.Endpoint),
		},
		prefix: viper.GetString("server.http_prefix"),
	}
}

func (h *HttpGateway) LoadBalancing(watcher *Watcher) {
	for k, v := range WatcherInstance().watchers {
		h.proxyServer.Endpoints[h.getHostHeader(k)] = v.endpoint
	}
}

func (h *HttpGateway) getHostHeader(name string) string {
	if strings.HasPrefix(h.prefix, "/") {
		h.prefix = strings.TrimPrefix(h.prefix, "/")
	}
	format := "/%s/%s/{service_code}/"
	return fmt.Sprintf(format, h.prefix, name)
}

func (h *HttpGateway) GetServiceTpl(name string) string {
	return h.cache[name]
}

func (h *HttpGateway) MakeHttpHandle(r *mux.Router, serverId string) http.Handler {
	return MakeHTTPGateway(r, h.proxyServer, serverId)
}
