package parser

import (
	"fmt"
	"pscript/ast"
	"pscript/lexer"
	"pscript/token"
)

// Parser struct representing the parser
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

// New creates new instance of parser
// with the curToken and peekToken populated
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns a slice of all errors
// encountered while parsing
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

// ParseProgram parses all given input
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt == nil {
		} else {
			fmt.Println("Declaration Statement", stmt)
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.KEYPHRASE_DECLARE:
		return p.parseDeclarationStatement()
	default:
		return nil
	}
}

func (p *Parser) parseDeclarationStatement() *ast.DeclarationStatement {
	stmt := &ast.DeclarationStatement{Token: p.curToken}

	if !p.expectPeek(token.ID) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return stmt
}
