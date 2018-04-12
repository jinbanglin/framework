package distributor

import (
	"github.com/golang/protobuf/proto"
	"github.com/jinbanglin/moss/endpoint"
)

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
