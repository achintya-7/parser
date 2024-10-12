package parser

import (
	"fmt"
	"parser/constants"
)

type // InfixExpression is for expressions like 5 + 5, 3 * 3, etc.
InfixExpression struct {
	Token    constants.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Lexeme }

func (ie *InfixExpression) Evaluate() (float64, error) {
	left, err := ie.Left.Evaluate()
	if err != nil {
		return 0, err
	}
	fmt.Println("Left: ", left)

	right, err := ie.Right.Evaluate()
	if err != nil {
		return 0, err
	}
	fmt.Println("Right: ", right)

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
			return 1, nil
		} else {
			return 0, nil
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
