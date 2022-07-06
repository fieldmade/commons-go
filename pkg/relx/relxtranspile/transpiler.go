package relxtranspile

import (
	"github.com/fieldmade/commons-go/pkg/queryx"
	"github.com/go-rel/rel"
)

type Transpiler struct {
	FieldMap map[string]string
}

func (s *Transpiler) fieldName(name string) string {
	val, ok := s.FieldMap[name]
	if ok {
		return val
	} else {
		return name
	}
}

func (s *Transpiler) TranspileQuery(query *queryx.Query) ([]rel.Querier, error) {
	var res []rel.Querier

	if query.OrderBy != nil {
		res = append(res, s.transpileOrderBy(query.OrderBy)...)
	}

	if query.Filter != nil {
		filterRes, err := s.transpileFilter(query.Filter)
		if err != nil {
			return nil, err
		}

		if filterRes != nil {
			res = append(res, filterRes)
		}
	}

	return res, nil
}
