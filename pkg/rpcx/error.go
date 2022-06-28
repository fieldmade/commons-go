package rpcx

import (
	"fmt"
	"github.com/bufbuild/connect-go"
)

func NewErrorf(code connect.Code, format string, args ...interface{}) *connect.Error {
	return connect.NewError(code, fmt.Errorf(format, args...))
}

func CanceledError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeCanceled, format, args...)
}

func UnknownError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeUnknown, format, args...)
}

func InvalidArgumentError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeInvalidArgument, format, args...)
}

func DeadlineExceededError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeDeadlineExceeded, format, args...)
}

func NotFoundError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeNotFound, format, args...)
}

func AlreadyExistsError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeAlreadyExists, format, args...)
}

func PermissionDeniedError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodePermissionDenied, format, args...)
}

func ResourceExhaustedError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeResourceExhausted, format, args...)
}

func FailedPreconditionError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeFailedPrecondition, format, args...)
}

func AbortedError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeAborted, format, args...)
}

func OutOfRangeError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeOutOfRange, format, args...)
}

func UnimplementedError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeUnimplemented, format, args...)
}

func InternalError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeInternal, format, args...)
}

func UnavailableError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeUnavailable, format, args...)
}

func DataLossError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeDataLoss, format, args...)
}

func UnauthenticatedError(format string, args ...interface{}) *connect.Error {
	return NewErrorf(connect.CodeUnauthenticated, format, args...)
}

func IsCanceledError(err error) bool {
	return HasErrorCode(err, connect.CodeCanceled)
}

func IsUnknownError(err error) bool {
	return HasErrorCode(err, connect.CodeUnknown)
}

func IsInvalidArgumentError(err error) bool {
	return HasErrorCode(err, connect.CodeInvalidArgument)
}

func IsDeadlineExceededError(err error) bool {
	return HasErrorCode(err, connect.CodeDeadlineExceeded)
}

func IsNotFoundError(err error) bool {
	return HasErrorCode(err, connect.CodeNotFound)
}

func IsAlreadyExistsError(err error) bool {
	return HasErrorCode(err, connect.CodeAlreadyExists)
}

func IsPermissionDeniedError(err error) bool {
	return HasErrorCode(err, connect.CodePermissionDenied)
}

func IsResourceExhaustedError(err error) bool {
	return HasErrorCode(err, connect.CodeResourceExhausted)
}

func IsFailedPreconditionError(err error) bool {
	return HasErrorCode(err, connect.CodeFailedPrecondition)
}

func IsAbortedError(err error) bool {
	return HasErrorCode(err, connect.CodeAborted)
}

func IsOutOfRangeError(err error) bool {
	return HasErrorCode(err, connect.CodeOutOfRange)
}

func IsUnimplementedError(err error) bool {
	return HasErrorCode(err, connect.CodeUnimplemented)
}

func IsInternalError(err error) bool {
	return HasErrorCode(err, connect.CodeInternal)
}

func IsUnavailableError(err error) bool {
	return HasErrorCode(err, connect.CodeUnavailable)
}

func IsDataLossError(err error) bool {
	return HasErrorCode(err, connect.CodeDataLoss)
}

func IsUnauthenticatedError(err error) bool {
	return HasErrorCode(err, connect.CodeUnauthenticated)
}

func HasErrorCode(err error, code connect.Code) bool {
	if err == nil {
		return false
	}

	st, ok := err.(*connect.Error)
	if !ok {
		return false
	}

	return st.Code() == code
}
