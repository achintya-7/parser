package parser

import "parser/constants"

// AST Node Interface
type (
	Node interface {
		TokenLiteral() string
		Evaluate() (float64, error)
		PartialEvaluate() (string, error)
	}

	Statement interface {
		Node
		String() string
	}

	Expression interface {
		Node
		String() string
	}
)

// Parser Interface
type Parserer interface {
	nextToken()
	Errors() []string
	printAST(any)
	ParseProgram() ProgramEvaluator
	parseStatement() Statement
	parseAssertStatement() *AssertStatement
	parseExpression(int) Expression
	prefixParseFns(tokenType constants.TokenType) func() Expression
	infixParseFns(tokenType constants.TokenType) func(Expression) Expression
	parseVariable() Expression
	parseNumberLiteral() Expression
	parsePrefixExpression() Expression
	parseInfixExpression(left Expression) Expression
	parseGroupedExpression() Expression
	noPrefixParseFnError(t constants.TokenType)
	expectPeek(t constants.TokenType) bool
	peekTokenIs(t constants.TokenType) bool
	peekError(t constants.TokenType)
	peekPrecedence() int
	curPrecedence() int
}

// Program Evaluator Interface
type ProgramEvaluator interface {
	SetValueMap(map[string]float64)
	Evaluate() ([]float64, []error, bool)
	PartialEvaluate() ([]string, []error, bool)
}

// AST PrintVisitor Interface
type PrintVisitor interface {
	VisitProgram(*Program, int)
	VisitExpression(Expression, int)
}
