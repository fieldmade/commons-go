package queryx

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type OrderBy struct {
	Fields []*OrderByField
}

type OrderByField struct {
	Field string
	Desc  bool
}

type OrderByRequest interface {
	GetOrderBy() string
}

type orderingParser struct {
	validFields []string
}

func (s *orderingParser) parsingError(msg string) error {
	return fmt.Errorf("error parsing order: %s", msg)
}

func (s *orderingParser) parseOrderBy(str string) (*OrderBy, error) {
	var fields []*OrderByField

	for _, strField := range strings.Split(str, ",") {
		field, err := s.parseOrderByField(strField)
		if err != nil {
			return nil, err
		}

		if field != nil {
			fields = append(fields, field)
		}
	}

	if len(fields) == 0 {
		return nil, nil
	}

	return &OrderBy{
		Fields: fields,
	}, nil
}

func (s *orderingParser) parseOrderByField(str string) (*OrderByField, error) {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return nil, nil
	}

	split := strings.Split(str, " ")
	if len(split) == 1 {
		return s.createOrderByField(split[0], false)
	}

	if len(split) > 2 {
		return nil, s.parsingError("invalid syntax")
	}

	op := strings.ToLower(strings.TrimSpace(split[1]))
	if op == "desc" {
		return s.createOrderByField(split[0], true)
	} else if op == "asc" {
		return s.createOrderByField(split[0], false)
	}

	return nil, s.parsingError("invalid syntax")
}

func (s *orderingParser) createOrderByField(field string, desc bool) (*OrderByField, error) {
	field = strings.TrimSpace(field)

	if !lo.Contains(s.validFields, field) {
		return nil, s.parsingError(fmt.Sprintf("unknown field [%s]", field))
	}

	return &OrderByField{
		Field: field,
		Desc:  desc,
	}, nil
}
