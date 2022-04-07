package pointerx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32(t *testing.T) {
	p := Float32(1.2)
	assert.NotNil(t, p)

	v1 := Float32R(p)
	assert.Equal(t, float32(1.2), v1)

	v2 := Float32R(nil)
	assert.Equal(t, float32(0), v2)
}

func TestFloat64(t *testing.T) {
	p := Float64(1.2)
	assert.NotNil(t, p)

	v1 := Float64R(p)
	assert.Equal(t, 1.2, v1)

	v2 := Float64R(nil)
	assert.Equal(t, float64(0), v2)
}
