package idx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUniqueId(t *testing.T) {
	id := GenerateUniqueId()
	assert.Equal(t, 32, len(id))
}

func TestGenerateUniqueToken(t *testing.T) {
	id := GenerateUniqueToken()
	assert.Equal(t, 64, len(id))
}
