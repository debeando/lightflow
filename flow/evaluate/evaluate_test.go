package evaluate_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/evaluate"
)

func TestExpressions(t *testing.T) {
	type TestExpressions struct {
		Formula string
		Valid   bool
	}

	var testExpressions = map[int]TestExpressions{}

	testExpressions[0] = TestExpressions{Formula: "", Valid: false}
	testExpressions[1] = TestExpressions{Formula: "1 == 1", Valid: true}
	testExpressions[2] = TestExpressions{Formula: "1 == 0", Valid: false}
	testExpressions[3] = TestExpressions{Formula: "1 = 0", Valid: false}
	testExpressions[4] = TestExpressions{Formula: "1 = 1", Valid: false}
	testExpressions[5] = TestExpressions{Formula: "(1 + 3) > 5", Valid: false}
	testExpressions[6] = TestExpressions{Formula: "(1 + 3) >= 4", Valid: true}

	for index, _ := range testExpressions {
		t.Log(testExpressions[index].Formula)
		t.Log(evaluate.Expression(testExpressions[index].Formula))
	}
}
