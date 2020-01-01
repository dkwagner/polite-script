package ast

import "pscript/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// Declaration statement
// token.DECLARE
type DeclarationStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *DeclarationStatement) statementNode()       {}
func (ls *DeclarationStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier Expression
// token.ID
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
