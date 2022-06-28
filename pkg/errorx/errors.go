package errorx

import (
	"github.com/pkg/errors"
)

var (
	ErrIllegalArgument  = errors.New("illegal_argument")
	ErrUnauthenticated  = errors.New("unauthenticated")
	ErrPermissionDenied = errors.New("permission_denied")
)

func NewErrIllegalArgument(format string, args ...interface{}) error {
	return WrapError(ErrIllegalArgument, format, args...)
}

func NewErrUnauthenticated(format string, args ...interface{}) error {
	return WrapError(ErrUnauthenticated, format, args...)
}

func NewErrPermissionDenied(format string, args ...interface{}) error {
	return WrapError(ErrPermissionDenied, format, args...)
}
