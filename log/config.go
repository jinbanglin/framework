package log

import (
	"encoding/json"
	"io/ioutil"
)

const (
	_DEBUG    level = iota + 1
	_INFO
	_WARN
	_ERR
	_DISASTER
)

var (
	coreDead    coreStatus = 2 //gLogger is dead
	coreBlock   coreStatus = 0 //gLogger is block
	coreRunning coreStatus = 1 //gLogger is running
)
var out = "stdout"
var maxSize int = 256 * 1024 * 1024
var bucketLen int = 1024
var bufSize int = 2 * 1024 * 1024
var fileName string = "moss"
var filePath string = getCurrentDirectory()
var levelFlag level = _DEBUG
var pollInterval = 500

type level = uint8
type coreStatus = uint32

func loadConfig() {
	//todo use viper
	b, err := ioutil.ReadFile("logx.json")
	if err != nil {
		b, err = ioutil.ReadFile("config.json")
		if err != nil {
			return
		}
	}
	var config config
	if err = json.Unmarshal(b, &config); err != nil {
		return
	}
	if x := config.Lbucketlen; x != 0 {
		bucketLen = x
	}
	if x := config.Lfilename; x != "" {
		fileName = x
	}
	if x := config.Lfilepath; x != "" {
		filePath = x
	}
	if x := config.Llevel; x != 0 {
		levelFlag = level(x)
	}
	if x := config.Lmaxsize; x != 0 {
		maxSize = x * 1024 * 1024
	}
	if x := config.Lout; x != "" {
		out = x
	}
	if x := config.Lpollerinterval; x != 0 {
		pollInterval = x
	}
}
