package http

import (
	"context"
	"net/http"

	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/payload"
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
