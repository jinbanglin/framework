package main

import (
	"time"

	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/_example/apigateway/opensvc"
	"github.com/jinbanglin/moss/distributor"
)

func main() {
	distributor.AppServer.SetupConfig(_example.ServiceNameAPIGateway)
	go distributor.AppServer.HTTPTLSGatewayStart(opensvc.MakeOpensvcHTTPHandler())
	distributor.AppServer.Stop(10*time.Second, func() {
		//TODO free
	})
	distributor.AppServer.Run()
}
