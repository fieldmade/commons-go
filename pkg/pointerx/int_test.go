package pointerx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	p := Int(1)
	assert.NotNil(t, p)

	v1 := IntR(p)
	assert.Equal(t, 1, v1)

	v2 := IntR(nil)
	assert.Equal(t, 0, v2)
}

func TestInt32(t *testing.T) {
	p := Int32(1)
	assert.NotNil(t, p)

	v1 := Int32R(p)
	assert.Equal(t, int32(1), v1)

	v2 := Int32R(nil)
	assert.Equal(t, int32(0), v2)
}

func TestInt64(t *testing.T) {
	p := Int64(1)
	assert.NotNil(t, p)

	v1 := Int64R(p)
	assert.Equal(t, int64(1), v1)

	v2 := Int64R(nil)
	assert.Equal(t, int64(0), v2)
}
