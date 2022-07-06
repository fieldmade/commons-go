package relxstore

import (
	"context"
	"errors"
	"github.com/go-rel/rel"
)

type entityStoreImpl[Entity interface{}] struct {
	db rel.Repository
}

func (s *entityStoreImpl[Entity]) FindOne(ctx context.Context, filters ...rel.Querier) (*Entity, error) {
	var res *Entity

	err := s.db.Find(ctx, &res, filters...)
	if errors.Is(err, rel.ErrNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *entityStoreImpl[Entity]) FindMany(ctx context.Context, page, pageSize int, filters ...rel.Querier) (*Entities[Entity], error) {
	filters = append(filters, rel.Offset((page-1)*pageSize))
	filters = append(filters, rel.Limit(pageSize+1))

	var res []*Entity
	err := s.db.FindAll(ctx, &res, filters...)
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
