package queryx

import (
	"github.com/fieldmade/commons-go/pkg/errorx"
	"go.einride.tech/aip/ordering"
)

func orderParsingError(err error) error {
	return errorx.NewErrIllegalArgument("error parsing order: %s", err.Error())
}

func parseOrderBy(def *QueryDefinition, req ordering.Request) (*ordering.OrderBy, error) {
	res, err := ordering.ParseOrderBy(req)
	if err != nil {
		return nil, orderParsingError(err)
	}

	err = res.ValidateForPaths(def.OrderFields...)
	if err != nil {
		return nil, orderParsingError(err)
	}

	return &res, nil
}

func parseOrderByIfNeeded(def *QueryDefinition, req interface{}) (*ordering.OrderBy, error) {
	typedReq, ok := req.(ordering.Request)
	if !ok {
		return nil, nil
	}

	return parseOrderBy(def, typedReq)
}
