package evaluate

import (
	"fmt"
	"riakao/pkg/expression"
	"riakao/pkg/token"
)

func Evaluate(expr expression.Expression, args map[interface{}]interface{}) (bool, error) {
	if err := replaceIdent(&expr, args); err != nil {
		return false, err
	}

	return check(expr)
}

func replaceIdent(expr *expression.Expression, args map[interface{}]interface{}) error {
	if expr.Token == token.Ident {
		key := expr.Value
		val, ok := args[key]
		if !ok {
			return fmt.Errorf("args missing key")
		}

		expr.Value = val
		return nil
	}

	for i := range expr.Children {
		if err := replaceIdent(&expr.Children[i], args); err != nil {
			return err
		}
	}

	return nil
}
