package rpcx

import (
	"errors"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		err   *connect.Error
		code  connect.Code
		check func(err error) bool
	}{
		{
			err:   CanceledError("some %s", "message"),
			code:  connect.CodeCanceled,
			check: IsCanceledError,
		},
		{
			err:   UnknownError("some %s", "message"),
			code:  connect.CodeUnknown,
			check: IsUnknownError,
		},
		{
			err:   InvalidArgumentError("some %s", "message"),
			code:  connect.CodeInvalidArgument,
			check: IsInvalidArgumentError,
		},
		{
			err:   DeadlineExceededError("some %s", "message"),
			code:  connect.CodeDeadlineExceeded,
			check: IsDeadlineExceededError,
		},
		{
			err:   NotFoundError("some %s", "message"),
			code:  connect.CodeNotFound,
			check: IsNotFoundError,
		},
		{
			err:   AlreadyExistsError("some %s", "message"),
			code:  connect.CodeAlreadyExists,
			check: IsAlreadyExistsError,
		},
		{
			err:   PermissionDeniedError("some %s", "message"),
			code:  connect.CodePermissionDenied,
			check: IsPermissionDeniedError,
		},
		{
			err:   ResourceExhaustedError("some %s", "message"),
			code:  connect.CodeResourceExhausted,
			check: IsResourceExhaustedError,
		},
		{
			err:   FailedPreconditionError("some %s", "message"),
			code:  connect.CodeFailedPrecondition,
			check: IsFailedPreconditionError,
		},
		{
			err:   AbortedError("some %s", "message"),
			code:  connect.CodeAborted,
			check: IsAbortedError,
		},
		{
			err:   OutOfRangeError("some %s", "message"),
			code:  connect.CodeOutOfRange,
			check: IsOutOfRangeError,
		},
		{
			err:   UnimplementedError("some %s", "message"),
			code:  connect.CodeUnimplemented,
			check: IsUnimplementedError,
		},
		{
			err:   InternalError("some %s", "message"),
			code:  connect.CodeInternal,
			check: IsInternalError,
		},
		{
			err:   UnavailableError("some %s", "message"),
			code:  connect.CodeUnavailable,
			check: IsUnavailableError,
		},
		{
			err:   DataLossError("some %s", "message"),
			code:  connect.CodeDataLoss,
			check: IsDataLossError,
		},
		{
			err:   UnauthenticatedError("some %s", "message"),
			code:  connect.CodeUnauthenticated,
			check: IsUnauthenticatedError,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.code), func(t *testing.T) {
			assert.Equal(t, test.code, test.err.Code())
			assert.Equal(t, "some message", test.err.Message())
			assert.False(t, test.check(nil))
			assert.False(t, test.check(errors.New("other")))
			assert.True(t, test.check(test.err))
		})
	}
}
