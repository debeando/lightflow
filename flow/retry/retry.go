package retry

import (
	"time"
)

func Retry(attempts int, wait int, fn func() bool) bool {
	if r := fn(); r == true {
		if attempts--; attempts > 0 {
			time.Sleep(time.Duration(wait) * time.Second)
			return Retry(attempts, wait, fn)
		}
	}

	return false
}
