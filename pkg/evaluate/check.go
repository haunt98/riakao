package evaluate

import (
	"fmt"
	"riakao/pkg/expression"
	"riakao/pkg/token"
)

func check(expr expression.Expression) (bool, error) {
	switch expr.Token {
	case token.And:
		return checkAnd(expr)
	case token.Or:
		return checkOr(expr)
	case token.Not:
		return checkNot(expr)
	case token.Equal:
		return checkEqual(expr)
	case token.In:
		return checkIn(expr)
	case token.LeftParenthesis:
		return checkParenthesis(expr)
	default:
		return false, fmt.Errorf("%s unimplement check", expr.Token.String())
	}
}

func checkAnd(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 2 {
		return false, fmt.Errorf("and operator missing 2 operands")
	}

	leftResult, err := check(expr.Children[0])
	if err != nil {
		return false, err
	}
	if !leftResult {
		return false, nil
	}

	rightResult, err := check(expr.Children[1])
	if err != nil {
		return false, err
	}

	return rightResult, nil
}

func checkOr(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 2 {
		return false, fmt.Errorf("or operator missing 2 operands")
	}

	leftResult, err := check(expr.Children[0])
	if err != nil {
		return false, err
	}
	if leftResult {
		return true, nil
	}

	rightResult, err := check(expr.Children[1])
	if err != nil {
		return false, err
	}

	return rightResult, nil
}

func checkNot(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 1 {
		return false, fmt.Errorf("not operator missing 1 operand")
	}

	childResult, err := check(expr.Children[0])
	if err != nil {
		return false, err
	}

	return !childResult, nil
}

func checkEqual(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 2 {
		return false, fmt.Errorf("equal operator missing 2 operands")
	}

	leftValue, err := (expr.Children[0]).GetValue()
	if err != nil {
		return false, err
	}

	rightValue, err := (expr.Children[1]).GetValue()
	if err != nil {
		return false, err
	}

	return leftValue == rightValue, nil
}

func checkIn(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 2 {
		return false, fmt.Errorf("in operator missing 2 operands")
	}

	leftValue, err := (expr.Children[0]).GetValue()
	if err != nil {
		return false, err
	}

	rightValues, err := (expr.Children[1]).GetValues()
	if err != nil {
		return false, err
	}

	return existInArray(rightValues, leftValue), nil
}

func checkParenthesis(expr expression.Expression) (bool, error) {
	if len(expr.Children) != 1 {
		return false, fmt.Errorf("parenthesis missing child")
	}

	return check(expr.Children[0])
}

func existInArray(arr []interface{}, item interface{}) bool {
	for i := range arr {
		if item == arr[i] {
			return true
		}
	}
	return false
}
