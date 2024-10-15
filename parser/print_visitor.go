package parser

import "fmt"

type PrintVisitorStruct struct{}

func NewPrintVisitor() PrintVisitor {
	return &PrintVisitorStruct{}
}

func (pv *PrintVisitorStruct) VisitProgram(p *Program, indent int) {
	fmt.Println("Program:")
	for _, stmt := range p.Statements {
		printIndent(indent + 1)
		fmt.Println("-> Statement:")
		if expr, ok := stmt.(Expression); ok {
			pv.VisitExpression(expr, indent+2)
		}
	}
}

func (pv *PrintVisitorStruct) VisitExpression(e Expression, indent int) {
	switch expr := e.(type) {
	case *InfixExpression:
		printIndent(indent)
		fmt.Println("InfixExpression:")

		printIndent(indent + 1)
		fmt.Println("Left:")
		pv.VisitExpression(expr.Left, indent+2)

		printIndent(indent + 1)
		fmt.Printf("Operator: %s\n", expr.Operator)

		printIndent(indent + 1)
		fmt.Println("Right:")
		pv.VisitExpression(expr.Right, indent+2)

	case *PrefixExpression:
		printIndent(indent)
		fmt.Println("PrefixExpression:")

		printIndent(indent + 1)
		fmt.Printf("Operator: %s\n", expr.Operator)

		printIndent(indent + 1)
		fmt.Println("Right:")
		pv.VisitExpression(expr.Right, indent+2)

	case *NumberLiteral:
		printIndent(indent)
		fmt.Printf("NumberLiteral: %.2f\n", expr.Value)

	case *Variable:
		printIndent(indent)
		fmt.Printf("Variable: %s\n", expr.Value)

	case *AssertStatement:
		printIndent(indent)
		fmt.Println("AssertStatement:")
		pv.VisitExpression(expr.Expression, indent+1)

	default:
		printIndent(indent)
		fmt.Println("Unknown expression type")
	}
}
