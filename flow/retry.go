package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
	"github.com/debeando/lightflow/flow/retry"
	"github.com/debeando/lightflow/flow/template"
)

func (f *Flow) Retry(fn func() bool) {
	r := retry.Retry{
		Attempt: f.GetRetryAttempts(),
		Wait:    f.GetRetryWait(),
	}

	r.Retry(
		func() bool {
			f.Attempt = f.GetRetryAttempts() - r.Attempt + 1
			fn()

			f.PrintRetry()
			return f.EvalRetry()
		})
}

// PrintRetry show the retry progress.
func (f *Flow) PrintRetry() {
	if f.GetRetryAttempts() < 1 {
		return
	}

	log.Info(
		fmt.Sprintf(
			"TASK[%s] SUB TASK[%s] PIPE[%s] RETRY[%d/%d]",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
			f.Attempt,
			f.GetRetryAttempts(),
		), nil)
}

func (f *Flow) EvalRetry() bool {
	if f.GetRetryWait() == 0 {
		return false
	}

	if f.GetRetryAttempts() == 0 {
		return false
	}

	expression, err := template.Render(f.GetRetryExpression(), f.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return evaluate.Expression(expression)
}

func (f *Flow) GetRetryAttempts() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts
}

func (f *Flow) GetRetryWait() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait
}

func (f *Flow) GetRetryExpression() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Expression
}
