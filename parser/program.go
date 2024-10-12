package parser

import (
	"fmt"
	"strings"
)

var (
	ValueMap map[string]float64
)

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) SetValueMap(vm map[string]float64) {
	ValueMap = vm
}

func (p *Program) Evaluate() (float64, error) {
	var result float64
	var err error
	for i, stmt := range p.Statements {
		fmt.Printf("Evaluating statement %d: %s\n", i, stmt.String())

		result, err = stmt.Evaluate()
		if err != nil {
			if strings.Contains(err.Error(), "assertion failed") {
				return result, err
			}

			return -1, err
		}
	}
	return result, nil
}
