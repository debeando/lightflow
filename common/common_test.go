package common_test

import (
	"testing"

	"github.com/swapbyt3s/lightflow/common"
)

func TestTrimNewlines(t *testing.T) {
	var cases = []string{
		"abc123",
		"\nabc123\n",
		"\n\nabc123\n\n",
		"\n\nabc\n123\n\n",
	}

	for _, c := range cases {
		t.Log(common.TrimNewlines(c))
	}
}
