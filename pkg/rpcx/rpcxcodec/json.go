package rpcxcodec

import (
	"fmt"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type jsonCodec struct {
	marshalOptions   *protojson.MarshalOptions
	unmarshalOptions *protojson.UnmarshalOptions
}

func (s *jsonCodec) Name() string {
	return "json"
}

func (s *jsonCodec) Marshal(message any) ([]byte, error) {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return nil, errNotProto(message)
	}

	return s.marshalOptions.Marshal(protoMessage)
}

func (s *jsonCodec) Unmarshal(binary []byte, message any) error {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return errNotProto(message)
	}

	return s.unmarshalOptions.Unmarshal(binary, protoMessage)
}

func NewJsonCodec() connect.Codec {
	return &jsonCodec{
		marshalOptions: &protojson.MarshalOptions{
			UseProtoNames: true,
		},
		unmarshalOptions: &protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}
}

func errNotProto(message any) error {
	return fmt.Errorf("%T doesn't implement proto.Message", message)
}
