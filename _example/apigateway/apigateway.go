package main

import (
	"time"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/_example/apigateway/opensvc"
)

func main() {
	moss.AppServer.SetupConfig(_example.ServiceNameAPIGateway)
	go moss.AppServer.HTTPTLSGatewayStart(opensvc.MakeOpensvcHTTPHandler())
	moss.AppServer.Stop(10*time.Second, func() {
		//TODO free
	})
	moss.AppServer.Run()
}
