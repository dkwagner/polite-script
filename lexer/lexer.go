package lexer

import (
	"fmt"
	"os"
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
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char in context of line)
	readPosition int  // current reading position in input (actual cursor in input, not relative to line)
	line         int  //current line being read of input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, position: 0, readPosition: 1, ch: input[0]}
	return l
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
		tok = comment(l)
	case '+':
		tok = newToken(token.OP_PLUS, "+", l.position, l.line)
	case '-':
		tok = minusOrNumber(l)
	case '*':
		tok = newToken(token.OP_MULT, "*", l.position, l.line)
	case '/':
		tok = newToken(token.OP_DIV, "/", l.position, l.line)
	case '"':
		tok = stringLiteral(l)
	case 'i':
		keyphrases := []string{token.OP_GREATER_OR_EQUAL,
			token.OP_GREATER,
			token.OP_LESS_OR_EQUAL,
			token.OP_LESS,
			token.OP_EQUAL,
			token.TYPE_INT}
		tok = lookupKeyphrase(l, keyphrases)
	case 0:
		tok = newToken(token.EOF, "", l.position, l.line)
	default:
		tok = identifierOrInteger(l)
	}

	l.readChar()
	return tok
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
		l.position = 0
		l.ch = l.input[l.readPosition]
		l.readPosition++
		return
	}

	l.position++
	l.readPosition++
}

func newToken(tokenType token.TokenType, literal string, pos int, line int) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Position: pos, Line: line}
}

func lookupKeyphrase(l *Lexer, tokenTypes []string) token.Token {

	line := l.line
	position := l.position

	for _, tt := range tokenTypes {
		keyphrase := reservedPhrases[tt]

		// If keyphrase not in list, throw an error, we made a goof
		if keyphrase == "" {
			error(fmt.Sprintf("ERROR: Invalid keyphrase type %s", tt))
		}

		// check that keyphrase is not longer than input
		// if so, skip this keyphrase in the list
		if len(keyphrase)+l.position > len(l.input) {
			continue
		}

		subString := string(l.input[l.position : l.position+len(keyphrase)])

		if subString == keyphrase {
			l.position = l.position + len(keyphrase)
			l.readPosition = l.readPosition + len(keyphrase)
			l.ch = l.input[l.readPosition-1]
			return newToken(token.TokenType(tt), subString, position, line)
		}
	}

	return identifier(l)
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

// Skips comment, only records that there was a comment
func comment(l *Lexer) token.Token {

	tokenStart := l.position
	line := l.line

	for l.ch != '\n' || l.ch == 0 {
		l.readChar()
	}

	return newToken(token.COMMENT, "#", tokenStart, line)
}

// Assumes that l.ch is a " to start
func stringLiteral(l *Lexer) token.Token {

	var sb strings.Builder
	tokenStart := l.position
	line := l.line

	l.readChar()

	for l.ch != '"' {

		if l.ch == 0 {
			error(fmt.Sprintf("LEXER ERROR: Reached EOF before string completed, line %d, column %d",
				l.line, l.position))
		}

		if l.ch == '\n' {
			error(fmt.Sprintf("LEXER ERROR: Illegal newline in string, line %d, column %d",
				l.line, l.position))
		}

		sb.WriteByte(l.ch)
		l.readChar()
	}

	return newToken(token.STRING, sb.String(), tokenStart, line)
}

func error(errorMsg string) {
	fmt.Println(errorMsg)
	os.Exit(-1)
}
