package flow

import (
	"fmt"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/chunk"
	"github.com/swapbyt3s/lightflow/variables"
)

func (f *Flow) Chunks() {
	var v = variables.Load()

	if f.IsValidChunk() {
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

			v.Set(map[string]interface{}{
				"offset": offset,
				"limit": c.Limit,
			})

			f.Retry()
		})
	} else {
		f.Retry()
	}
}
