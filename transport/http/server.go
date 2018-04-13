package http

import (
	"context"
	"net/http"

	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"
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
		log.Errorf("dec |request=%v |err=%v",request,err)
		s.errorEncoder(ctx, request, w)
		return
	}
	response, err := s.e(ctx, request)
	if err != nil {
		log.Errorf("FROM |response=%v |err=%v",response,err)
		s.errorEncoder(ctx, response, w)
		return
	}
	s.enc(ctx, w, response)
}

type ErrorEncoder func(ctx context.Context, response interface{}, w http.ResponseWriter)
