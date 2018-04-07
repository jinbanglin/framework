package register

import (
	"moss/_example/pb"
	"moss/log"

	"golang.org/x/net/context"
)

func RegisterHandler(_ context.Context, request interface{}) (response interface{}, err error) {
	log.Info("RegisterHandler")
	req := request.(*pb.RegisterReq)
	res := &pb.RegisterRes{
		UserName:  req.UserName,
		UserPhone: req.UserPhone,
		UserId:    "123456789",
		Audience:  "987654321",
	}
	return res, nil
}
