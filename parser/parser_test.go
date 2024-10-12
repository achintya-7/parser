package parser

import (
	"parser/lexer"
	"testing"

	"github.com/stretchr/testify/require"
)

type ParserTestCase struct {
	input           string
	expectedResults []float64
	succeed         bool
	valueMap        map[string]float64
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
