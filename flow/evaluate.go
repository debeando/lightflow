package flow

import (
	"fmt"
	"os"
	"strings"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
)

func (f *Flow) evaluate() {
	evals := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Evaluate

	for _, eval := range evals {
		if len(eval.Expression) == 0 {
			log.Warning(
				fmt.Sprintf(
					"%s/%s/%s Evaluate expression for '%s' is empty.",
					f.TaskName(),
					f.SubTaskName(),
					f.PipeName(),
					eval.Action,
				),
				nil,
			)
		}

		switch strings.ToUpper(eval.Action) {
		case "TERMINATE":
			message, result := f.evaluateGeneral("Terminate", eval.Expression)

			if f.When == true && result {
				evaluateLog(eval.Level, message)

				os.Exit(eval.ExitCode)
			}
		case "SKIP":
			message, result := f.evaluateGeneral("Skip", eval.Expression)

			if f.When == true && result {
				evaluateLog(eval.Level, message)

				f.Skip = result
				f.Variables.Set(map[string]interface{}{
					"skip": f.Skip,
				})
				return
			}
		case "LOGGING":
			message, result := f.evaluateGeneral("Looging", eval.Expression)

			if f.When == true && result {
				evaluateLog(eval.Level, message)
			}
		case "WHEN":
			message, result := f.evaluateGeneral("When", eval.Expression)

			if f.In == false && ! result {
				evaluateLog(eval.Level, message)
				f.When = false
				return
			}
		default:
			log.Warning(
				fmt.Sprintf(
					"%s/%s/%s Evaluate action invalid: %s",
					f.TaskName(),
					f.SubTaskName(),
					f.PipeName(),
					eval.Action,
				),
				nil,
			)
		}
	}
}

func (f *Flow) evaluateGeneral(eval_type string, expression string) (string, bool) {
	rendered := f.Render(expression)
	result := evaluate.Expression(rendered)

	message := fmt.Sprintf(
		"%s/%s/%s %s: %s => %s => %#v",
		f.TaskName(),
		f.SubTaskName(),
		f.PipeName(),
		eval_type,
		expression,
		rendered,
		result,
	)

	return message, result
}

func evaluateLog(eval_type string, message string) {
	switch strings.ToUpper(eval_type) {
	case "INFO":
		log.Info(
			message,
			nil,
		)
	case "WARNING":
		log.Warning(
			message,
			nil,
		)
	case "ERROR":
		log.Error(
			message,
			nil,
		)
	case "DEBUG":
		log.Debug(
			message,
			nil,
		)
	}
}
