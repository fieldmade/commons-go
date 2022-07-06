package relxstore

import (
	"context"
	"github.com/go-rel/rel"
)

type Entities[Entity interface{}] struct {
	Items   []*Entity
	HasMore bool
}

type EntityStore[Entity interface{}] interface {
	FindOne(ctx context.Context, filters ...rel.Querier) (*Entity, error)
	FindMany(ctx context.Context, page, pageSize int, filters ...rel.Querier) (*Entities[Entity], error)
}

func NewEntityStore[Entity interface{}](
	db rel.Repository,
) EntityStore[Entity] {
	return &entityStoreImpl[Entity]{
		db: db,
	}
}
