package parser

import (
	"fmt"
	"parser/constants"
	"parser/lexer"
	"strconv"
)

type Parser struct {
	l         lexer.Lexerer
	curToken  constants.Token
	peekToken constants.Token
	errors    []string
}

func NewParser(l lexer.Lexerer) Parserer {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NewToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) printAST(node any) {
	switch n := node.(type) {
	case *Program:
		fmt.Println("Program:")
		for _, stmt := range n.Statements {
			fmt.Printf("  -> %s\n", stmt.String())
		}
	case Expression:
		fmt.Printf("Expression: %s\n", n.String())
	default:
		fmt.Println("Unknown node type")
	}
}

func (p *Parser) ParseProgram() ProgramEvaluator {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != constants.TOKEN_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Lexeme {
	case "assert":
		return p.parseAssertStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssertStatement() *AssertStatement {
	stmt := &AssertStatement{Token: p.curToken}

	p.nextToken()

	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns(p.curToken.Type)
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns(p.peekToken.Type)
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	p.printAST(leftExp)

	return leftExp
}

func (p *Parser) prefixParseFns(tokenType constants.TokenType) func() Expression {
	switch tokenType {
	case constants.TOKEN_VARIABLE:
		return p.parseVariable
	case constants.TOKEN_NUMBER:
		return p.parseNumberLiteral
	case constants.TOKEN_NOT:
		return p.parsePrefixExpression
	case constants.TOKEN_LEFT_PAREN:
		return p.parseGroupedExpression
	default:
		return nil
	}
}

func (p *Parser) infixParseFns(tokenType constants.TokenType) func(Expression) Expression {
	switch tokenType {
	case constants.TOKEN_PLUS, constants.TOKEN_MINUS, constants.TOKEN_DIVIDE, constants.TOKEN_MULTIPLY, constants.TOKEN_DOUBLE_EQUAL, constants.TOKEN_NOT_EQUAL:
		return p.parseInfixExpression
	default:
		return nil
	}
}

func (p *Parser) parseVariable() Expression {
	return &Variable{Token: p.curToken, Value: p.curToken.Lexeme}
}

func (p *Parser) parseNumberLiteral() Expression {
	lit := &NumberLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Lexeme, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Lexeme)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Lexeme,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Lexeme,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(constants.TOKEN_RIGHT_PAREN) {
		return nil
	}

	return exp
}

func (p *Parser) noPrefixParseFnError(t constants.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %v found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t constants.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekTokenIs(t constants.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t constants.TokenType) {
	msg := fmt.Sprintf("expected next token to be %v, got %v instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Operator precedence
const (
	_ int = iota
	LOWEST
	EQUALS    // ==
	NOT_EQUAL // <=
	SUBTRACT  // -
	SUM       // +
	PRODUCT   // *
	DIVISION  // /
	PREFIX    // -X or !X
)

var precedences = map[constants.TokenType]int{
	constants.TOKEN_DOUBLE_EQUAL: EQUALS,
	constants.TOKEN_NOT_EQUAL:    NOT_EQUAL,
	constants.TOKEN_MINUS:        SUBTRACT,
	constants.TOKEN_PLUS:         SUM,
	constants.TOKEN_MULTIPLY:     PRODUCT,
	constants.TOKEN_DIVIDE:       DIVISION,
	constants.TOKEN_NOT:          PREFIX,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
