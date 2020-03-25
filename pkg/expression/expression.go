package expression

import (
	"riakao/pkg/token"
)

type Expression struct {
	Token    token.Token
	Value    interface{}
	Children []Expression
}
