package grpc

import (
	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"

	"context"
)

type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (interface{}, error)
}

type Server struct {
	e      endpoint.Endpoint
	logger log.Logger
}

func NewServer(e endpoint.Endpoint, ) *Server { return &Server{e: e} }

func (s Server) ServeGRPC(ctx context.Context, req interface{}) (response interface{}, err error) {
	return s.e(ctx, req)
}
