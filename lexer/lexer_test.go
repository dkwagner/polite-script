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

	l := New(input)

	var expectedCh byte = '('
	expectedReadPosition := 1
	expectedPosition := 0

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

func TestReadChar_WhenNewLine(t *testing.T) {
	input := `
(`
	l := New(input)

	l.readChar()

	if l.position != 0 {
		t.Errorf("Lexer.position wrong, expected 0, got %d", l.position)
	}

	if l.line != 2 {
		t.Errorf("Lexer.line wrong, expected 1, got %d", l.line)
	}

	if l.readPosition != 2 {
		t.Errorf("Lexer.readPosition wrong, expected 2, got %d", l.readPosition)
	}

	if l.ch != '(' {
		t.Errorf("Lexer.ch wrong, expected (, got %q", l.ch)
	}
}

func TestNextToken(t *testing.T) {
	input := `( ) { } + - * /
# This is a comment
"This is a string"
is greater than or equal to
is greater than
is less than or equal to
is less than
is equal to
integer
i
Please
Thanks
check if
set
loop while
define function
with arguments
that returns
boolean
string
equal to
true
false
abcde 123 -123`

	tests := []struct {
		expectedType     token.TokenType
		expectedLiteral  string
		expectedPosition int
		expectedLine     int
	}{
		{token.LPAREN, "(", 0, 1},
		{token.RPAREN, ")", 2, 1},
		{token.LBRACE, "{", 4, 1},
		{token.RBRACE, "}", 6, 1},
		{token.OP_PLUS, "+", 8, 1},
		{token.OP_MINUS, "-", 10, 1},
		{token.OP_MULT, "*", 12, 1},
		{token.OP_DIV, "/", 14, 1},
		{token.COMMENT, "#", 0, 2},
		{token.STRING, "This is a string", 0, 3},
		{token.OP_GREATER_OR_EQUAL, "is greater than or equal to", 0, 4},
		{token.OP_GREATER, "is greater than", 0, 5},
		{token.OP_LESS_OR_EQUAL, "is less than or equal to", 0, 6},
		{token.OP_LESS, "is less than", 0, 7},
		{token.OP_EQUAL, "is equal to", 0, 8},
		{token.TYPE_INT, "integer", 0, 9},
		{token.ID, "i", 0, 10},
		{token.KEYPHRASE_START, "Please", 0, 11},
		{token.KEYPHRASE_END, "Thanks", 0, 12},
		{token.KEYPHRASE_IF, "check if", 0, 13},
		{token.KEYPHRASE_DECLARE, "set", 0, 14},
		{token.KEYPHRASE_LOOP, "loop while", 0, 15},
		{token.KEYPHRASE_FUNC_DECL, "define function", 0, 16},
		{token.KEYPHRASE_ARG_DECL, "with arguments", 0, 17},
		{token.KEYPHRASE_RETURN, "that returns", 0, 18},
		{token.TYPE_BOOL, "boolean", 0, 19},
		{token.TYPE_STRING, "string", 0, 20},
		{token.OP_ASSIGN, "equal to", 0, 21},
		{token.KEYWORD_TRUE, "true", 0, 22},
		{token.KEYWORD_FALSE, "false", 0, 23},
		{token.ID, "abcde", 0, 24},
		{token.INT, "123", 6, 24},
		{token.INT, "-123", 10, 24},
		{token.EOF, "", 13, 24},
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
