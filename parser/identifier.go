package parser

import (
	"fmt"
	"parser/constants"
)

type Identifier struct {
	Token constants.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Lexeme }

// todo
func (i *Identifier) Evaluate() (float64, error) {
	value, ok := ValueMap[i.Value]
	if ok {
		return value, nil
	} else {
		return 0, fmt.Errorf("unknown identifier: %s", i.Value)
	}
}

func (i *Identifier) PartialEvaluate() (string, error) {
	return i.Value, nil
}

func (i *Identifier) String() string {
	return i.Value
}
