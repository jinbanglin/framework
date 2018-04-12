package grpc

import (
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/log"

	"context"
)

type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (interface{}, error)
}

type Server struct {
	e      moss.Endpoint
	logger log.Logger
}

func NewServer(e moss.Endpoint, ) *Server { return &Server{e: e} }

func (s Server) ServeGRPC(ctx context.Context, req interface{}) (response interface{}, err error) {
	return s.e(ctx, req)
}
