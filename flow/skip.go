package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
)

// Skip evaluate condition to set skip flag.
func (f *Flow) skip() {
	if len(f.GetProperty("Skip")) == 0 {
		return
	}

	expression := f.Render(f.GetProperty("Skip"))

	f.Skip = evaluate.Expression(expression)
	debug_vars := make(map[string]interface{})
	debug_vars["Expression"] = f.GetProperty("Skip")
	debug_vars["Rendered"] = expression
	debug_vars["Result"] = f.Skip

	f.Variables.Set(map[string]interface{}{
		"skip": f.Skip,
	})

	log.Info(
		fmt.Sprintf(
			"%s/%s/%s Skip: %s => %s => %#v",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
			debug_vars["Expression"],
			debug_vars["Rendered"],
			debug_vars["Result"],
		),
		nil,
	)
}