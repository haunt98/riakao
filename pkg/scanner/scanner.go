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
		case and:
			return token.And, text
		case or:
			return token.Or, text
		case not:
			return token.Not, text
		case in:
			return token.In, text
		default:
			return token.Ident, text
		}
	case scanner.Int:
		return token.Int, text
	case scanner.String:
		return token.String, text
	case equalSign:
		if expect := s.textScanner.Scan(); expect != equalSign {
			return token.Illegal, text + s.textScanner.TokenText()
		}
		return token.Equal, text + s.textScanner.TokenText()
	case leftParenthesis:
		return token.LeftParenthesis, text
	case rightParenthesis:
		return token.RightParenthesis, text
	case leftCurlyBracket:
		return token.LeftCurlyBracket, text
	case rightCurlyBracket:
		return token.RightCurlyBracket, text
	case leftSquareBracket:
		return token.LeftSquareBracket, text
	case rightSquareBracket:
		return token.RightSquareBracket, text
	case comma:
		return token.Comma, text
	default:
		return token.Illegal, text
	}
}
