package rpcx

import "fmt"

const (
	ErrorCodeOK                 ErrorCode = 0
	ErrorCodeCanceled           ErrorCode = 1
	ErrorCodeUnknown            ErrorCode = 2
	ErrorCodeInvalidArgument    ErrorCode = 3
	ErrorCodeDeadlineExceeded   ErrorCode = 4
	ErrorCodeNotFound           ErrorCode = 5
	ErrorCodeAlreadyExists      ErrorCode = 6
	ErrorCodePermissionDenied   ErrorCode = 7
	ErrorCodeResourceExhausted  ErrorCode = 8
	ErrorCodeFailedPrecondition ErrorCode = 9
	ErrorCodeAborted            ErrorCode = 10
	ErrorCodeOutOfRange         ErrorCode = 11
	ErrorCodeUnimplemented      ErrorCode = 12
	ErrorCodeInternal           ErrorCode = 13
	ErrorCodeUnavailable        ErrorCode = 14
	ErrorCodeDataLoss           ErrorCode = 15
	ErrorCodeUnauthenticated    ErrorCode = 16
)

type ErrorCode int

type rpcError struct {
	code    ErrorCode
	message string
}

func (s *rpcError) Error() string {
	return s.message
}

func newRpcError(code ErrorCode, format string, args ...interface{}) *rpcError {
	return &rpcError{
		code:    code,
		message: fmt.Sprintf(format, args...),
	}
}

func ToErrorCode(err error) ErrorCode {
	if err == nil {
		return ErrorCodeOK
	}

	detailed, ok := err.(*rpcError)
	if !ok {
		return ErrorCodeUnknown
	}

	return detailed.code
}

func HasErrorCode(err error, code ErrorCode) bool {
	return ToErrorCode(err) == code
}

func CanceledError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeCanceled, format, args...)
}

func UnknownError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeUnknown, format, args...)
}

func InvalidArgumentError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeInvalidArgument, format, args...)
}

func DeadlineExceededError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeDeadlineExceeded, format, args...)
}

func NotFoundError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeNotFound, format, args...)
}

func AlreadyExistsError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeAlreadyExists, format, args...)
}

func PermissionDeniedError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodePermissionDenied, format, args...)
}

func ResourceExhaustedError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeResourceExhausted, format, args...)
}

func FailedPreconditionError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeFailedPrecondition, format, args...)
}

func AbortedError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeAborted, format, args...)
}

func OutOfRangeError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeOutOfRange, format, args...)
}

func UnimplementedError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeUnimplemented, format, args...)
}

func InternalError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeInternal, format, args...)
}

func UnavailableError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeUnavailable, format, args...)
}

func DataLossError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeDataLoss, format, args...)
}

func UnauthenticatedError(format string, args ...interface{}) error {
	return newRpcError(ErrorCodeUnauthenticated, format, args...)
}

func IsCanceledError(err error) bool {
	return HasErrorCode(err, ErrorCodeCanceled)
}

func IsUnknownError(err error) bool {
	return HasErrorCode(err, ErrorCodeUnknown)
}

func IsInvalidArgumentError(err error) bool {
	return HasErrorCode(err, ErrorCodeInvalidArgument)
}

func IsDeadlineExceededError(err error) bool {
	return HasErrorCode(err, ErrorCodeDeadlineExceeded)
}

func IsNotFoundError(err error) bool {
	return HasErrorCode(err, ErrorCodeNotFound)
}

func IsAlreadyExistsError(err error) bool {
	return HasErrorCode(err, ErrorCodeAlreadyExists)
}

func IsPermissionDeniedError(err error) bool {
	return HasErrorCode(err, ErrorCodePermissionDenied)
}

func IsResourceExhaustedError(err error) bool {
	return HasErrorCode(err, ErrorCodeResourceExhausted)
}

func IsFailedPreconditionError(err error) bool {
	return HasErrorCode(err, ErrorCodeFailedPrecondition)
}

func IsAbortedError(err error) bool {
	return HasErrorCode(err, ErrorCodeAborted)
}

func IsOutOfRangeError(err error) bool {
	return HasErrorCode(err, ErrorCodeOutOfRange)
}

func IsUnimplementedError(err error) bool {
	return HasErrorCode(err, ErrorCodeUnimplemented)
}

func IsInternalError(err error) bool {
	return HasErrorCode(err, ErrorCodeInternal)
}

func IsUnavailableError(err error) bool {
	return HasErrorCode(err, ErrorCodeUnavailable)
}

func IsDataLossError(err error) bool {
	return HasErrorCode(err, ErrorCodeDataLoss)
}

func IsUnauthenticatedError(err error) bool {
	return HasErrorCode(err, ErrorCodeUnauthenticated)
}
