package lexer

import (
	"parser/constants"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected []constants.Token
	}{
		{
			input: "+-*/()",
			expected: []constants.Token{
				{Type: constants.TOKEN_PLUS, Lexeme: "+", Line: 1},
				{Type: constants.TOKEN_MINUS, Lexeme: "-", Line: 1},
				{Type: constants.TOKEN_MULTIPLY, Lexeme: "*", Line: 1},
				{Type: constants.TOKEN_DIVIDE, Lexeme: "/", Line: 1},
				{Type: constants.TOKEN_LEFT_PAREN, Lexeme: "(", Line: 1},
				{Type: constants.TOKEN_RIGHT_PAREN, Lexeme: ")", Line: 1},
				{Type: constants.TOKEN_EOF, Lexeme: "", Line: 1},
			},
		},
		{
			input: "1 + 25",
			expected: []constants.Token{
				{Type: constants.TOKEN_NUMBER, Lexeme: "1", Line: 1},
				{Type: constants.TOKEN_PLUS, Lexeme: "+", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "25", Line: 1},
				{Type: constants.TOKEN_EOF, Lexeme: "", Line: 1},
			},
		},
		{
			input: "0 / 22",
			expected: []constants.Token{
				{Type: constants.TOKEN_NUMBER, Lexeme: "0", Line: 1},
				{Type: constants.TOKEN_DIVIDE, Lexeme: "/", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "22", Line: 1},
				{Type: constants.TOKEN_EOF, Lexeme: "", Line: 1},
			},
		},
		{
			input: "(1 + 2 * 3) + (x - y / 2)",
			expected: []constants.Token{
				{Type: constants.TOKEN_LEFT_PAREN, Lexeme: "(", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "1", Line: 1},
				{Type: constants.TOKEN_PLUS, Lexeme: "+", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "2", Line: 1},
				{Type: constants.TOKEN_MULTIPLY, Lexeme: "*", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "3", Line: 1},
				{Type: constants.TOKEN_RIGHT_PAREN, Lexeme: ")", Line: 1},
				{Type: constants.TOKEN_PLUS, Lexeme: "+", Line: 1},
				{Type: constants.TOKEN_LEFT_PAREN, Lexeme: "(", Line: 1},
				{Type: constants.TOKEN_VARIABLE, Lexeme: "x", Line: 1},
				{Type: constants.TOKEN_MINUS, Lexeme: "-", Line: 1},
				{Type: constants.TOKEN_VARIABLE, Lexeme: "y", Line: 1},
				{Type: constants.TOKEN_DIVIDE, Lexeme: "/", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "2", Line: 1},
				{Type: constants.TOKEN_RIGHT_PAREN, Lexeme: ")", Line: 1},
				{Type: constants.TOKEN_EOF, Lexeme: "", Line: 1},
			},
		},
		{
			input: "assert (x * 6.00)",
			expected: []constants.Token{
				{Type: constants.TOKEN_VARIABLE, Lexeme: "assert", Line: 1},
				{Type: constants.TOKEN_LEFT_PAREN, Lexeme: "(", Line: 1},
				{Type: constants.TOKEN_VARIABLE, Lexeme: "x", Line: 1},
				{Type: constants.TOKEN_MULTIPLY, Lexeme: "*", Line: 1},
				{Type: constants.TOKEN_NUMBER, Lexeme: "6.00", Line: 1},
				{Type: constants.TOKEN_RIGHT_PAREN, Lexeme: ")", Line: 1},
				{Type: constants.TOKEN_EOF, Lexeme: "", Line: 1},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		for i, expectedToken := range tt.expected {
			tok := l.NewToken()
			// t.Log(tok.ToString()) // Uncomment this line to see the tokens
			if tok != expectedToken {
				t.Fatalf("test[%d] - token wrong. expected=%q, got=%q", i+1, expectedToken, tok)
			}
		}
	}
}
