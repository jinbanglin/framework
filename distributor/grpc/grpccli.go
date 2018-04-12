package grpc

import (
	"github.com/jinbanglin/moss"
	"google.golang.org/grpc"
	"github.com/jinbanglin/moss/payload"
	transportgrpc "github.com/jinbanglin/moss/transport/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) (endpoint moss.Endpoint) {
	return transportgrpc.NewClient(
		conn,
		"payload.Invoking",
		"Invoking",
		payload.MossPacket{},
	).Endpoint()
}
