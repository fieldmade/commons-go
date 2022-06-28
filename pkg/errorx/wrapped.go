package errorx

import (
	"fmt"
)

type wrappedError struct {
	message string
	err     error
}

func (s *wrappedError) Error() string {
	return s.message
}

func (s *wrappedError) Unwrap() error {
	return s.err
}

func WrapError(err error, format string, args ...interface{}) error {
	return &wrappedError{
		message: fmt.Sprintf(format, args...),
		err:     err,
	}
}
