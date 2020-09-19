package retry

import (
	"time"

	"github.com/swapbyt3s/lightflow/registry"
)

func Retry(attempts int, fn func() bool) bool {
	if r := fn(); r == true {
		if attempts--; attempts > 0 {
			time.Sleep(registry.Load().GetRetryWait() * time.Second)
			return Retry(attempts, fn)
		}
	}

	return false
}
