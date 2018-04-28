package http

import (
	"context"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/payload"
	"github.com/kavu/go_reuseport"
	"github.com/valyala/fasthttp"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

type Server struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	errorEncoder ErrorEncoder
}

func NewServer(e endpoint.Endpoint, dec DecodeRequestFunc, enc EncodeResponseFunc, errorEncoder ErrorEncoder) *Server {
	return &Server{e: e, dec: dec, enc: enc, errorEncoder: errorEncoder}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request, err := s.dec(ctx, r)
	if err != nil {
		log.Errorf("MOSS |dec |request=%v |err=%v", request, err)
		s.errorEncoder(ctx, request, w)
		return
	}
	log.Infof("MOSS |FROM |request=%v ", request)
	response, err := s.e(ctx, request)
	if err != nil || response == nil {
		log.Infof("MOSS |FROM RPC|response=%v |err=%v", request, err)
		response = payload.MossPacket{MossMessage: payload.StatusText(payload.StatusInternalServerError)}
		s.errorEncoder(ctx, response, w)
		return
	}
	log.Infof("MOSS |FROM RPC|response=%v ", response)
	s.enc(ctx, w, response)
}

type ErrorEncoder func(ctx context.Context, response interface{}, w http.ResponseWriter)

func ListenAndServe(address string, handler func(*fasthttp.RequestCtx)) {
	hdr := func(listener net.Listener) {
		defer listener.Close()
		server := &fasthttp.Server{
			Handler: handler,
			ReadTimeout: 60 * time.Second,
			DisableKeepalive:true,
		}
		server.Serve(listener)
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		l, err := reuseport.Listen("tcp4", address)
		if err != nil {
			panic(err)
		}
		if (i + 1 ) == runtime.NumCPU() {
			hdr(l)
		} else {
			go hdr(l)
		}
	}
}