package queryx

import (
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
)

type Query struct {
	Pagination *Pagination
	OrderBy    *ordering.OrderBy
	Filter     *filtering.Filter
}

func (s *Query) GetPage() int32 {
	if s.Pagination != nil {
		return s.Pagination.Page
	} else {
		return 1
	}
}

func (s *Query) GetPageSize() int32 {
	if s.Pagination != nil {
		return s.Pagination.PageSize
	} else {
		return 10
	}
}

type QueryDefinition struct {
	MaxPage         int32
	MaxPageSize     int32
	DefaultPageSize int32
	OrderFields     []string
	FilterFields    FilterFields
}

func (s *QueryDefinition) ParseQuery(req interface{}) (*Query, error) {
	pagination, err := buildPaginationIfNeeded(s, req)
	if err != nil {
		return nil, err
	}

	orderBy, err := parseOrderByIfNeeded(s, req)
	if err != nil {
		return nil, err
	}

	filter, err := parseFilterIfNeeded(s, req)
	if err != nil {
		return nil, err
	}

	return &Query{
		Pagination: pagination,
		OrderBy:    orderBy,
		Filter:     filter,
	}, nil
}
