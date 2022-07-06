package relxstore

import (
	"context"
	"github.com/fieldmade/commons-go/pkg/queryx"
	"github.com/fieldmade/commons-go/pkg/relx/relxtranspile"
	"github.com/go-rel/rel"
)

type transpilerStoreImpl[Entity interface{}] struct {
	EntityStore[Entity]
	Transpiler *relxtranspile.Transpiler
}

func (s *transpilerStoreImpl[Entity]) FindManyQuery(ctx context.Context, query *queryx.Query, filters ...rel.Querier) (*Entities[Entity], error) {
	queryFilters, err := s.Transpiler.TranspileQuery(query)
	if err != nil {
		return nil, err
	}

	allFilters := append(filters, queryFilters...)
	return s.FindMany(ctx, int(query.GetPage()), int(query.GetPageSize()), allFilters...)
}
