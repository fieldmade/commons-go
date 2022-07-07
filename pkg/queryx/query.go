package queryx

import (
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
)

type Query struct {
	Pagination *Pagination
	OrderBy    *OrderBy
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
	orderBy, err := s.parseOrderByIfNeeded(req)
	if err != nil {
		return nil, err
	}

	filter, err := s.parseFilterIfNeeded(req)
	if err != nil {
		return nil, err
	}

	return &Query{
		Pagination: s.buildPaginationIfNeeded(req),
		OrderBy:    orderBy,
		Filter:     filter,
	}, nil
}

func (s *QueryDefinition) buildPaginationIfNeeded(req interface{}) *Pagination {
	typedReq, ok := req.(PaginationRequest)
	if !ok {
		return nil
	}

	builder := paginationBuilder{
		maxPage:         s.MaxPage,
		maxPageSize:     s.MaxPageSize,
		defaultPageSize: s.DefaultPageSize,
	}

	builder.init()
	return builder.buildPagination(typedReq.GetPage(), typedReq.GetPageSize())
}

func (s *QueryDefinition) parseOrderByIfNeeded(req interface{}) (*OrderBy, error) {
	typedReq, ok := req.(ordering.Request)
	if !ok {
		return nil, nil
	}

	parser := orderingParser{
		validFields: s.OrderFields,
	}

	return parser.parseOrderBy(typedReq.GetOrderBy())
}

func (s *QueryDefinition) parseFilterIfNeeded(req interface{}) (*filtering.Filter, error) {
	typedReq, ok := req.(filtering.Request)
	if !ok {
		return nil, nil
	}

	parser := filteringParser{
		fields: s.FilterFields,
	}

	return parser.parseFilter(typedReq.GetFilter())
}
