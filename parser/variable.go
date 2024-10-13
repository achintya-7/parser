package parser

import (
	"fmt"
	"parser/constants"
)

type Variable struct {
	Token constants.Token
	Value string
}

func (v *Variable) TokenLiteral() string { return v.Token.Lexeme }

func (v *Variable) Evaluate() (float64, error) {
	value, ok := ValueMap[v.Value]
	if ok {
		return value, nil
	} else {
		return 0, fmt.Errorf("unknown variable: %s", v.Value)
	}
}

func (v *Variable) PartialEvaluate() (string, error) {
	return v.Value, nil
}

func (v *Variable) String() string {
	return v.Value
}
