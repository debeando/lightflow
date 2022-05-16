package core

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/evaluate"
	"github.com/debeando/lightflow/plugins/retry"
	"github.com/debeando/lightflow/plugins/template"
)

func (core *Core) Retry(fn func()) {
	r := retry.Retry{
		Attempt: core.GetRetryAttempts(),
		Wait:    core.GetRetryWait(),
	}

	r.Retry(
		func() bool {
			core.Attempt = core.GetRetryAttempts() - r.Attempt + 1
			fn()

			core.PrintRetry()
			if core.Attempt == core.GetRetryAttempts() && core.EvalRetry() {
				log.Warning(
					fmt.Sprintf(
						"%s/%s Retry end, Attempts exhausted",
						core.TaskName(),
						core.PipeName(),
					), nil)
			}
			return core.EvalRetry()
		})
}

// PrintRetry show the retry progress.
func (core *Core) PrintRetry() {
	if core.GetRetryAttempts() < 1 {
		return
	}

	log.Info(
		fmt.Sprintf(
			"%s/%s Retry %d/%d",
			core.TaskName(),
			core.PipeName(),
			core.Attempt,
			core.GetRetryAttempts(),
		), nil)
}

func (core *Core) EvalRetry() bool {
	if core.GetRetryWait() == 0 {
		return false
	}

	if core.GetRetryAttempts() == 0 {
		return false
	}

	expression, err := template.Render(core.GetRetryExpression(), core.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return evaluate.Expression(expression)
}
