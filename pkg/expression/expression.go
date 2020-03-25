package expression

import (
	"fmt"
	"riakao/pkg/token"
)

type Expression struct {
	Token    token.Token
	Value    interface{}
	Children []Expression
}

func (expr Expression) GetValue() (interface{}, error) {
	switch expr.Token {
	case token.Ident, token.Int, token.String:
		return expr.Value, nil
	case token.LeftCurlyBracket:
		if len(expr.Children) != 1 {
			return nil, fmt.Errorf("ident missing chilren")
		}

		return (expr.Children[0]).GetValue()
	default:
		return nil, fmt.Errorf("%s not implement get value", expr.Token.String())
	}
}

func (expr Expression) GetValues() ([]interface{}, error) {
	switch expr.Token {
	case token.LeftSquareBracket:
		values := make([]interface{}, 0, len(expr.Children))
		for _, child := range expr.Children {
			value, err := child.GetValue()
			if err != nil {
				return nil, err
			}

			values = append(values, value)
		}
		return values, nil
	default:
		return nil, fmt.Errorf("%s not implement get values", expr.Token.String())
	}
}
