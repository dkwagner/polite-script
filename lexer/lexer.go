package lexer

import (
	"fmt"
	"polite-script/token"
	"polite-script/util"
	"strings"
)

var reservedPhrases map[string]string = make(map[string]string)

func init() {
	reservedPhrases[token.KEYPHRASE_START] = "Please"
	reservedPhrases[token.KEYPHRASE_END] = "Thanks"
	reservedPhrases[token.KEYPHRASE_IF] = "check if"
	reservedPhrases[token.KEYPHRASE_DECLARE] = "set"
	reservedPhrases[token.KEYPHRASE_LOOP] = "loop while"
	reservedPhrases[token.KEYPHRASE_FUNC_DECL] = "define function"
	reservedPhrases[token.KEYPHRASE_ARG_DECL] = "with arguments"
	reservedPhrases[token.KEYPHRASE_RETURN_DECAL] = "that returns"
	reservedPhrases[token.TYPE_BOOL] = "boolean"
	reservedPhrases[token.TYPE_STRING] = "string"
	reservedPhrases[token.TYPE_INT] = "integer"
	reservedPhrases[token.OP_ASSIGN] = "equal to"
	reservedPhrases[token.OP_EQUAL] = "is equal to"
	reservedPhrases[token.OP_NOT_EQUAL] = "does not equal"
	reservedPhrases[token.OP_GREATER] = "is greater than"
	reservedPhrases[token.OP_LESS] = "is less than"
	reservedPhrases[token.OP_GREATER_OR_EQUAL] = "is greater than or equal to"
	reservedPhrases[token.OP_LESS_OR_EQUAL] = "is less than or equal to"
	reservedPhrases[token.OP_PLUS] = "plus"
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char in context of line)
	readPosition int  // current reading position in input (actual cursor in input, not relative to line)
	line         int  //current line being read of input
	ch           byte // current char under examination
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	for util.IsWhitespace(l.ch) {
		l.readChar()
	}

	switch l.ch {
	case '(':
		tok = newToken(token.LPAREN, "(", l.position, l.line)
	case ')':
		tok = newToken(token.RPAREN, ")", l.position, l.line)
	case '{':
		tok = newToken(token.LBRACE, "{", l.position, l.line)
	case '}':
		tok = newToken(token.RBRACE, "}", l.position, l.line)
	case '#':
		tok = newToken(token.COMMENT, "#", l.position, l.line)
	case '+':
		tok = newToken(token.OP_PLUS, "+", l.position, l.line)
	case '-':
		tok = minusOrNumber(l)
	case '*':
		tok = newToken(token.OP_MULT, "*", l.position, l.line)
	case '/':
		tok = newToken(token.OP_DIV, "/", l.position, l.line)
	case 0:
		tok = newToken(token.EOF, "", l.position, l.line)
	default:
		tok = identifierOrInteger(l)
	}

	l.readChar()
	return tok
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, position: -1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}

	l.ch = l.input[l.readPosition]
	fmt.Printf("reading: %q @ position: %d\n", l.ch, l.position)

	// if newline, set position to 0, and increment line
	// still increment readPosition to next character in input for next readChar call
	if l.ch == '\n' {
		l.line++
		l.position = -1
		l.readPosition++
		return
	}

	l.position++
	l.readPosition++
}

func newToken(tokenType token.TokenType, literal string, pos int, line int) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Position: pos, Line: line}
}

func identifierOrInteger(l *Lexer) token.Token {

	if l.ch == '-' || util.IsDigit(l.ch) {
		return integer(l)
	}

	return identifier(l)
}

func identifier(l *Lexer) token.Token {

	var sb strings.Builder
	tokenStart := l.position

	for util.IsLetter(l.ch) || util.IsDigit(l.ch) || l.ch == '_' {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	return newToken(token.ID, sb.String(), tokenStart, l.line)
}

func integer(l *Lexer) token.Token {

	var sb strings.Builder
	tokenStart := l.position

	if l.ch == '-' {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	for util.IsDigit(l.ch) {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	return newToken(token.INT, sb.String(), tokenStart, l.line)
}

func minusOrNumber(l *Lexer) token.Token {

	if (l.readPosition) <= len(l.input)-1 {
		if util.IsWhitespace(l.input[l.readPosition]) {
			return newToken(token.OP_MINUS, "-", l.position, l.line)
		}
	}

	return integer(l)
}
