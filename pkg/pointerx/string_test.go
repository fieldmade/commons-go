package pointerx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	p := String("abc")
	assert.NotNil(t, p)

	v1 := StringR(p)
	assert.Equal(t, "abc", v1)

	v2 := StringR(nil)
	assert.Equal(t, "", v2)
}
