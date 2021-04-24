package flow

import (
	// "fmt"
	// "strings"

	// "github.com/debeando/lightflow/common/log"
)

func (f *Flow) RunPipe() bool {
	f.evaluate()
	if f.In == false && f.When {
		// fmt.Println("> try to run pipe...")
		f.In = true
		return true
	}
		// if f.When() {
	return false
	// evals := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Evaluate

	// for _, eval := range evals {
	// 	if len(eval.Expression) == 0 {
	// 		log.Warning(
	// 			fmt.Sprintf(
	// 				"%s/%s/%s Evaluate expression for '%s' is empty.",
	// 				f.TaskName(),
	// 				f.SubTaskName(),
	// 				f.PipeName(),
	// 				eval.Action,
	// 			),
	// 			nil,
	// 		)
	// 	}

	// 	switch strings.ToUpper(eval.Action) {
	// 	case "WHEN":
	// 		message, result := f.evaluateGeneral("When", eval.Expression)

	// 		if ! result {
	// 			evaluateLog(eval.Level, message)
	// 			return false
	// 		}
	// 	}
	// }

	// return true
}
