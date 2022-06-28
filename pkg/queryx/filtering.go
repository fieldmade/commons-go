package queryx

import (
	"github.com/fieldmade/commons-go/pkg/errorx"
	"go.einride.tech/aip/filtering"
	"google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type FilterField struct {
	Field string
	Type  *expr.Type
}

func filterParsingError(err error) error {
	return errorx.NewErrIllegalArgument("error parsing filter: %s", err.Error())
}

func filterDeclarationOptions(def *QueryDefinition) []filtering.DeclarationOption {
	res := []filtering.DeclarationOption{
		filtering.DeclareStandardFunctions(),
	}

	for _, field := range def.FilterFields {
		res = append(res, filtering.DeclareIdent(field.Field, field.Type))
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
