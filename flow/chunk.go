package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/chunk"
)

// Chunks is the handle method for generic Chunk package.
func (f *Flow) Chunks() {
	if f.IsValidChunk() {
		c := chunk.Chunk{
			Total: f.GetChunkTotal(),
			Limit: f.GetChunkLimit(),
		}

		c.Chunk(func(step int, chunks int, offset int, pct int) {
			log.Info(
				fmt.Sprintf(
					"TASK[%s] SUB TASK[%s] PIPE CHUNK[%d%%]",
					f.TaskName(),
					f.SubTaskName(),
					pct,
				), nil)

			f.Variables.Set(map[string]interface{}{
				"offset": offset,
				"limit":  c.Limit,
			})

			f.Execute()
		})
	} else {
		f.Execute()
	}
}

// GetChunkTotal get the total of chunks from config or variables.
func (f *Flow) GetChunkTotal() int {
	// Aqui hay que poner un interface to int y nos ahorramos validaciones.
	if total := f.Variables.Get("chunk.total"); total != nil {
		if value, ok := total.(string); ok {
			if value := common.StringToInt(value); value > 0 {
				f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total = value
			}
		}
	}

	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total
}

// GetChunkLimit get the limit of chunks from config or variables.
func (f *Flow) GetChunkLimit() int {
	limit := f.Variables.Get("chunk.limit")

	if limit != nil {
		if value, ok := limit.(string); ok {
			if value := common.StringToInt(value); value > 0 {
				f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit = value
			}
		}
	}

	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit
}

// IsValidChunk is the sanity check for settings.
func (f *Flow) IsValidChunk() bool {
	if f.GetChunkTotal() < 2 {
		return false
	}
	if f.GetChunkLimit() < 2 {
		return false
	}

	return true
}
