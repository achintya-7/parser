package parser

import (
	"parser/lexer"
	"testing"

	"github.com/stretchr/testify/require"
)

type ParserTestCase struct {
	input                  string
	expectedResults        []float64
	expectedPartialResults []string
	succeed                bool
	valueMap               map[string]float64
}

func TestParser(t *testing.T) {
	testCases := []ParserTestCase{
		{
			input:           "assert !(1 * 5) == 0",
			expectedResults: []float64{0},
			succeed:         true,
		},
		{
			input:           "assert 1 + 2 * 3 == 5",
			expectedResults: []float64{1},
			succeed:         false,
		},
		{
			input:           "assert 1 + 2 * 3 == 7",
			expectedResults: []float64{0},
			succeed:         true,
		},
		{
			input:           "assert (1 + 2) * 3 == 9",
			expectedResults: []float64{0},
			succeed:         true,
		},
		{
			input:           "assert 1 + 2 == (1 / 2) * 6",
			expectedResults: []float64{0},
			succeed:         true,
		},
		{
			input:           "assert 1 + 2 * 3",
			expectedResults: []float64{7},
			succeed:         false,
		},
		{
			input:           "assert (2 + 2) * 3",
			expectedResults: []float64{12},
			succeed:         false,
		},
		{
			input:           "assert (x + 2) * 3",
			expectedResults: []float64{12},
			succeed:         false,
			valueMap: map[string]float64{
				"x": 2,
			},
		},
		{
			input:           "assert x + y == z",
			expectedResults: []float64{0},
			succeed:         true,
			valueMap: map[string]float64{
				"x": 1,
				"y": 2,
				"z": 3,
			},
		},
		{
			input:           "assert (x + y) * 4 == z * 2",
			expectedResults: []float64{0},
			succeed:         true,
			valueMap: map[string]float64{
				"x": 1,
				"y": 2,
				"z": 6,
			},
		},
		{
			input:           "assert (x + y) * 4",
			expectedResults: []float64{0},
			succeed:         true,
			valueMap: map[string]float64{
				"x": 2,
				"y": -2,
			},
		},
		{
			input:           "assert !(1 * 5) == 0\n assert 1 + 2 * 3 == 7",
			expectedResults: []float64{0, 0},
			succeed:         true,
		},
		{
			input:           "assert !(1 * 5) == 0\n assert 1 + 2 * 3 == 5",
			expectedResults: []float64{0, 1},
			succeed:         false,
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

		if len(testCase.valueMap) > 0 {
			program.SetValueMap(testCase.valueMap)
		}

		results, _, success := program.Evaluate()
		require.Equal(t, testCase.succeed, success)
		require.Equal(t, testCase.expectedResults, results)
	}
}

func TestPartialEvaluation(t *testing.T) {
	testCases := []ParserTestCase{
		{
			input:                  "assert x * (2 * 3)",
			expectedPartialResults: []string{"assert x * 6.00"},
			succeed:                true,
		},
		{
			input:                  "assert x * (2 * 3) == 6",
			expectedPartialResults: []string{"assert x * 6.00 == 6.00"},
			succeed:                true,
		},
		{
			input:                  "assert x * y + (2 * 3)\n assert (z + 2) * 3",
			expectedPartialResults: []string{"assert x * y + 6.00", "assert z + 2.00 * 3.00"},
			succeed:                true,
		},
		{
			input:                  "assert !(1 * 0) * x + (5 * 6 - y)",
			expectedPartialResults: []string{"assert 1.00 * x + 30.00 - y"},
			succeed:                true,
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

		if len(testCase.valueMap) > 0 {
			program.SetValueMap(testCase.valueMap)
		}

		results, _, success := program.PartialEvaluate()
		require.Equal(t, testCase.succeed, success)
		require.Equal(t, testCase.expectedPartialResults, results)
	}
}

type DualParserTestCase struct {
	initialInput    string
	partialEvaluatedInput    []string
	valueMap        map[string]float64
	expectedResults []float64
}

func TestDualParser(t *testing.T) {
	testCases := []DualParserTestCase{
		{
			initialInput:    "assert x * (2 * 3)",
			partialEvaluatedInput:    []string{"assert x * 6.00"},
			valueMap:        map[string]float64{"x": 0},
			expectedResults: []float64{0},
		},
	}

	for _, testCase := range testCases {
		t.Log("Testing:", testCase.initialInput)
		l := lexer.NewLexer(testCase.initialInput)
		p := NewParser(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Errorf("Parser errors: %v", p.Errors())
			continue
		}

		if len(testCase.valueMap) > 0 {
			program.SetValueMap(testCase.valueMap)
		}

		results, _, success := program.PartialEvaluate()
		require.True(t, success)
		require.Equal(t, testCase.partialEvaluatedInput, results)


		l = lexer.NewLexer(results[0])
		p = NewParser(l)
		program = p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Errorf("Parser errors: %v", p.Errors())
			continue
		}

		parsedResults, _, success := program.Evaluate()
		require.True(t, success)
		require.Equal(t, testCase.expectedResults, parsedResults)
	}
}
