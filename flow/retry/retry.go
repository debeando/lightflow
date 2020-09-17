package retry

import (
	"time"
)

func Retry(attempts int, fn func() bool) bool {
	if r := fn(); r == true {
		if attempts--; attempts > 0 {
			time.Sleep(3 * time.Second)
			return Retry(attempts, fn)
		}
	}

	return false
}
