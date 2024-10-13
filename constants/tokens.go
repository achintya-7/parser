package constants

import "fmt"

type TokenType int

const (
	TOKEN_EOF TokenType = iota
	TOKEN_VARIABLE
	TOKEN_NUMBER
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_DIVIDE
	TOKEN_DOUBLE_EQUAL
	TOKEN_NOT_EQUAL
	TOKEN_EQUAL
	TOKEN_NOT
	TOKEN_LEFT_PAREN
	TOKEN_RIGHT_PAREN
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) ToString() string {
	return fmt.Sprintf("Type: %d, Lexeme: %s, Line: %d", t.Type, t.Lexeme, t.Line)
}
