package main

import (
	"time"

	"moss/_example"
	"moss/_example/addgrpcsvc/hello_moss"
	"moss/_example/addgrpcsvc/register"
	"moss/_example/pb"
	"moss/kernel"
)

func main() {
	kernel.AppServer.SetupConfig(_example.Addgrpcsvc)
	kernel.GrpcInstance().RegisterProtoInvokeFunc(900001, &pb.HelloMoss{}, hello_moss.HelloWorldHandler)
	kernel.GrpcInstance().RegisterProtoInvokeFunc(900003, &pb.RegisterReq{}, register.RegisterHandler)
	kernel.AppServer.GrpcServerStart()
	kernel.AppServer.Stop(10*time.Second, func() {
		//TODO 资源释放函数
	})
	kernel.AppServer.Run()
}
