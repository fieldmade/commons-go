package relxtranspile

import (
	"fmt"
	"github.com/go-rel/rel"
	"go.einride.tech/aip/filtering"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"time"
)

type compareExprFunc func(field string, value interface{}) rel.FilterQuery

type logicalExprFunc func(inner ...rel.FilterQuery) rel.FilterQuery

func (s *Transpiler) transpileFilter(from *filtering.Filter) (rel.Querier, error) {
	if from.CheckedExpr == nil {
		return nil, nil
	}

	resultExpr, err := s.transpileBoolExpr(from.CheckedExpr.Expr)
	if err != nil {
		return nil, err
	}

	return *resultExpr, nil
}

func (s *Transpiler) transpileBoolExpr(from *expr.Expr) (*rel.FilterQuery, error) {
	kind := from.ExprKind

	switch kind.(type) {
	case *expr.Expr_CallExpr:
		return s.transpileBoolCallExpr(from.GetCallExpr())
	case *expr.Expr_IdentExpr:
		return s.transpileBoolIdentExpr(from.GetIdentExpr())
	default:
		return nil, fmt.Errorf("unsupported expr: %T", kind)
	}
}

func (s *Transpiler) transpileBoolCallExpr(from *expr.Expr_Call) (*rel.FilterQuery, error) {
	switch from.Function {
	case filtering.FunctionEquals:
		return s.transpileComparisonCallExpr(from, rel.Eq)
	case filtering.FunctionNotEquals:
		return s.transpileComparisonCallExpr(from, rel.Ne)
	case filtering.FunctionLessThan:
		return s.transpileComparisonCallExpr(from, rel.Lt)
	case filtering.FunctionLessEquals:
		return s.transpileComparisonCallExpr(from, rel.Lte)
	case filtering.FunctionGreaterThan:
		return s.transpileComparisonCallExpr(from, rel.Gt)
	case filtering.FunctionGreaterEquals:
		return s.transpileComparisonCallExpr(from, rel.Gte)
	case filtering.FunctionOr:
		return s.transpileBinaryLogicalCallExpr(from, rel.Or)
	case filtering.FunctionAnd:
		return s.transpileBinaryLogicalCallExpr(from, rel.And)
	case filtering.FunctionNot:
		return s.transpileNotCallExpr(from)
	default:
		return nil, fmt.Errorf("unsupported function call: %s", from.Function)
	}
}

func (s *Transpiler) transpileBoolIdentExpr(from *expr.Expr_Ident) (*rel.FilterQuery, error) {
	field := from.GetName()

	result := rel.Eq(field, true)
	return &result, nil
}

func (s *Transpiler) transpileComparisonCallExpr(from *expr.Expr_Call, fn compareExprFunc) (*rel.FilterQuery, error) {
	if len(from.Args) != 2 {
		return nil, fmt.Errorf(
			"unexpected number of arguments to `%s`: %d",
			from.GetFunction(),
			len(from.Args),
		)
	}

	lhsExpr, err := s.transpileFieldExpr(from.Args[0])
	if err != nil {
		return nil, err
	}

	rhsExpr, err := s.transpileValueExpr(from.Args[1])
	if err != nil {
		return nil, err
	}

	result := fn(lhsExpr, rhsExpr)
	return &result, nil
}

func (s *Transpiler) transpileBinaryLogicalCallExpr(from *expr.Expr_Call, fn logicalExprFunc) (*rel.FilterQuery, error) {
	if len(from.Args) != 2 {
		return nil, fmt.Errorf(
			"unexpected number of arguments to `%s`: %d",
			from.GetFunction(),
			len(from.Args),
		)
	}

	lhsExpr, err := s.transpileBoolExpr(from.Args[0])
	if err != nil {
		return nil, err
	}

	rhsExpr, err := s.transpileBoolExpr(from.Args[1])
	if err != nil {
		return nil, err
	}

	result := fn(*lhsExpr, *rhsExpr)
	return &result, nil
}

func (s *Transpiler) transpileFieldExpr(from *expr.Expr) (string, error) {
	kind := from.ExprKind

	switch kind.(type) {
	case *expr.Expr_IdentExpr:
		return s.transpileFieldIdentExpr(from.GetIdentExpr())
	default:
		return "", fmt.Errorf("unsupported expr: %T", kind)
	}
}

func (s *Transpiler) transpileFieldIdentExpr(from *expr.Expr_Ident) (string, error) {
	return s.findFieldName(from.GetName()), nil
}

func (s *Transpiler) transpileValueExpr(from *expr.Expr) (interface{}, error) {
	kind := from.ExprKind

	switch kind.(type) {
	case *expr.Expr_ConstExpr:
		return s.transpileValueConstExpr(from.GetConstExpr())
	case *expr.Expr_CallExpr:
		return s.transpileValueCallExpr(from.GetCallExpr())
	default:
		return "", fmt.Errorf("unsupported expr: %T", kind)
	}
}

func (s *Transpiler) transpileValueConstExpr(from *expr.Constant) (interface{}, error) {
	switch kind := from.ConstantKind.(type) {
	case *expr.Constant_BoolValue:
		return kind.BoolValue, nil
	case *expr.Constant_DoubleValue:
		return kind.DoubleValue, nil
	case *expr.Constant_Int64Value:
		return kind.Int64Value, nil
	case *expr.Constant_StringValue:
		return kind.StringValue, nil
	case *expr.Constant_Uint64Value:
		return int64(kind.Uint64Value), nil
	default:
		return nil, fmt.Errorf("unsupported const expr: %T", kind)
	}
}

func (s *Transpiler) transpileValueCallExpr(from *expr.Expr_Call) (interface{}, error) {
	switch from.Function {
	case filtering.FunctionTimestamp:
		return s.transpileTimestampCallExpr(from)
	default:
		return nil, fmt.Errorf("unsupported function call: %s", from.Function)
	}
}

func (s *Transpiler) transpileNotCallExpr(from *expr.Expr_Call) (*rel.FilterQuery, error) {
	if len(from.Args) != 1 {
		return nil, fmt.Errorf(
			"unexpected number of arguments to `%s` expression: %d",
			filtering.FunctionNot,
			len(from.Args),
		)
	}

	rhsExpr, err := s.transpileBoolExpr(from.Args[0])
	if err != nil {
		return nil, err
	}

	result := rel.Not(*rhsExpr)
	return &result, nil
}

func (s *Transpiler) transpileTimestampCallExpr(from *expr.Expr_Call) (interface{}, error) {
	if len(from.Args) != 1 {
		return nil, fmt.Errorf(
			"unexpected number of arguments to `%s`: %d", from.Function, len(from.Args),
		)
	}

	constArg, ok := from.Args[0].ExprKind.(*expr.Expr_ConstExpr)
	if !ok {
		return nil, fmt.Errorf("expected constant string arg to %s", from.Function)
	}

	stringArg, ok := constArg.ConstExpr.ConstantKind.(*expr.Constant_StringValue)
	if !ok {
		return nil, fmt.Errorf("expected constant string arg to %s", from.Function)
	}

	timeArg, err := time.Parse(time.RFC3339, stringArg.StringValue)
	if err != nil {
		return nil, fmt.Errorf("invalid string arg to %s: %w", from.Function, err)
	}

	return timeArg, nil
}

func (s *Transpiler) findFieldName(name string) string {
	if s.FieldMap == nil {
		return name
	}

	overrideName, ok := s.FieldMap[name]
	if !ok {
		return name
	}

	return overrideName
}
