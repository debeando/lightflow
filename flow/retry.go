package flow

/*
 * TODO: Meter dentro del command el retry y no al revez.
 */

import (
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/retry"
	"github.com/swapbyt3s/lightflow/variables"
)

func (f *Flow) Retry() {
	retry.Retry(
		f.GetRetryAttempts(),
		f.GetRetryWait(),
		func() bool {
			return f.ExecuteCommand()
		})

	f.PrintDebugVariables()
}

func (f *Flow) PrintDebugVariables() {
	for variable, value := range variables.Load().Items {
		switch variable {
		case "exit_code":
			if value.(int) > 0 {
				log.Error(f.GetTitle(), map[string]interface{}{
					"Exit Code": value,
				})
			}
			break
		default:
			log.Debug(f.GetTitle(), map[string]interface{}{
				variable: value,
			})
		}
	}
}
