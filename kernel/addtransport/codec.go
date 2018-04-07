package addtransport

import (
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
)

//only support json or protobuf
type Codecer interface {
	Marshal(v proto.Message) ([]byte, error)
	Unmarshal(payload []byte, message proto.Message) error
}

type JSONCodec struct{}

func (JSONCodec) Marshal(v proto.Message) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (JSONCodec) Unmarshal(payload []byte, message proto.Message) error {
	return jsoniter.Unmarshal(payload, message)
}

type PROTOCodec struct{}

func (PROTOCodec) Marshal(v proto.Message) ([]byte, error) {
	return proto.Marshal(v)
}

func (PROTOCodec) Unmarshal(payload []byte, message proto.Message) error {
	return proto.Unmarshal(payload, message)
}

var gJSONCodec = JSONCodec{}
var gPROTOCodec = PROTOCodec{}

func GetCodecerByServiceCode(serviceCode uint32) Codecer {
	switch serviceCode&1 == 0 {
	case true:
		return gPROTOCodec
	case false:
		return gJSONCodec
	}
	return nil
}
