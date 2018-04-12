package distributor

import (
	"github.com/golang/protobuf/proto"
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/distributor/grpc"
)

type GRPCServer struct {
	Scheduler *grpc.GPRCInvoking
}

var gGRPCServer *GRPCServer

func init() {
	gGRPCServer = &GRPCServer{Scheduler: &grpc.GPRCInvoking{Scheduler: make(map[uint32]*grpc.SchedulerHandler)}}
}

func RegisterGRPCHandler(serviceCode uint32, request proto.Message, endpoint moss.Endpoint) {
	gGRPCServer.Scheduler.RegisterHandler(serviceCode, request, endpoint)
}
