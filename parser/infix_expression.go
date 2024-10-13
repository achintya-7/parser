package parser

import (
	"fmt"
	"parser/constants"
)

// InfixExpression is for expressions like 5 + 5, 3 * 3, etc.
type InfixExpression struct {
	Token    constants.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Lexeme }

func (ie *InfixExpression) PartialEvaluate() (string, error) {
	left, err := ie.Left.PartialEvaluate()
	if err != nil {
		return "", err
	}

	right, err := ie.Right.PartialEvaluate()
	if err != nil {
		return "", err
	}

	_, err = IsConstant(left)
	if err != nil {
		return fmt.Sprintf("%s %s %s", left, ie.Operator, right), nil
	}

	_, err = IsConstant(right)
	if err != nil {
		return fmt.Sprintf("%s %s %s", left, ie.Operator, right), nil
	}

	evaluatedValue, err := ie.Evaluate()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", evaluatedValue), nil
}

func (ie *InfixExpression) Evaluate() (float64, error) {
	left, err := ie.Left.Evaluate()
	if err != nil {
		return -1, err
	}

	right, err := ie.Right.Evaluate()
	if err != nil {
		return -1, err
	}

	fmt.Printf("Evaluating: %s %s %s\n", ie.Left.String(), ie.Operator, ie.Right.String())

	switch ie.Operator {
	case "+":
		return left + right, nil
	case "*":
		return left * right, nil
	case "-":
		return left - right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		} else {
			return left / right, nil
		}
	case "==":
		if left == right {
			return 0, nil
		} else {
			return 1, nil
		}
	case "<=":
		if left != right {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return left, nil
	}
}

func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}
