package relxstore

import (
	"context"
	"github.com/fieldmade/commons-go/pkg/queryx"
	"github.com/fieldmade/commons-go/pkg/relx/relxtranspile"
	"github.com/go-rel/rel"
)

type TranspilerStore[Entity interface{}] interface {
	EntityStore[Entity]
	FindManyQuery(ctx context.Context, query *queryx.Query, filters ...rel.Querier) (*Entities[Entity], error)
}

func NewTranspilerStore[Entity interface{}](
	db rel.Repository,
	transpiler *relxtranspile.Transpiler,
) TranspilerStore[Entity] {
	return &transpilerStoreImpl[Entity]{
		EntityStore: NewEntityStore[Entity](db),
		Transpiler:  transpiler,
	}
}
