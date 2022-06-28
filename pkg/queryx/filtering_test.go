package queryx

import (
	"github.com/stretchr/testify/assert"
	"go.einride.tech/aip/filtering"
	"testing"
)

type mockFilterRequest struct {
	filter string
}

func (s *mockFilterRequest) GetFilter() string {
	return s.filter
}

func Test_parseFilter(t *testing.T) {
	def := &QueryDefinition{
		FilterFields: []*FilterField{
			{
				Field: "a",
				Type:  filtering.TypeString,
			},
			{
				Field: "b",
				Type:  filtering.TypeInt,
			},
		},
	}

	req := &mockFilterRequest{
		filter: "a = 'one' AND b = 1",
	}

	filterExpr, err := parseFilterIfNeeded(def, req)

	assert.NoError(t, err)
	assert.NotNil(t, filterExpr)
	assert.NotNil(t, filterExpr.CheckedExpr)
}
