package main

import (
	"time"

	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/_example/sns/register"
	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/distributor"
)

func main() {
	distributor.AppServer.SetupConfig(_example.ServiceNameSns)
	distributor.RegisterGRPCHandler(900003, &pb.RegisterReq{}, register.RegisterHandler)
	distributor.AppServer.GRPCServerStart()
	distributor.AppServer.Stop(10*time.Second, func() {
		//TODO free
	})
	distributor.AppServer.Run()
}
