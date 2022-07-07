package queryx

import (
	"github.com/stretchr/testify/assert"
	"go.einride.tech/aip/filtering"
	"testing"
)

func Test_parseFilter(t *testing.T) {
	parser := &filteringParser{
		fields: FilterFields{
			"a": filtering.TypeString,
			"b": filtering.TypeInt,
		},
	}

	filterExpr, err := parser.parseFilter("a = 'one' AND b = 1")

	assert.NoError(t, err)
	assert.NotNil(t, filterExpr)
	assert.NotNil(t, filterExpr.CheckedExpr)
}
