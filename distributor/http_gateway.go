package distributor

import (
	"fmt"
	"net/http"
	"strings"

	distributorhttp "github.com/jinbanglin/moss/distributor/http"

	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss"
	"github.com/spf13/viper"
)

type HttpGateway struct {
	cache       map[moss.ServiceName]string
	proxyServer distributorhttp.MutilEndpoints
	prefix      string
}

func NewHTTPGateway() *HttpGateway {
	return &HttpGateway{
		cache: make(map[moss.ServiceName]string),
		proxyServer: distributorhttp.MutilEndpoints{
			Endpoints: make(map[string]moss.Endpoint),
		},
		prefix: viper.GetString("server.http_prefix"),
	}
}

func (h *HttpGateway) loadBalancing(watcher *Watcher) {
	for k, v := range WatcherInstance().watchers {
		h.proxyServer.Endpoints[h.getHostHeader(k)] = v.endpoint
	}
}

func (h *HttpGateway) getHostHeader(name moss.ServiceName) string {
	if strings.HasPrefix(h.prefix, "/") {
		h.prefix = strings.TrimPrefix(h.prefix, "/")
	}
	format := "/%s/%s/{service_code}/"
	return fmt.Sprintf(format, h.prefix, name)
}

func (h *HttpGateway) GetServiceTpl(name moss.ServiceName) string {
	return h.cache[name]
}

func (h *HttpGateway) MakeHttpHandle(r *mux.Router, serverId string) http.Handler {
	return distributorhttp.MakeHTTPGateway(r, h.proxyServer, serverId)
}
