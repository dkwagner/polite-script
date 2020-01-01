package parser

import (
	"pscript/ast"
	"pscript/lexer"
	"pscript/token"
)

type Parser struct {
	l *lexer.Lexer

	CurToken  token.Token
	PeekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.CurToken = p.PeekToken
	p.PeekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
