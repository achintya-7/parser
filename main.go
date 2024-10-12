package main

import (
	"fmt"
	"parser/lexer"
	"parser/parser"
)

func main() {
	testCases := []string{
		// "assert 1 + 2 * 3",
		// "assert (2 + 2) * 3",
		"assert !(1 * 5) == 0",
		// "assert 1 + 2 * 3 == 5",
		// "assert 1 + 2 * 3 = 6",
		// "assert (1 + 2) * 3 = 9",
		// "assert 1 + 2 == (1 / 2) * 6",
		// "assert 1 + 2 * 3 = 5",
		// "assert x + y = z",
	}

	for _, testCase := range testCases {
		fmt.Printf("Testing: %s\n", testCase)
		l := lexer.NewLexer(testCase)
		p := parser.NewParser(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			fmt.Println("Parser errors:")
			for _, err := range p.Errors() {
				fmt.Printf("\t%s\n", err)
			}
			continue
		}

		result, err := program.Evaluate()
		if err != nil {
			fmt.Printf("Evaluation error: %s\n", err)
		} else {
			fmt.Printf("Result: %v\n", result)
		}

		fmt.Println()
	}
}
