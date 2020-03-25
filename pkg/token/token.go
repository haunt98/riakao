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

	LeftParenthesis    // (
	RightParenthesis   // )
	LeftCurlyBracket   // {
	RightCurlyBracket  // }
	LeftSquareBracket  // [
	RightSquareBracket // ]
	Comma              // ,
)

var tokenStrings = map[Token]string{
	Illegal: "Illegal",
	EOF:     "EOF",

	Ident:  "Ident",
	Int:    "Int",
	String: "String",

	And:   "And",
	Or:    "Or",
	Not:   "Not",
	In:    "In",
	Equal: "Equal",

	LeftParenthesis:    "LeftParenthesis",
	RightParenthesis:   "RightParenthesis",
	LeftCurlyBracket:   "LeftCurlyBracket",
	RightCurlyBracket:  "RightCurlyBracket",
	LeftSquareBracket:  "LeftSquareBracket",
	RightSquareBracket: "RightSquareBracket",
	Comma:              "Comma",
}

func (tok Token) String() string {
	result, ok := tokenStrings[tok]
	if !ok {
		return "Unknown"
	}
	return result
}

const (
	LowestPrecedence = 0
)

var tokenPrecedences = map[Token]int{
	Or:    LowestPrecedence + 1,
	And:   LowestPrecedence + 2,
	Equal: LowestPrecedence + 3,
	In:    LowestPrecedence + 3,
	Not:   LowestPrecedence + 4,
}

func (tok Token) Precedence() int {
	result, ok := tokenPrecedences[tok]
	if !ok {
		return LowestPrecedence
	}
	return result
}
