package kernel

import (
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/kernel/addtransport"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/transport/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
)

type GrpcServer struct {
	gScheduler *addtransport.Scheduler
}

var gGrpcServer *GrpcServer

func GrpcInstance() *GrpcServer {
	if gGrpcServer == nil {
		gGrpcServer = &GrpcServer{
			gScheduler: addtransport.NewGrpcServer(opentracing.GlobalTracer(), []grpc.ServerOption{
				grpc.ServerErrorLogger(log.Logger{}),
			}),
		}
	}
	return gGrpcServer
}

func (g *GrpcServer) RegisterProtoInvokeFunc(serviceCode uint32, request proto.Message, endpoint moss.Endpoint) {
	g.gScheduler.RegisterHandler(serviceCode, request, endpoint)
}
