package flow_test

import (
	"os"
	"testing"

	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow"
	"github.com/debeando/lightflow/variables"
)

var f = flow.Flow{}

func TestMain(m *testing.M) {
	f.Variables = *variables.Load()
	f.Config = *config.Load()
	f.Config.Read("../tests/flow/chunk.yaml")

	os.Exit(m.Run())
}

func TestGetChunkTotal(t *testing.T) {
	// var values = []interface{}{
	// 	"10",
	// 	"3",
	// 	"1",
	// 	"12",
	// }

	// for _, value := range values {
	// 	f.Variables.Set(map[string]interface{}{
	// 		"chunk.total": value,
	// 	})
	// 	t.Log("V:", f.Variables.Get("chunk.total"))
	// 	t.Log("M:", f.GetChunkTotal())
	// }

	f.Index.Task = 0
	f.Index.Subtask = 0
	f.Index.Pipe = 0
	f.Chunks()
}
