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
		return -1, err
	}

	fmt.Println("Assert value:", value)

	if value != 0 {
		return value, fmt.Errorf("assertion failed: %s", as.Expression.String())
	}

	return value, nil
}

func (as *AssertStatement) PartialEvaluate() (string, error) {
	value, err := as.Expression.PartialEvaluate()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("assert %s", value), nil
}

func (as *AssertStatement) String() string {
	return fmt.Sprintf("assert %s", as.Expression.String())
}
