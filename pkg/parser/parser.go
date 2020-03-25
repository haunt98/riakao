// https://tdop.github.io/
package parser

import (
	"fmt"
	"riakao/pkg/expression"
	"riakao/pkg/scanner"
	"riakao/pkg/token"
	"strconv"
	"strings"
)

type Parser struct {
	bs     *scanner.BufferScanner
	nudFns map[token.Token]nudFn // nud short for null denotation
	ledFns map[token.Token]ledFn // led short for left denotation
}

type nudFn func(token.Token, string) (expression.Expression, error)
type ledFn func(token.Token, string, expression.Expression) (expression.Expression, error)

func NewParser(bs *scanner.BufferScanner) *Parser {
	p := &Parser{
		bs: bs,
	}

	p.nudFns = map[token.Token]nudFn{
		token.Ident:             p.nudIdent,
		token.Int:               p.nudInt,
		token.String:            p.nudString,
		token.Not:               p.nudNot,
		token.LeftParenthesis:   p.nudParenthesis,
		token.LeftCurlyBracket:  p.nudCurlyBracket,
		token.LeftSquareBracket: p.nudSquareBracket,
	}
	p.ledFns = map[token.Token]ledFn{
		token.And:   p.ledInfix,
		token.Or:    p.ledInfix,
		token.In:    p.ledInfix,
		token.Equal: p.ledInfix,
	}

	return p
}

func (p *Parser) Parse() (expression.Expression, error) {
	return p.parseExpression(0)
}

func (p *Parser) parseExpression(precedence int) (expression.Expression, error) {
	tok, text := p.bs.Scan()
	left, err := p.nullDenotation(tok, text)
	if err != nil {
		return expression.Expression{}, err
	}

	for {
		peekTok, _ := p.bs.Peek()
		if precedence >= peekTok.Precedence() {
			break
		}

		tok, text = p.bs.Scan()
		left, err = p.leftDenotation(tok, text, left)
		if err != nil {
			return expression.Expression{}, err
		}
	}

	return left, nil
}

func (p *Parser) nullDenotation(tok token.Token, text string) (expression.Expression, error) {
	fn, ok := p.nudFns[tok]
	if !ok {
		return expression.Expression{}, fmt.Errorf("%s not impelement null denotation", tok.String())
	}

	return fn(tok, text)
}

func (p *Parser) nudIdent(tok token.Token, text string) (result expression.Expression, err error) {
	result = expression.Expression{
		Token:    tok,
		Value:    text,
		Children: nil,
	}
	return
}

func (p *Parser) nudInt(tok token.Token, text string) (result expression.Expression, err error) {
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return expression.Expression{}, err
	}

	result = expression.Expression{
		Token:    tok,
		Value:    value,
		Children: nil,
	}
	return
}

func (p *Parser) nudString(tok token.Token, text string) (result expression.Expression, err error) {
	result = expression.Expression{
		Token:    token.String,
		Value:    strings.Trim(text, `"`),
		Children: nil,
	}
	return
}

func (p *Parser) nudNot(tok token.Token, _ string) (result expression.Expression, err error) {
	var expr expression.Expression
	expr, err = p.parseExpression(0)
	if err != nil {
		return
	}

	result = expression.Expression{
		Token: tok,
		Value: nil,
		Children: []expression.Expression{
			expr,
		},
	}
	return
}

func (p *Parser) nudParenthesis(tok token.Token, _ string) (result expression.Expression, err error) {
	var expr expression.Expression
	expr, err = p.parseExpression(0)
	if err != nil {
		return
	}

	if !p.expectToken(token.RightParenthesis) {
		err = fmt.Errorf("expect %s", token.RightParenthesis.String())
		return
	}

	result = expression.Expression{
		Token: tok,
		Value: nil,
		Children: []expression.Expression{
			expr,
		},
	}
	return
}

func (p *Parser) nudCurlyBracket(tok token.Token, _ string) (result expression.Expression, err error) {
	var expr expression.Expression
	expr, err = p.parseExpression(0)
	if err != nil {
		return
	}

	if !p.expectToken(token.RightCurlyBracket) {
		err = fmt.Errorf("expect %s", token.RightCurlyBracket.String())
		return
	}

	result = expression.Expression{
		Token: tok,
		Value: nil,
		Children: []expression.Expression{
			expr,
		},
	}
	return
}

func (p *Parser) nudSquareBracket(tok token.Token, _ string) (result expression.Expression, err error) {
	expr := expression.Expression{
		Token:    tok,
		Value:    nil,
		Children: nil,
	}

	var child expression.Expression
	for {
		peekTok, _ := p.bs.Peek()
		if peekTok == token.RightSquareBracket {
			break
		}

		child, err = p.parseExpression(0)
		if err != nil {
			return
		}
		expr.Children = append(expr.Children, child)

		peekTok, _ = p.bs.Peek()
		if peekTok != token.Comma {
			break
		}

		// skip comma
		_, _ = p.bs.Scan()
	}

	if !p.expectToken(token.RightSquareBracket) {
		err = fmt.Errorf("expect %s", token.RightSquareBracket.String())
		return
	}

	result = expr
	return
}

func (p *Parser) leftDenotation(tok token.Token, text string, expr expression.Expression) (result expression.Expression, err error) {
	fn, ok := p.ledFns[tok]
	if !ok {
		err = fmt.Errorf("%s not implement left denotation", tok.String())
	}

	return fn(tok, text, expr)
}

func (p *Parser) ledInfix(tok token.Token, _ string, expr expression.Expression) (result expression.Expression, err error) {
	var rightExpr expression.Expression
	rightExpr, err = p.parseExpression(tok.Precedence())
	if err != nil {
		return
	}

	result = expression.Expression{
		Token: tok,
		Children: []expression.Expression{
			expr,
			rightExpr,
		},
	}
	return
}

func (p *Parser) expectToken(tok token.Token) bool {
	gotTok, _ := p.bs.Scan()
	return gotTok == tok
}
