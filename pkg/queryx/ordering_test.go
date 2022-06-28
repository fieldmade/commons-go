package queryx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockOrderingRequest struct {
	orderBy string
}

func (s *mockOrderingRequest) GetOrderBy() string {
	return s.orderBy
}

func Test_parseOrderBy(t *testing.T) {
	def := &QueryDefinition{
		OrderFields: []string{"a", "b", "c"},
	}

	req := &mockOrderingRequest{
		orderBy: "a asc, b desc, c",
	}

	orderBy, err := parseOrderByIfNeeded(def, req)

	assert.NoError(t, err)
	assert.NotNil(t, orderBy)
	assert.Equal(t, 3, len(orderBy.Fields))
}
