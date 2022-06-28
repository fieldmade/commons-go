package errorx

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrappedErrors(t *testing.T) {
	tests := []struct {
		err       error
		actualErr error
	}{
		{
			err:       NewErrIllegalArgument("some %s", "message"),
			actualErr: ErrIllegalArgument,
		},
		{
			err:       NewErrUnauthenticated("some %s", "message"),
			actualErr: ErrUnauthenticated,
		},
		{
			err:       NewErrPermissionDenied("some %s", "message"),
			actualErr: ErrPermissionDenied,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(errors.Unwrap(test.err).Error()), func(t *testing.T) {
			assert.True(t, errors.Is(test.err, test.actualErr))
			assert.Equal(t, "some message", test.err.Error())
		})
	}
}
