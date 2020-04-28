package parser

import (
	"errors"
	"fmt"

	"github.com/dkwagner/pscript/ast"
	"github.com/dkwagner/pscript/lexer"
	"github.com/dkwagner/pscript/token"
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
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Double check that a statement is legit
// This is done because thats the way it is
// But really this has to do with pointers getting returned that are
// actually pointers pointing to nil
// imagine something like *nil, but you cant do that with golang
// so here we are
func (p *Parser) validateStatement(stmt ast.Statement, err error) ast.Statement {
	if err != nil {
		return nil
	}
	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.KEYPHRASE_DECLARE:
		return p.validateStatement(p.parseDeclarationStatement())
	case token.KEYPHRASE_RETURN:
		return p.validateStatement(p.parseReturnStatemnt())
	default:
		return nil
	}
}

func (p *Parser) parseDeclarationStatement() (*ast.DeclarationStatement, error) {
	stmt := &ast.DeclarationStatement{Token: p.curToken}

	if !p.expectPeek(token.ID) {
		return nil, errors.New("Did not get expected token ID")
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return stmt, nil
}

func (p *Parser) parseReturnStatemnt() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// TODO implement expression parsing
	for !p.curTokenIs(token.END_LINE) {
		if p.curTokenIs(token.EOF) {
			return nil, errors.New("Did not get end line token before EOF")
		}
		p.nextToken()
	}

	return stmt, nil
}
