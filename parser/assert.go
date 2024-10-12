package parser

import (
	"fmt"
	"parser/constants"
)

type AssertStatement struct {
	Token      constants.Token
	Expression Expression
}

func (as *AssertStatement) TokenLiteral() string { return as.Token.Lexeme }

func (as *AssertStatement) Evaluate() (float64, error) {
	value, err := as.Expression.Evaluate()
	if err != nil {
		return 0, err
	}

	if value != 1 {
		return value, nil
	}

	return 1, nil
}

func (as *AssertStatement) String() string {
	return fmt.Sprintf("assert %s", as.Expression.String())
}
