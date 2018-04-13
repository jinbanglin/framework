package grpc

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jinbanglin/moss/endpoint"
	"google.golang.org/grpc"
)

type Client struct {
	client      *grpc.ClientConn
	serviceName string
	method      string
	grpcReply   reflect.Type
}

func NewClient(cc *grpc.ClientConn, serviceName string, method string, grpcReply interface{}) *Client {
	return &Client{
		client:    cc,
		method:    fmt.Sprintf("/%s/%s", serviceName, method),
		grpcReply: reflect.TypeOf(reflect.Indirect(reflect.ValueOf(grpcReply)).Interface()),
	}
}

func (c Client) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		response = reflect.New(c.grpcReply).Interface()
		err = c.client.Invoke(ctx, c.method, request, response)
		return
	}
}
