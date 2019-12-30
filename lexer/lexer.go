package lexer

import (
	"polite-script/token"
)

// var reservedPhrases map[string]string

// func init() {
// 	reservedPhrases[token.KEYPHRASE_START] = "Please"
// 	reservedPhrases[token.KEYPHRASE_END] = "Thanks"
// 	reservedPhrases[token.KEYPHRASE_IF] = "check if"
// 	reservedPhrases[token.KEYPHRASE_DECLARE] = "set"
// 	reservedPhrases[token.KEYPHRASE_LOOP] = "loop while"
// 	reservedPhrases[token.OP_ASSIGN] = "equal to"
// 	reservedPhrases[token.OP_EQUAL] = "is equal to"
// 	reservedPhrases[token.OP_NOT_EQUAL] = "does not equal"
// 	reservedPhrases[token.OP_GREATER] = "is greater than"
// 	reservedPhrases[token.OP_LESS] = "is less than"
// 	reservedPhrases[token.OP_GREATER_OR_EQUAL] = "is greater than or equal to"
// 	reservedPhrases[token.OP_LESS_OR_EQUAL] = "is less than or equal to"
// }

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '(':
		tok = newToken(token.LPAREN, "(")
	case ')':
		tok = newToken(token.RPAREN, ")")
	case '{':
		tok = newToken(token.LBRACE, "{")
	case '}':
		tok = newToken(token.RBRACE, "}")
	case '#':
		tok = newToken(token.COMMENT, "#")
	case 0:
		tok = newToken(token.EOF, "")
	}

	l.readChar()
	return tok
}
