package parser

import "parser/constants"

type Identifier struct {
	Token constants.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Lexeme }

// todo
func (i *Identifier) Evaluate() (float64, error) {
	// For simplicity, we'll just return 1 for any identifier
	return 1, nil
}

func (i *Identifier) String() string {
	return i.Value
}
