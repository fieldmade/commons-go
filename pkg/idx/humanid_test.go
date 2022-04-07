package idx

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateHumanId(t *testing.T) {
	id := GenerateHumanId()
	assert.Equal(t, 9, len(id))
	assert.Equal(t, 1, strings.Count(id, "-"))
}
