package scanner

import (
	"io"
	"riakao/pkg/token"
	"strings"
	"text/scanner"
)

type Scanner struct {
	textScanner *scanner.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	textScanner := &scanner.Scanner{}
	textScanner.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanInts
	textScanner.Init(r)

	return &Scanner{
		textScanner: textScanner,
	}
}

func (s *Scanner) Scan() (token.Token, string) {
	ch := s.textScanner.Scan()
	text := s.textScanner.TokenText()

	var tok token.Token
	switch ch {
	case scanner.EOF:
		tok = token.EOF
	case scanner.Ident:
		switch strings.ToLower(text) {
		case And:
			tok = token.And
		case Or:
			tok = token.Or
		case Not:
			tok = token.Not
		case In:
			tok = token.In
		default:
			tok = token.Ident
		}
	case scanner.Int:
		tok = token.Int
	case scanner.String:
		tok = token.String
	case LeftParenthesis:
		tok = token.LeftParenthesis
	case RightParenthesis:
		tok = token.RightParenthesis
	case LeftCurlyBracket:
		tok = token.LeftCurlyBracket
	case RightCurlyBracket:
		tok = token.RightCurlyBracket
	case LeftSquareBracket:
		tok = token.LeftSquareBracket
	case RightSquareBracket:
		tok = token.RightSquareBracket
	case Comma:
		tok = token.Comma
	default:
		tok = token.Illegal
	}

	return tok, text
}
