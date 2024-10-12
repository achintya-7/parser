package parser

import (
	"parser/lexer"
	"testing"
)

type ParserTestCase struct {
	input    string
	expected float64
}

func TestParser(t *testing.T) {
	testCases := []ParserTestCase{
		{
			input:    "assert !(1 * 5) == 0",
			expected: 1,
		},
		{
			input:    "assert 1 + 2 * 3 == 5",
			expected: 0,
		},
		{
			input:    "assert 1 + 2 * 3 == 7",
			expected: 1,
		},
		{
			input:    "assert (1 + 2) * 3 == 9",
			expected: 1,
		},
		{
			input:    "assert 1 + 2 == (1 / 2) * 6",
			expected: 1,
		},
		{
			input:    "assert 1 + 2 * 3",
			expected: 7,
		},
		{
			input:    "assert (2 + 2) * 3",
			expected: 12,
		},
	}

	for _, testCase := range testCases {
		t.Log("Testing:", testCase.input)
		l := lexer.NewLexer(testCase.input)
		p := NewParser(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Errorf("Parser errors: %v", p.Errors())
			continue
		}

		result, err := program.Evaluate()
		if err != nil {
			t.Errorf("Evaluation error: %s", err)
		} else if result != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}
