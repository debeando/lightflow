package expression_test

import (
	"testing"

	"github.com/swapbyt3s/lightflow/flow/expression"
	"github.com/swapbyt3s/lightflow/variables"
)

type Case struct {
	Expression string
	Result bool
}

func TestEvaluate(t *testing.T) {
	variables.Load().Set(map[string]interface{}{
		"X": "Z",
		"Y": "Z",
	})

	var cases = []Case{
		{ Expression: `X == "Z"`,      Result: true },
		{ Expression: `"Z" == Y`,      Result: true },
		{ Expression: `1 == 1`,        Result: true },
		{ Expression: `1 == 2`,        Result: false },
		{ Expression: `"A" == "A"`,    Result: true },
		{ Expression: `"A" == "B"`,    Result: false },
		{ Expression: `"1A" == "1A"`,  Result: true },
		{ Expression: `"1A" == "2B"`,  Result: false },
		{ Expression: `"!A" == "!A"`,  Result: true },
		{ Expression: `"!A" == "!@"`,  Result: false },
	}

	for _, c := range cases {
		if r := expression.Evaluate(c.Expression); r != c.Result {
			t.Errorf("Expression [%s] equeal %t; want %t", c.Expression, r, c.Result)
		}
	}
}
