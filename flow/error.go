package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/evaluate"
)

// Error evaluate expression to identify any error or suggest error.
func (f *Flow) error() {
	error := f.GetProperty("Error")
	if len(error) == 0 {
		error = "{{ .exit_code }} != 0 || {{ len .error }} > 0"
	}

	expression := f.Render(error)
	result := evaluate.Expression(expression)

	debug_vars := make(map[string]interface{})
	debug_vars["Expression"] = error
	debug_vars["Rendered"] = expression
	debug_vars["Result"] = result

	if result {
		log.Error(
			fmt.Sprintf(
				"%s/%s Error: %s => %s => %#v",
				f.TaskName(),
				f.PipeName(),
				debug_vars["Expression"],
				debug_vars["Rendered"],
				debug_vars["Result"],
			),
			nil,
		)
	}
}
