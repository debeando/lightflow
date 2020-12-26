package flow

import (
	"github.com/debeando/lightflow/flow/retry"
)

func (f *Flow) Retry(fn func() bool) {
	r := retry.Retry{
		Attempt: f.GetRetryAttempts(),
		Wait:    f.GetRetryWait(),
	}

	r.Retry(
		func() bool {
			f.Attempt = f.GetRetryAttempts() - r.Attempt + 1
			return fn()
		})
}

func (f *Flow) GetRetryAttempts() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts
}

func (f *Flow) GetRetryWait() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait
}

func (f *Flow) GetRetryExitCode() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.ExitCode
}

func (f *Flow) GetRetryError() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Error

	if len(value) == 0 {
		value = "error"
	}

	return value
}

func (f *Flow) GetRetryStatus() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Status

	if len(value) == 0 {
		value = "status"
	}

	return value
}

func (f *Flow) GetRetryDone() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Done

	if len(value) == 0 {
		value = "done"
	}

	return value
}
