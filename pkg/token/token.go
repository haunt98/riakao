package token

type Token int

const (
	Illegal Token = iota
	EOF

	Ident
	Int
	String

	And
	Or
	Not
	In
	Equal

	LeftParenthesis
	RightParenthesis
	LeftCurlyBracket
	RightCurlyBracket
	LeftSquareBracket
	RightSquareBracket
	Comma
)

var tokenStrings = map[Token]string{
	Illegal: "illegal",
	EOF:     "EOF",

	Ident:  "ident",
	Int:    "int",
	String: "string",

	And:   "and",
	Or:    "or",
	Not:   "not",
	In:    "in",
	Equal: "==",

	LeftParenthesis:    "(",
	RightParenthesis:   ")",
	LeftCurlyBracket:   "{",
	RightCurlyBracket:  "}",
	LeftSquareBracket:  "[",
	RightSquareBracket: "]",
	Comma:              ",",
}

func (tok Token) String() string {
	result, ok := tokenStrings[tok]
	if !ok {
		return "Unknown"
	}
	return result
}

// https://en.wikipedia.org/wiki/Order_of_operations
var tokenPrecedences = map[Token]int{
	Or:    1,
	And:   2,
	Equal: 3,
	In:    3,
	Not:   4,
}

func (tok Token) Precedence() int {
	result, ok := tokenPrecedences[tok]
	if !ok {
		return 0
	}
	return result
}
