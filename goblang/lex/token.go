package token

// token types are int constants
type Type int

const (
	// special tokens
	INVALID TokenType = iota
	EOF
	// structural tokens
	IDENTIFIER
	KEYWORD
	SEPARATOR
	OPERATOR
	LITERAL
	COMMENT
)

type Unit struct {
	tokType   TokenType
	tokString string
}
