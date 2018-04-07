package main

import (
	"time"

	"moss/_example"
	"moss/_example/apigateway/opensvc"
	"moss/kernel"
)

func main() {
	kernel.AppServer.SetupConfig(_example.Apigateway)
	go kernel.AppServer.HttpGatewayStart(opensvc.MakeOpensvcHTTPHandler())
	kernel.AppServer.Stop(10*time.Second, func() {
		//TODO 资源释放函数
	})
	kernel.AppServer.Run()
}
