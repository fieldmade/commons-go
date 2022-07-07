package queryx

import (
	"github.com/stretchr/testify/assert"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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
		FilterFields: map[string]*expr.Type{
			"a": TypeString,
			"b": TypeInt,
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
