package rpcxerr

import (
	"context"
	"errors"
	"github.com/bufbuild/connect-go"
	"github.com/fieldmade/commons-go/pkg/errorx"
)

func translateError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, errorx.ErrIllegalArgument) {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	if errors.Is(err, errorx.ErrUnauthenticated) {
		return connect.NewError(connect.CodeUnauthenticated, err)
	}

	if errors.Is(err, errorx.ErrPermissionDenied) {
		return connect.NewError(connect.CodePermissionDenied, err)
	}

	return err
}

func NewServerInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			return res, translateError(err)
		}
	}
}
