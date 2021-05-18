package chunk

import (
	"math"

	"github.com/debeando/lightflow/flow/duration"
)

// Chunk save the common settings:
type Chunk struct {
	Total         int // Total of items.
	Limit         int // Items for each chunk.
	ExecutionTime string
}

// Chunk is the main handler for loop .
func (c *Chunk) Chunk(fn func(step int, chunks int, offset int, percentage int)) {
	// Calc total number of iterations:
	chunks := int(math.Ceil(float64(c.Total) / float64(c.Limit)))

	c.ExecutionTime = duration.Start(func() {
		// Generic iteration like Yield:
		for step := 0; step <= chunks; step++ {
			offset := (step * c.Limit)
			pct := percentage(c.Total, offset)
			fn(step, chunks, offset, pct)
		}
	})
}

// percentage get the overall process chunk in percentage format.
func percentage(total int, position int) int {
	pct := ((100 * position) / total)

	if pct < 0 {
		return 0
	} else if pct > 100 {
		return 100
	}

	return pct
}
