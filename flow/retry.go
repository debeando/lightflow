package flow

import (
	"github.com/debeando/lightflow/flow/retry"
)

func (f *Flow) Retry(fn func() bool) {
	retry.Retry(
		f.GetRetryAttempts(),
		f.GetRetryWait(),
		func () bool {
			return fn()
		})
}
