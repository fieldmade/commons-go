package queryx

import "math"

type Pagination struct {
	Page     int32
	PageSize int32
}

type PaginationRequest interface {
	GetPage() int32
	GetPageSize() int32
}

type paginationBuilder struct {
	maxPage         int32
	maxPageSize     int32
	defaultPageSize int32
}

func (s *paginationBuilder) init() {
	if s.maxPage <= 0 {
		s.maxPage = math.MaxInt32
	}

	if s.maxPageSize <= 0 {
		s.maxPageSize = 50
	}

	if s.defaultPageSize <= 0 {
		s.defaultPageSize = 10
	}
}

func (s *paginationBuilder) buildPagination(page, pageSize int32) *Pagination {
	if page > s.maxPage {
		page = s.maxPage
	}

	if page <= 0 {
		page = 1
	}

	if pageSize > s.maxPageSize {
		pageSize = s.maxPageSize
	}

	if pageSize <= 0 {
		pageSize = s.defaultPageSize
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}
