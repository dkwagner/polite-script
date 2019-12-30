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
		expectedLine         int
	}{
		expectedInput:        input,
		expectedPosition:     0,
		expectedReadPosition: 1,
		expectedCh:           '(',
		expectedLine:         1,
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

	if l.line != test.expectedLine {
		t.Errorf("Lexer.line wrong, expected %d, got %d",
			test.expectedLine, l.line)
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
	input := `(){}#+-*/abcde`

	tests := []struct {
		expectedType     token.TokenType
		expectedLiteral  string
		expectedPosition int
		expectedLine     int
	}{
		{token.LPAREN, "(", 0, 1},
		{token.RPAREN, ")", 1, 1},
		{token.LBRACE, "{", 2, 1},
		{token.RBRACE, "}", 3, 1},
		{token.COMMENT, "#", 4, 1},
		{token.OP_PLUS, "+", 5, 1},
		{token.OP_MINUS, "-", 6, 1},
		{token.OP_MULT, "*", 7, 1},
		{token.OP_DIV, "/", 8, 1},
		{token.ID, "abcde", 9, 1},
		{token.EOF, "", 14, 1},
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

		if tok.Position != tt.expectedPosition {
			t.Errorf("tests[%d] - Position wrong. expected=%d, got=%d",
				i, tt.expectedPosition, tok.Position)
		}

		if tok.Line != tt.expectedLine {
			t.Errorf("tests[%d] - Line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}
	}
}
