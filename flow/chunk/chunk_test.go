package chunk_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/chunk"
)

func TestChunk(t *testing.T) {
	type TestChunks struct {
		Step int
		Chunks int
		Chunk int
		Percentage int
		Limit int
	}

	var testChunks = map[int]TestChunks{}
	testChunks[0] = TestChunks{Chunks: 5, Chunk:  0, Percentage:   0, Limit: 2}
	testChunks[1] = TestChunks{Chunks: 5, Chunk:  2, Percentage:  22, Limit: 2}
	testChunks[2] = TestChunks{Chunks: 5, Chunk:  4, Percentage:  44, Limit: 2}
	testChunks[3] = TestChunks{Chunks: 5, Chunk:  6, Percentage:  66, Limit: 2}
	testChunks[4] = TestChunks{Chunks: 5, Chunk:  8, Percentage:  88, Limit: 2}
	testChunks[5] = TestChunks{Chunks: 5, Chunk: 10, Percentage: 100, Limit: 2}

	c := chunk.Chunk{
		Total: 9,
		Limit: 2,
	}
	c.Chunk(func (step int, chunks int, chunk int, pct int){
		if testChunks[step].Chunks != chunks {
			t.Errorf("Expected %d, got %d.", testChunks[step].Chunks, chunks)
		}
		if testChunks[step].Chunk != chunk {
			t.Errorf("Expected %d, got %d.", testChunks[step].Chunk, chunk)
		}
		if testChunks[step].Percentage != pct {
			t.Errorf("Expected %d, got %d.", testChunks[step].Percentage, pct)
		}
		if testChunks[step].Limit != c.Limit {
			t.Errorf("Expected %d, got %d.", testChunks[step].Limit, c.Limit)
		}
	})
}
