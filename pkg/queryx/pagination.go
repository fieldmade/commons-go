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

func buildPagination(def *QueryDefinition, req PaginationRequest) (*Pagination, error) {
	maxPage := def.MaxPage
	if maxPage <= 0 {
		maxPage = math.MaxInt32
	}

	maxPageSize := def.MaxPageSize
	if maxPageSize <= 0 {
		maxPageSize = 50
	}

	defaultPageSize := def.DefaultPageSize
	if defaultPageSize <= 0 {
		defaultPageSize = 10
	}

	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page > maxPage {
		page = maxPage
	}

	if page <= 0 {
		page = 1
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func buildPaginationIfNeeded(def *QueryDefinition, req interface{}) (*Pagination, error) {
	typedReq, ok := req.(PaginationRequest)
	if !ok {
		return nil, nil
	}

	return buildPagination(def, typedReq)
}
