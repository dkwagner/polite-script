package token

// TokenType is the type of token
type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position int
	Line     int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	ID     = "ID"
	INT    = "INT"
	STRING = "STRING"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	COMMA  = ","

	// Keyphrases
	KEYPHRASE_START     = "START"
	KEYPHRASE_END       = "END"
	KEYPHRASE_IF        = "IF"
	KEYPHRASE_DECLARE   = "DECLARE"
	KEYPHRASE_LOOP      = "LOOP"
	KEYPHRASE_FUNC_DECL = "FUNCTION_DECLARATION"
	KEYPHRASE_ARG_DECL  = "ARG_DECLARATION"
	KEYPHRASE_RETURN    = "RETURN_DECLARATION"

	KEYWORD_TRUE  = "TRUE"
	KEYWORD_FALSE = "FALSE"

	TYPE_BOOL   = "BOOL"
	TYPE_INT    = "INTEGER"
	TYPE_STRING = "STRING"

	// Operators
	OP_ASSIGN           = "ASSIGN"
	OP_EQUAL            = "EQUAL"
	OP_NOT_EQUAL        = "NOT_EQUAL"
	OP_GREATER          = "GREATER_THAN"
	OP_LESS             = "LESS_THAN"
	OP_GREATER_OR_EQUAL = "GREATER_THAN_OR_EQUAL"
	OP_LESS_OR_EQUAL    = "LESS_THAN_OR_EQUAL"
	OP_PLUS             = "PLUS"
	OP_MINUS            = "MINUS"
	OP_MULT             = "MULT"
	OP_DIV              = "DIV"

	// Delimiters
	END_LINE = "END_LINE"

	COMMENT = "#"
)
