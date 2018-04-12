package http

import (
	"context"
	"net/http"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/log"
)

type Server struct {
	e            moss.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	errorEncoder ErrorEncoder
}

func NewServer(e moss.Endpoint, dec DecodeRequestFunc, enc EncodeResponseFunc, errorEncoder ErrorEncoder) *Server {
	return &Server{e: e, dec: dec, enc: enc, errorEncoder: errorEncoder}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request, err := s.dec(ctx, r)
	if err != nil {
		log.Error(err)
		s.errorEncoder(ctx, request, w)
		return
	}
	response, err := s.e(ctx, request)
	if err != nil {
		log.Error(err)
		s.errorEncoder(ctx, response, w)
		return
	}
	s.enc(ctx, w, response)
}

type ErrorEncoder func(ctx context.Context, response interface{}, w http.ResponseWriter)
