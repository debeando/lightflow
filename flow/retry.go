package flow

import (
	"github.com/swapbyt3s/lightflow/flow/retry"
)

func (f *Flow) Retry(fn func() bool) {
	retry.Retry(
		f.GetRetryAttempts(),
		f.GetRetryWait(),
		func () bool {
			return fn()
		})
}
