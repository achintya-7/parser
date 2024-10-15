package main

import (
	"fmt"
	"os"
	"parser/lexer"
	"parser/parser"
	"strings"
)

func main() {
	assertValues := `
		assert (2 + 6) * (x - 3) 
		assert (y + 6) 
		assert z * 2
	`
	valueMap := map[string]float64{
		"x": 3, 
		"y": -6, 
		"z": 0,
	}

	l := lexer.NewLexer(assertValues)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	simplifiedResult, errors, isSuccess := program.PartialEvaluate()
	if !isSuccess {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	fmt.Println("Simplified results :-")
	for _, result := range simplifiedResult {
		fmt.Println(result)
	}

	combinedPartialResults := strings.Join(simplifiedResult, "\n")

	fmt.Println("\nAdding value map for the asserts")
	
	program.SetValueMap(valueMap)
	fmt.Println()

	l = lexer.NewLexer(combinedPartialResults)
	p = parser.NewParser(l)
	program = p.ParseProgram()

	_, errors, isSuccess = program.Evaluate()
	if !isSuccess {
		for _, err := range errors {
			fmt.Println(err)
		}
		fmt.Println("Asserts Failed [X]")
		os.Exit(1)
	}

	fmt.Println("Asserts Passed [✓]")
}
