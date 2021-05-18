package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
)

// When a condition is true allow execute pipe.
func (f *Flow) when() bool {
	if len(f.GetProperty("When")) == 0 {
		return true
	}

	expression := f.Render(f.GetProperty("When"))
	value := evaluate.Expression(expression)

	debug_vars := make(map[string]interface{})
	debug_vars["Expression"] = f.GetProperty("When")
	debug_vars["Rendered"] = expression
	debug_vars["Result"] = value

	log.Info(
		fmt.Sprintf(
			"%s/%s/%s When: %s => %s => %#v",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
			debug_vars["Expression"],
			debug_vars["Rendered"],
			debug_vars["Result"],
		),
		nil,
	)

	return value
}
