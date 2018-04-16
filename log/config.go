package log

import (
	"github.com/spf13/viper"
)

type level = uint8
type coreStatus = uint32

const (
	_DEBUG    level = iota + 1
	_INFO
	_WARN
	_ERR
	_DISASTER
)

const (
	B  = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

var (
	coreDead    coreStatus = 2 //gLogger is dead
	coreBlock   coreStatus = 0 //gLogger is block
	coreRunning coreStatus = 1 //gLogger is running
)
var gSetOut = "file"
var gSetMaxSize = 256 * MB
var gSetBucketLen = 1024
var gSetBufSize = 2 * MB
var gSetFilename = "✨MOSS✨"
var gSetFilePath = getCurrentDirectory()
var gSetLevel = _DEBUG
var gSetPollerInterval = 500

func setupConfig() {
	if value := viper.GetInt("log.bucketlen"); value > 0 {
		gSetBucketLen = value
	}
	if value := viper.GetString("log.filename"); value != "" {
		gSetFilename = value
	}
	if value := viper.GetString("log.filepath"); value != "" {
		gSetFilePath = value
	}
	if value := viper.GetInt("log.level"); value > 0 {
		gSetLevel = level(value)
	}
	if value := viper.GetInt("log.maxsize"); value > 0 {
		gSetMaxSize = value*MB
	}
	if value := viper.GetString("log.out"); value != "" {
		gSetOut = value
	}
	if value := viper.GetInt("log.interval"); value > 0 {
		gSetPollerInterval = value
	}
}
