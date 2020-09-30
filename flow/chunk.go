package flow

import (
	"fmt"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/chunk"
)

func (f *Flow) Chunks() {
	if f.IsValidChunk() {
		log.Info("Starting chunk loop...", nil)
		c := chunk.Chunk{
			Total: f.GetChunkTotal(),
			Limit: f.GetChunkLimit(),
		}

		c.Chunk(func (step int, chunks int, offset int, pct int){
			log.Info(
				f.GetTitle(),
				map[string]interface{}{
					"Chunk Percentage": fmt.Sprintf("%d%%", pct),
			})

			f.Variables.Set(map[string]interface{}{
				"offset": offset,
				"limit": c.Limit,
			})

			f.Execute()
		})
	} else {
		f.Execute()
	}
}
