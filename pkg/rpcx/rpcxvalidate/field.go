package rpcxvalidate

import (
	"github.com/bufbuild/connect-go"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/protobuf/proto"
)

func toProtoMessage(req interface{}) proto.Message {
	return req.(proto.Message)
}

func validateFieldBehaviour(req proto.Message) error {
	err := fieldbehavior.ValidateRequiredFields(req)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	} else {
		return nil
	}
}
