package core_test

import (
	"os"
	"testing"

	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/core"
	"github.com/debeando/lightflow/variables"
)

var f = core.Core{}

func TestMain(m *testing.M) {
	core.Variables = *variables.Load()
	core.Config = *config.Load()
	core.Config.Read("../tests/core/chunk.yaml")

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
	// 	core.Variables.Set(map[string]interface{}{
	// 		"chunk.total": value,
	// 	})
	// 	t.Log("V:", core.Variables.Get("chunk.total"))
	// 	t.Log("M:", core.GetChunkTotal())
	// }

	core.Index.Task = 0
	core.Index.Subtask = 0
	core.Index.Pipe = 0
	core.Chunks()
}
