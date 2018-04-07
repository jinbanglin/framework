package hello_moss

import (
	"moss/_example/pb"

	"golang.org/x/net/context"
)

func HelloWorldHandler(_ context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*pb.HelloMoss)
	req.HelloMoss = "hello boy"
	return req, nil
}
