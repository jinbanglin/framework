package main

import (
	"time"

	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/_example/addgrpcsvc/hello_moss"
	"github.com/jinbanglin/moss/_example/addgrpcsvc/register"
	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/kernel"
)

func main() {
	kernel.AppServer.SetupConfig(_example.Addgrpcsvc)
	kernel.GrpcInstance().RegisterProtoInvokeFunc(900001, &pb.HelloMoss{}, hello_moss.HelloWorldHandler)
	kernel.GrpcInstance().RegisterProtoInvokeFunc(900003, &pb.RegisterReq{}, register.RegisterHandler)
	kernel.AppServer.GrpcServerStart()
	kernel.AppServer.Stop(10*time.Second, func() {
		//TODO free
	})
	kernel.AppServer.Run()
}
