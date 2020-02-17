package parser

import (
	"testing"

	"github.com/dkwagner/pscript/ast"
	"github.com/dkwagner/pscript/lexer"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
}

func TestNew(t *testing.T) {
	input := "token1 token2"

	lexer := lexer.New(input)

	tests := struct {
		expectedCurTokenValue  string
		expectedPeekTokenValue string
	}{
		expectedCurTokenValue:  "token1",
		expectedPeekTokenValue: "token2",
	}

	actual := New(lexer)

	if actual.curToken.Literal != tests.expectedCurTokenValue {
		t.Errorf("FAIL on CurToken value: expected %s, got %s", tests.expectedCurTokenValue, actual.curToken.Literal)
	}

	if actual.peekToken.Literal != tests.expectedPeekTokenValue {
		t.Errorf("FAIL on PeekToken value: expected %s, got %s", tests.expectedPeekTokenValue, actual.peekToken.Literal)
	}
}

func TestDeclarationStatements(t *testing.T) {
	input := `
Please set x equal to 5
Please set a equal to 3
Please set b equal to -1
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"a"},
		{"b"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testDeclarationStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testDeclarationStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "set" {
		t.Errorf("s.tokenLiteral not 'set'. got %q", s.TokenLiteral())
	}

	declStmt, ok := s.(*ast.DeclarationStatement)
	if !ok {
		t.Errorf("s not *ast.DeclarationStatement. got %T", s)
		return false
	}

	if declStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got %s", name, declStmt.Name.Value)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
Returns 5
Returns "Hello"
Returns identifier	
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.ReturnStatemnt. Got %T", stmt)
		}
		if returnStmt.TokenLiteral() != "Returns" {
			t.Errorf("returnStmt.TokenLiteral not 'Returns'. Got %q", returnStmt.TokenLiteral())
		}
	}

}
