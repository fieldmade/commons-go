package relxtranspile

import (
	"github.com/go-rel/rel"
	"go.einride.tech/aip/ordering"
)

func (s *Transpiler) transpileOrderBy(from *ordering.OrderBy) []rel.Querier {
	var res []rel.Querier

	for _, field := range from.Fields {
		res = append(res, s.transpileOrderField(&field))
	}

	return res
}

func (s *Transpiler) transpileOrderField(from *ordering.Field) rel.Querier {
	name := s.fieldName(from.Path)

	if from.Desc {
		return rel.SortDesc(name)
	} else {
		return rel.SortAsc(name)
	}
}
