package main

import (
	"time"

	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/_example/sns/register"
	"github.com/jinbanglin/moss/ipc"
)

func main() {
	moss.AppServer.SetupConfig(_example.ServiceNameSns)
	ipc.RegisterGRPCHandler(900003, &pb.RegisterReq{}, register.RegisterHandler)
	moss.AppServer.GRPCServerStart()
	moss.AppServer.Stop(10*time.Second, func() {
		//TODO free
	})
	moss.AppServer.Run()
}
