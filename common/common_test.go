package common_test

import (
	"testing"

	"github.com/debeando/lightflow/common"
)

func TestTrimNewlines(t *testing.T) {
	type TestTexts struct {
		Text   string
		Result string
	}

	var testTexts = map[int]TestTexts{}
	testTexts[0] = TestTexts{Text: "abc123", Result: "abc123"}
	testTexts[1] = TestTexts{Text: "\nabc123\n", Result: "abc123"}
	testTexts[2] = TestTexts{Text: "\n\nabc123\n\n", Result: "abc123"}
	testTexts[3] = TestTexts{Text: "\n\nabc\n123\n\n", Result: "abc\n123"}

	for index, _ := range testTexts {
		if common.TrimNewlines(testTexts[index].Text) != testTexts[index].Result {
			t.Errorf("Expected %s, got %s.", testTexts[index].Result, common.TrimNewlines(testTexts[index].Text))
		}
	}
}
