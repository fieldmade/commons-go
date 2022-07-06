package relxdb

import (
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"
)

func NewPostgresDb(dsn string, logging bool) (rel.Repository, error) {
	adapter, err := postgres.Open(dsn)
	if err != nil {
		return nil, err
	}

	return rel.New(adapter), nil
}
