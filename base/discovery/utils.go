package discovery

import (
	"github.com/gogo/protobuf/proto"
	"github.com/perlin-network/noise/internal/protobuf"
	"github.com/pkg/errors"
)

const (
	DiscoveryServiceID   = 5
	opCodePing           = 1
	opCodePong           = 2
	opCodeLookupRequest  = 3
	opCodeLookupResponse = 4
)

func toProtobufMessage(opcode int, content proto.Message) (*protobuf.Message, error) {
	raw, err := proto.Marshal(content)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal reply")
	}
	msg := &protobuf.Message{
		Message: raw,
		Opcode:  uint32(opcode),
	}
	return msg, nil
}

func fromProtobufMessage(msg *protobuf.Message) (proto.Message, error) {

	var content proto.Message
	if err := proto.Unmarshal(msg.Message, content); err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal type")
	}

	return content, nil
}