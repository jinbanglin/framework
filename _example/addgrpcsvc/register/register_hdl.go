package register

import (
	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/log"

	"context"
	"fmt"
)

func RegisterEndpoint(_ context.Context, request interface{}) (response interface{}, err error) {
	fmt.Println("--------------1--------")
	log.Info("RegisterEndpoint")
	req := request.(*pb.RegisterReq)
	res := &pb.RegisterRes{
		UserName:  req.UserName,
		UserPhone: req.UserPhone,
		UserId:    "123456789",
		Audience:  "987654321",
	}
	return res, nil
}
