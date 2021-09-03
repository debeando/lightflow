package retry

import (
	"time"
)

// Retry save the common variables:
type Retry struct {
	Attempt int // Current attempt
	Wait    int // Wait between attempt
}

// Retry is a method to call many times until satify return value.
func (r *Retry) Retry(fn func() bool) bool {
	if c := fn(); c == true {
		if r.Attempt--; r.Attempt > 0 {
			time.Sleep(time.Duration(r.Wait) * time.Second)
			return r.Retry(fn)
		}
	}

	return false
}
