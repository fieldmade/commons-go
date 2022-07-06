package queryx

import (
	"fmt"
	"go.einride.tech/aip/filtering"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var (
	TypeString    = filtering.TypeString
	TypeInt       = filtering.TypeInt
	TypeFloat     = filtering.TypeFloat
	TypeBool      = filtering.TypeBool
	TypeDuration  = filtering.TypeDuration
	TypeTimestamp = filtering.TypeTimestamp
)

type FilterFields map[string]*expr.Type

func filterParsingError(err error) error {
	return fmt.Errorf("error parsing filter: %s", err.Error())
}

func filterDeclarationOptions(def *QueryDefinition) []filtering.DeclarationOption {
	res := []filtering.DeclarationOption{
		filtering.DeclareStandardFunctions(),
	}

	for field, fieldType := range def.FilterFields {
		res = append(res, filtering.DeclareIdent(field, fieldType))
	}

	return res
}

func parseFilter(def *QueryDefinition, req filtering.Request) (*filtering.Filter, error) {
	opts := filterDeclarationOptions(def)
	decl, err := filtering.NewDeclarations(opts...)
	if err != nil {
		return nil, filterParsingError(err)
	}

	res, err := filtering.ParseFilter(req, decl)
	if err != nil {
		return nil, filterParsingError(err)
	}

	return &res, nil
}

func parseFilterIfNeeded(def *QueryDefinition, req interface{}) (*filtering.Filter, error) {
	typedReq, ok := req.(filtering.Request)
	if !ok {
		return nil, nil
	}

	return parseFilter(def, typedReq)
}
