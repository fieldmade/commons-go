package queryx

import (
	"fmt"
	"go.einride.tech/aip/filtering"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"strings"
)

type FilterFields map[string]*expr.Type

type filteringParser struct {
	fields FilterFields
}

func (s *filteringParser) parsingError(err error) error {
	return fmt.Errorf("error parsing filter: %s", err.Error())
}

func (s *filteringParser) filterDeclarationOptions() []filtering.DeclarationOption {
	res := []filtering.DeclarationOption{
		filtering.DeclareStandardFunctions(),
		// filtering.DeclareIdent("true", ...)
		// filtering.DeclareIdent("false", ...)
	}

	for field, fieldType := range s.fields {
		res = append(res, filtering.DeclareIdent(field, fieldType))
	}

	return res
}

func (s *filteringParser) parseFilter(str string) (*filtering.Filter, error) {
	str = strings.TrimSpace(str)

	if str == "" {
		return nil, nil
	}

	filterParser := &filtering.Parser{}
	filterParser.Init(str)

	parsedExpr, err := filterParser.Parse()
	if err != nil {
		return nil, err
	}

	opts := s.filterDeclarationOptions()

	decl, err := filtering.NewDeclarations(opts...)
	if err != nil {
		return nil, s.parsingError(err)
	}

	checker := &filtering.Checker{}
	checker.Init(parsedExpr.Expr, parsedExpr.SourceInfo, decl)

	checkedExpr, err := checker.Check()
	if err != nil {
		return nil, err
	}

	return &filtering.Filter{
		CheckedExpr: checkedExpr,
	}, nil
}
