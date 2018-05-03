package ipc

import (
	"github.com/gogo/protobuf/proto"
	"github.com/jinbanglin/moss/endpoint"
)

import (
	"github.com/jinbanglin/moss/payload"
	transportgrpc "github.com/jinbanglin/moss/transport/grpc"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) (endpoint endpoint.Endpoint) {
	return transportgrpc.NewClient(conn, "payload.Invoking", "Invoking", payload.MossPacket{}).Endpoint()
}

type GRPCServer struct {
	Scheduler *GPRCInvoking
}

var GGRPCServer *GRPCServer

func init() {
	GGRPCServer = &GRPCServer{Scheduler: &GPRCInvoking{Scheduler: make(map[uint32]*SchedulerHandler)}}
}

func RegisterGRPCHandler(serviceCode uint32, request proto.Message, endpoint endpoint.Endpoint) {
	GGRPCServer.Scheduler.RegisterHandler(serviceCode, request, endpoint)
}
