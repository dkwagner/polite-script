package lexer

import (
	"polite-script/token"
	"testing"
)

func TestNew(t *testing.T) {
	input := `(`

	test := struct {
		expectedInput        string
		expectedPosition     int
		expectedReadPosition int
		expectedCh           byte
	}{
		expectedInput:        input,
		expectedPosition:     0,
		expectedReadPosition: 1,
		expectedCh:           '(',
	}

	l := New(input)

	if l.input != test.expectedInput {
		t.Errorf("Lexer.input wrong, expected %q, got %q",
			test.expectedInput, l.input)
	}

	if l.position != test.expectedPosition {
		t.Errorf("Lexer.position wrong, expected %q, got %q",
			test.expectedPosition, l.position)
	}

	if l.readPosition != test.expectedReadPosition {
		t.Errorf("Lexer.readPosition wrong, expected %q, got %q",
			test.expectedReadPosition, l.readPosition)
	}

	if l.ch != test.expectedCh {
		t.Errorf("Lexer.input wrong, expected %q, got %q",
			test.expectedCh, l.ch)
	}
}

func TestReadChar_WhenReadPositionWithinInput(t *testing.T) {
	input := `((`

	l := &Lexer{input: input}

	var expectedCh byte = '('
	expectedReadPosition := 1
	expectedPosition := 0

	l.readChar()

	if l.ch != expectedCh {
		t.Errorf("Lexer.ch wrong, expected %q, got %q",
			expectedCh, l.ch)
	}

	if l.position != expectedPosition {
		t.Errorf("Lexer.position wrong, expected %d, got %d",
			expectedPosition, l.position)
	}

	if l.readPosition != expectedReadPosition {
		t.Errorf("Lexer.readPosition wrong, expected %d, got %d",
			expectedReadPosition, l.readPosition)
	}
}

func TestNextToken(t *testing.T) {
	input := `(){}#`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMENT, "#"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
