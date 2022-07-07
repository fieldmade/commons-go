package queryx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseOrderBy(t *testing.T) {
	parser := &orderingParser{
		validFields: []string{"a", "b", "c"},
	}

	orderBy, err := parser.parseOrderBy("a, b asc, c desc")

	assert.NoError(t, err)
	assert.NotNil(t, orderBy)
	assert.Equal(t, 3, len(orderBy.Fields))
}
