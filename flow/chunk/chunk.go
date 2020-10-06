package chunk

import (
	"math"
)

type Chunk struct {
	Total int
	Limit int
}

func (c *Chunk) Chunk(fn func(step int, chunks int, offset int, percentage int)) {
	chunks := int(math.Ceil(float64(c.Total) / float64(c.Limit)))

	for step := 0; step <= chunks; step++ {
		offset := (step * c.Limit)
		pct := percentage(c.Total, offset)
		fn(step, chunks, offset, pct)
	}
}

func percentage(total int, position int) int {
	pct := ((100 * position) / total)

	if pct < 0 {
		return 0
	} else if pct > 100 {
		return 100
	}

	return pct
}
