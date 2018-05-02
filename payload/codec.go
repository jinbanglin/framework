package payload

import (
	"github.com/gogo/protobuf/proto"
	"github.com/jinbanglin/log"
	"github.com/json-iterator/go"
)

//only support json or protobuf
type Codecer interface {
	Marshal(v proto.Message) ([]byte)
	Unmarshal(payload []byte, message proto.Message) error
}

type JSONCodec struct{}

func (JSONCodec) Marshal(v proto.Message) ([]byte) {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		log.Errorf("MOSS |Marshal |err=%v", err)
	}
	return b
}

func (JSONCodec) Unmarshal(payload []byte, message proto.Message) error {
	return jsoniter.Unmarshal(payload, message)
}

type PROTOCodec struct{}

func (PROTOCodec) Marshal(v proto.Message) ([]byte) {
	b, err := proto.Marshal(v)
	if err != nil {
		log.Errorf("MOSS |Marshal |err=%v", err)
	}
	return b
}

func (PROTOCodec) Unmarshal(payload []byte, message proto.Message) error {
	return proto.Unmarshal(payload, message)
}

var gJSONCodec = JSONCodec{}
var gPROTOCodec = PROTOCodec{}

func GetCodec(serviceCode uint32) Codecer {
	switch serviceCode&1 == 0 {
	case true:
		return gPROTOCodec
	case false:
		return gJSONCodec
	}
	return nil
}
