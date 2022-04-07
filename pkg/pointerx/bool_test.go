package pointerx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	p := Bool(true)
	assert.NotNil(t, p)

	v1 := BoolR(p)
	assert.Equal(t, true, v1)

	v2 := BoolR(nil)
	assert.Equal(t, false, v2)
}
