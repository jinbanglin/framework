package distributor

import (
	"github.com/jinbanglin/moss"
	"github.com/golang/protobuf/proto"
)

type GRPCServer struct {
	Scheduler *GPRCInvoking
}

type registerFunc func() (serviceCode uint32, request proto.Message, endpoint moss.Endpoint)

var gGRPCServer *GRPCServer

func NewGRPCServer() *GRPCServer {
	gGRPCServer = &GRPCServer{Scheduler: &GPRCInvoking{make(map[uint32]*SchedulerHandler)}}
	return gGRPCServer
}

func (g *GRPCServer) registerHandler(serviceCode uint32, request proto.Message, endpoint moss.Endpoint) {
	g.Scheduler.RegisterHandler(serviceCode, request, endpoint)
}

func RegisterHandler(handlers ...registerFunc) {
	for _, v := range handlers {
		gGRPCServer.Scheduler.RegisterHandler(v())
	}
}
