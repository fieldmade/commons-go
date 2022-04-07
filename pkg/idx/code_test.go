package idx

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	t.Run("generate-3", func(t *testing.T) {
		code := GenerateCode(3)
		assert.Equal(t, 3, len(code))
		assert.Equal(t, 0, strings.Count(code, "-"))
	})

	t.Run("generate-3-3", func(t *testing.T) {
		code := GenerateCode(3, 3)
		assert.Equal(t, 7, len(code))
		assert.Equal(t, 1, strings.Count(code, "-"))
	})

	t.Run("generate-3-3-3", func(t *testing.T) {
		code := GenerateCode(3, 3, 3)
		assert.Equal(t, 11, len(code))
		assert.Equal(t, 2, strings.Count(code, "-"))
	})

	t.Run("uniqueness", func(t *testing.T) {
		m := map[string]bool{}

		for i := 0; i < 10000; i++ {
			v := GenerateCode(4, 4)
			_, dupe := m[v]

			assert.False(t, dupe)

			m[v] = true
		}
	})
}
