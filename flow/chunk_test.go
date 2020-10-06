package flow_test

import (
	"os"
	"testing"

	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow"
	"github.com/debeando/lightflow/variables"
	// "github.com/debeando/lightflow/flow/chunk"
)

var f = flow.Flow{}

func TestMain(m *testing.M) {
	f.Variables = *variables.Load()
	// f.Config = config.Structure
	f.Config = *config.Load()
	f.Config.Read("../tests/flow/chunk.yaml")
	// f.Config.Read("../tests/empty.yaml")

	os.Exit(m.Run())
}

func TestGetChunkTotal(t *testing.T) {
	var values = []interface{}{
		// "a",
		"10",
		"3",
		1,
		"12",
		// 2,
	}

	// t.Log("a")

	for _, value := range values {
		// t.Log(value)
		f.Variables.Set(map[string]interface{}{
			"chunk.total": value,
		})
		t.Log("V: ", f.Variables.Get("chunk.total"))
		t.Log("M", f.GetChunkTotal())
	}
}
