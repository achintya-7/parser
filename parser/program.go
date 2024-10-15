package parser

import (
	"fmt"
)

var (
	ValueMap map[string]float64
)

type Program struct {
	Statements []Statement
}

func (p *Program) SetValueMap(vm map[string]float64) {
	ValueMap = vm
}

func (p *Program) Evaluate() ([]float64, []error, bool) {
	var results []float64
	var errs []error
	success := true

	for i, stmt := range p.Statements {
		fmt.Printf("\nEvaluating statement %d: %s\n", i + 1, stmt.String())

		result, err := stmt.Evaluate()
		fmt.Printf("Result: %f\n\n", result)
		if err != nil {
			success = false
			errs = append(errs, fmt.Errorf("error evaluating statement %d: %s", i, err))
		}

		results = append(results, result)
	}

	return results, errs, success
}

func (p *Program) PartialEvaluate() ([]string, []error, bool) {
	var results []string
	var errs []error
	success := true

	for i, stmt := range p.Statements {
		fmt.Printf("\nEvaluating statement %d: %s\n", i + 1, stmt.String())

		result, err := stmt.PartialEvaluate()
		fmt.Printf("Result: %s\n\n", result)
		if err != nil {
			success = false
			errs = append(errs, fmt.Errorf("error evaluating statement %d: %s", i, err))
		}

		results = append(results, result)
	}

	return results, errs, success
}
