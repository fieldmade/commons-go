package relxstore

import (
	"context"
	"errors"
	"github.com/fieldmade/commons-go/pkg/queryx"
	"github.com/fieldmade/commons-go/pkg/relx/relxtranspile"
	"github.com/go-rel/rel"
)

type EntityStore[Entity any] struct {
	Repository rel.Repository
	Transpiler *relxtranspile.Transpiler
}

func (s *EntityStore[Entity]) FindOne(ctx context.Context, filters ...rel.Querier) (*Entity, error) {
	var res *Entity

	err := s.Repository.Find(ctx, res, filters...)
	if errors.Is(err, rel.ErrNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EntityStore[Entity]) FindMany(ctx context.Context, page, pageSize int, filters ...rel.Querier) (*Entities[Entity], error) {
	filters = append(filters, rel.Offset((page-1)*pageSize))
	filters = append(filters, rel.Limit(pageSize+1))

	var res []*Entity
	err := s.Repository.FindAll(ctx, &res, filters...)
	if err != nil {
		return nil, err
	}

	hasMore := false
	if len(res) > pageSize {
		hasMore = true
		res = res[0:pageSize]
	}

	return &Entities[Entity]{
		Items:   res,
		HasMore: hasMore,
	}, nil
}

func (s *EntityStore[Entity]) FindManyQuery(ctx context.Context, query *queryx.Query, filters ...rel.Querier) (*Entities[Entity], error) {
	queryFilters, err := s.Transpiler.TranspileQuery(query)
	if err != nil {
		return nil, err
	}

	allFilters := append(filters, queryFilters...)
	return s.FindMany(ctx, int(query.GetPage()), int(query.GetPageSize()), allFilters...)
}
