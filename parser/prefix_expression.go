package parser

import (
	"fmt"
	"parser/constants"
)

type PrefixExpression struct {
	Token    constants.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Lexeme }

// todo: verify if this is correct
func (pe *PrefixExpression) Evaluate() (float64, error) {
	right, err := pe.Right.Evaluate()
	if err != nil {
		return 0, err
	}

	switch pe.Operator {
	case "!":
		if right == 0 {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("unknown operator: %s", pe.Operator)
	}
}

func (pe *PrefixExpression) PartialEvaluate() (string, error) {
	right, err := pe.Right.PartialEvaluate()
	if err != nil {
		return "", err
	}

	_, err = IsConstant(right)
	if err != nil {
		return fmt.Sprintf("(%s %s)", pe.Operator, right), nil
	}

	evaluatedValue, err := pe.Evaluate()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", evaluatedValue), nil
}

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String())
}
