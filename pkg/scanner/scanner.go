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

	switch ch {
	case scanner.EOF:
		return token.EOF, text
	case scanner.Ident:
		switch strings.ToLower(text) {
		case token.And.String():
			return token.And, text
		case token.Or.String():
			return token.Or, text
		case token.Not.String():
			return token.Not, text
		case token.In.String():
			return token.In, text
		default:
			return token.Ident, text
		}
	case scanner.Int:
		return token.Int, text
	case scanner.String:
		return token.String, text
	case '=':
		if expect := s.textScanner.Scan(); expect != '=' {
			return token.Illegal, text + s.textScanner.TokenText()
		}
		return token.Equal, text + s.textScanner.TokenText()
	case '(':
		return token.LeftParenthesis, text
	case ')':
		return token.RightParenthesis, text
	case '{':
		return token.LeftCurlyBracket, text
	case '}':
		return token.RightCurlyBracket, text
	case '[':
		return token.LeftSquareBracket, text
	case ']':
		return token.RightSquareBracket, text
	case ',':
		return token.Comma, text
	default:
		return token.Illegal, text
	}
}
