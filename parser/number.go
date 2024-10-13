package parser

import (
	"fmt"
	"parser/constants"
)

type NumberLiteral struct {
	Token constants.Token
	Value float64
}

func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Lexeme }

func (nl *NumberLiteral) Evaluate() (float64, error) { return nl.Value, nil }

func (nl *NumberLiteral) PartialEvaluate() (string, error) { return fmt.Sprintf("%.2f", nl.Value), nil }

func (nl *NumberLiteral) String() string { return nl.Token.Lexeme }
