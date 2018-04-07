package kernel

import (
	"fmt"

	"moss/log"

	"github.com/json-iterator/go"
	"github.com/spf13/viper"

	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"strings"
	"os"
)

type ConnectionType = string
type ServiceName = string

const (
	CONNECTION_TYPE_DEBUG   = "DEBUG"
	CONNECTION_TYPE_METRICS = "METRICS"
	CONNECTION_TYPE_HTTP    = "HTTP"
	CONNECTION_TYPE_GRPC    = "GRPC"
)

type ConfigManager struct {
	configFile    string
	EtcdEndPoints *EtcdV3
	Connections   []*Connection
	Watchers      []*Watch
}

type EtcdV3 struct {
	ServerId      string   `json:"server_id"`
	EtcdGroup     string   `json:"etcd_group"`
	EtcdEndpoints []string `json:"etcd_endpoints"`
}

type Connection struct {
	Addr     string         `json:"addr"`
	ConnType ConnectionType `json:"conn_type"`
}

type Watch struct {
	ServiceName ServiceName `json:"service_name"`
}

//use "github.com/spf13/viper" to config app

func (c *ConfigManager) setupConfig(serviceName ServiceName, f ...func()) {
	c.configFile = GetCurrentDirectory() + "/" + serviceName + ".toml"
	f = append(f, func() {
		{
			b, _ := jsoniter.Marshal(viper.Get("connection"))
			if err := jsoniter.Unmarshal(b, &c.Connections); err != nil {
				log.Error("chaos error |json Unmarshal", err)
				return
			}
		}
		{
			b, _ := jsoniter.Marshal(viper.Get("etcdv3"))
			if err := jsoniter.Unmarshal(b, c.EtcdEndPoints); err != nil {
				log.Error("chaos error |json Unmarshal", err)
				return
			}
		}
		{
			b, _ := jsoniter.Marshal(viper.Get("watch"))
			if err := jsoniter.Unmarshal(b, &c.Watchers); err != nil {
				log.Error("chaos error |json Unmarshal", err)
				return
			}
		}
	})
	c.fsnotify(f...)
}

func (c *ConfigManager) fsnotify(f ...func()) {
	viper.SetConfigType("toml")
	viper.SetConfigFile(c.configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	for _, v := range f {
		v()
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("ConfigManager |Config file changed:", e.Name)
		for _, v := range f {
			v()
		}
	})
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
