package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers and literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	CHAR   = "CHAR"

	//Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	//Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords
	FUNCTION = "FUNC"
	VARIABLE = "VAR"
)