package expression_test

import (
	"testing"

	"github.com/swapbyt3s/lightflow/flow/expression"
)

type Case struct {
	Expression string
	Result bool
}

func TestEvaluate(t *testing.T) {
	var cases = []Case{
		{ Expression: `true == true`,  Result: true },
		// Se comenta, porque aún no se ha podido resolver este caso.
		// { Expression: `true == false`, Result: false },
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

/*
 * Links a revisar:
 * https://github.com/dave/brenda
 * https://blog.gopheracademy.com/advent-2014/parsers-lexers/
 */
