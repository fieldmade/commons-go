package queryx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildPagination(t *testing.T) {
	tests := []struct {
		page           int32
		pageSize       int32
		actualPage     int32
		actualPageSize int32
	}{
		{
			page:           -1,
			pageSize:       -1,
			actualPage:     1,
			actualPageSize: 20,
		},
		{
			page:           0,
			pageSize:       0,
			actualPage:     1,
			actualPageSize: 20,
		},
		{
			page:           5,
			pageSize:       30,
			actualPage:     5,
			actualPageSize: 30,
		},
		{
			page:           5,
			pageSize:       120,
			actualPage:     5,
			actualPageSize: 100,
		},
	}

	builder := &paginationBuilder{
		defaultPageSize: 20,
		maxPageSize:     100,
	}

	builder.init()

	for _, test := range tests {
		name := fmt.Sprintf("page=%v:pageSze=%v", test.page, test.pageSize)

		t.Run(name, func(t *testing.T) {
			pagination := builder.buildPagination(test.page, test.pageSize)

			assert.NotNil(t, pagination)
			assert.Equal(t, test.actualPage, pagination.Page)
			assert.Equal(t, test.actualPageSize, pagination.PageSize)
		})
	}
}
