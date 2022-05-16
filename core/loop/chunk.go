package core

// import (
// 	"fmt"

// 	"github.com/debeando/lightflow/common"
// 	"github.com/debeando/lightflow/common/log"
// 	"github.com/debeando/lightflow/plugins/chunk"
// 	"github.com/debeando/lightflow/plugins/duration"
// )

// // Chunks is the handle method for generic Chunk package.
// func (core *Core) Chunks() {
// 	var et string

// 	if core.IsValidChunk() {
// 		c := chunk.Chunk{
// 			Total: core.GetChunkTotal(),
// 			Limit: core.GetChunkLimit(),
// 		}

// 		c.Chunk(func(step int, chunks int, offset int, pct int) {
// 			log.Info(
// 				fmt.Sprintf(
// 					"%s/%s Progress %d%%",
// 					core.TaskName(),
// 					core.PipeName(),
// 					pct,
// 				), nil)

// 			core.Variables.Set(map[string]interface{}{
// 				"chunk_limit":  c.Limit,
// 				"chunk_total":  c.Total,
// 				"chunk_offset": offset,
// 				"chunk_end":    offset + c.Limit,
// 				"chunk_step":   step,
// 			})

// 			// core.Execute()
// 		})
// 		et = c.ExecutionTime
// 	} else {
// 		et = duration.Start(func() {
// 			// core.Execute()
// 		})
// 	}

// 	log.Info(
// 		fmt.Sprintf(
// 			"%s/%s Finished %s",
// 			core.TaskName(),
// 			core.PipeName(),
// 			et,
// 		), nil)
// }

// // GetChunkTotal get the total of chunks from config or variables.
// func (core *Core) GetChunkTotal() int {
// 	// Aqui hay que poner un interface to int y nos ahorramos validaciones.
// 	if total := core.Variables.Get("chunk.total"); total != nil {
// 		if value, ok := total.(string); ok {
// 			if value := common.StringToInt(value); value > 0 {
// 				core.Config.Pipes[core.Index.Pipe].Chunk.Total = value
// 			}
// 		}
// 	}

// 	return core.Config.Pipes[core.Index.Pipe].Chunk.Total
// }

// // GetChunkLimit get the limit of chunks from config or variables.
// func (core *Core) GetChunkLimit() int {
// 	limit := core.Variables.Get("chunk.limit")

// 	if limit != nil {
// 		if value, ok := limit.(string); ok {
// 			if value := common.StringToInt(value); value > 0 {
// 				core.Config.Pipes[core.Index.Pipe].Chunk.Limit = value
// 			}
// 		}
// 	}

// 	return core.Config.Pipes[core.Index.Pipe].Chunk.Limit
// }

// // IsValidChunk is the sanity check for settings.
// func (core *Core) IsValidChunk() bool {
// 	if core.GetChunkTotal() < 2 {
// 		return false
// 	}
// 	if core.GetChunkLimit() < 2 {
// 		return false
// 	}

// 	return true
// }
