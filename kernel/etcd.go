package kernel

import (
	"golang.org/x/net/context"
	"time"

	"github.com/jinbanglin/moss/discovery/etcdv3"
)

func NewEtcdV3Client() etcdv3.Client {
	client, err := etcdv3.NewClient(context.Background(), []string{"127.0.0.1:2379"}, etcdv3.ClientOptions{
		CACert:        "",
		Cert:          "",
		Key:           "",
		Username:      "",
		Password:      "",
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 30,
	})
	if err != nil {
		panic(err)
	}
	return client
}

