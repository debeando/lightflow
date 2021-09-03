package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
)

// Print specific variable with value.
func (f *Flow) print() {
	names := f.GetPipePrint()
	if len(names) > 0 {
		vars := make(map[string]interface{})

		for _, key := range names {
			vars[key] = f.Variables.Get(key)
		}

		log.Info(
			fmt.Sprintf(
				"%s/%s",
				f.TaskName(),
				f.PipeName(),
			),
			vars,
		)
	}
}
