package rpcx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		err   error
		code  ErrorCode
		check func(err error) bool
	}{
		{
			err:   CanceledError("some %s", "message"),
			code:  ErrorCodeCanceled,
			check: IsCanceledError,
		},
		{
			err:   UnknownError("some %s", "message"),
			code:  ErrorCodeUnknown,
			check: IsUnknownError,
		},
		{
			err:   InvalidArgumentError("some %s", "message"),
			code:  ErrorCodeInvalidArgument,
			check: IsInvalidArgumentError,
		},
		{
			err:   DeadlineExceededError("some %s", "message"),
			code:  ErrorCodeDeadlineExceeded,
			check: IsDeadlineExceededError,
		},
		{
			err:   NotFoundError("some %s", "message"),
			code:  ErrorCodeNotFound,
			check: IsNotFoundError,
		},
		{
			err:   AlreadyExistsError("some %s", "message"),
			code:  ErrorCodeAlreadyExists,
			check: IsAlreadyExistsError,
		},
		{
			err:   PermissionDeniedError("some %s", "message"),
			code:  ErrorCodePermissionDenied,
			check: IsPermissionDeniedError,
		},
		{
			err:   ResourceExhaustedError("some %s", "message"),
			code:  ErrorCodeResourceExhausted,
			check: IsResourceExhaustedError,
		},
		{
			err:   FailedPreconditionError("some %s", "message"),
			code:  ErrorCodeFailedPrecondition,
			check: IsFailedPreconditionError,
		},
		{
			err:   AbortedError("some %s", "message"),
			code:  ErrorCodeAborted,
			check: IsAbortedError,
		},
		{
			err:   OutOfRangeError("some %s", "message"),
			code:  ErrorCodeOutOfRange,
			check: IsOutOfRangeError,
		},
		{
			err:   UnimplementedError("some %s", "message"),
			code:  ErrorCodeUnimplemented,
			check: IsUnimplementedError,
		},
		{
			err:   InternalError("some %s", "message"),
			code:  ErrorCodeInternal,
			check: IsInternalError,
		},
		{
			err:   UnavailableError("some %s", "message"),
			code:  ErrorCodeUnavailable,
			check: IsUnavailableError,
		},
		{
			err:   DataLossError("some %s", "message"),
			code:  ErrorCodeDataLoss,
			check: IsDataLossError,
		},
		{
			err:   UnauthenticatedError("some %s", "message"),
			code:  ErrorCodeUnauthenticated,
			check: IsUnauthenticatedError,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("code_%v", test.code), func(t *testing.T) {
			assert.Equal(t, test.code, ToErrorCode(test.err))
			assert.Equal(t, "some message", test.err.Error())
			assert.False(t, test.check(nil))
			assert.True(t, test.check(test.err))
		})
	}
}
