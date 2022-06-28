package rpcxvalidate

import (
	"context"
	"github.com/bufbuild/connect-go"
)

func NewServerInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			err := validateFieldBehaviour(toProtoMessage(req.Any()))
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}
