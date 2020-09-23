package retry

import (
	"time"
)

func Retry(attempts int, wait time.Duration, fn func() bool) bool {
	if r := fn(); r == true {
		if attempts--; attempts > 0 {
			time.Sleep(wait)
			return Retry(attempts, wait, fn)
		}
	}

	return false
}
