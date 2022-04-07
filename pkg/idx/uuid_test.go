package idx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUuid(t *testing.T) {
	id := GenerateUuid()
	assert.Equal(t, 36, len(id))
}

func TestGenerateUuid32(t *testing.T) {
	id := GenerateUuid32()
	assert.Equal(t, 32, len(id))
}

func TestGenerateUuidBase32(t *testing.T) {
	id := GenerateUuidBase32()
	assert.Equal(t, 26, len(id))
}
