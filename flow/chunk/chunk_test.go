package chunk_test

import (
	"fmt"
	"testing"

	"github.com/swapbyt3s/lightflow/flow/chunk"
)

func Test(t *testing.T) {
	c := chunk.Chunk{
		Total: 9,
		Limit: 2,
	}

	c.Chunk(func (step int, chunks int, chunk int, pct int){
		fmt.Printf("%d/%d %d%%: %d - %d\n", step, chunks, pct, chunk, c.Limit)
	})
}
