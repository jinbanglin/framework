package grpc

import (
	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/log"

	"context"
)

type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (interface{}, error)
}

type Server struct {
	e endpoint.Endpoint
}

func NewServer(e endpoint.Endpoint) *Server { return &Server{e: e} }

func (s Server) ServeGRPC(ctx context.Context, req interface{}) (response interface{}, err error) {
	response, err = s.e(ctx, req)
	if err != nil {
		log.Errorf("MOSS |err=%v", err)
		return
	}
	return response, nil
}
