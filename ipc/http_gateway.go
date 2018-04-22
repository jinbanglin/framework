package ipc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss/endpoint"
	"github.com/spf13/viper"
)

type HTTPGateway struct {
	cache       map[string]string
	proxyServer MutilEndpoints
	prefix      string
}

func NewHTTPGateway() *HTTPGateway {
	return &HTTPGateway{
		cache: make(map[string]string),
		proxyServer: MutilEndpoints{
			Endpoints: make(map[string]endpoint.Endpoint),
		},
		prefix: viper.GetString("server.prefix"),
	}
}

func (h *HTTPGateway) LoadBalancing(watcher *Watcher) {
	for k, v := range WatcherInstance().watchers {
		h.proxyServer.Endpoints[h.getHostHeader(k)] = v.endpoint
	}
}

func (h *HTTPGateway) getHostHeader(name string) string {
	if strings.HasPrefix(h.prefix, "/") {
		h.prefix = strings.TrimPrefix(h.prefix, "/")
	}
	format := "/%s/%s/{service_code}"
	return fmt.Sprintf(format, h.prefix, name)
}

func (h *HTTPGateway) GetServiceTpl(name string) string {
	return h.cache[name]
}

func (h *HTTPGateway) MakeHttpHandle(r *mux.Router) http.Handler {
	return MakeHTTPGateway(r, h.proxyServer)
}
