package parser

import (
	"parser/lexer"
	"testing"

	"github.com/stretchr/testify/require"
)

type ParserTestCase struct {
	input    string
	expected float64
	succeed  bool
	valueMap map[string]float64
}

func TestParser(t *testing.T) {
	testCases := []ParserTestCase{
		{
			input:    "assert !(1 * 5) == 0",
			expected: 0,
			succeed: true,
		},
		{
			input:    "assert 1 + 2 * 3 == 5",
			expected: 1,
			succeed:  false,
		},
		{
			input:    "assert 1 + 2 * 3 == 7",
			expected: 0,
			succeed: true,
		},
		{
			input:    "assert (1 + 2) * 3 == 9",
			expected: 0,
			succeed: true,
		},
		{
			input:    "assert 1 + 2 == (1 / 2) * 6",
			expected: 0,
			succeed: true,
		},
		{
			input:    "assert 1 + 2 * 3",
			expected: 7,
			succeed: false,
		},
		{
			input:    "assert (2 + 2) * 3",
			expected: 12,
			succeed: false,
		},
		{
			input:    "assert (x + 2) * 3",
			expected: 12,
			succeed: false,
			valueMap: map[string]float64{
				"x": 2,
			},
		},
		{
			input:    "assert x + y == z",
			expected: 0,
			succeed: true,
			valueMap: map[string]float64{
				"x": 1,
				"y": 2,
				"z": 3,
			},
		},
		{
			input:    "assert (x + y) * 4 == z * 2",
			expected: 0,
			succeed: true,
			valueMap: map[string]float64{
				"x": 1,
				"y": 2,
				"z": 6,
			},
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

		result, err := program.Evaluate()
		if testCase.succeed {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
		require.Equal(t, testCase.expected, result)
	}
}
