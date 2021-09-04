package core

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
)

// Print specific variable with value.
func (core *Core) print() {
	names := core.GetPipePrint()
	if len(names) > 0 {
		vars := make(map[string]interface{})

		for _, key := range names {
			vars[key] = core.Variables.Get(key)
		}

		log.Info(
			fmt.Sprintf(
				"%s/%s",
				core.TaskName(),
				core.PipeName(),
			),
			vars,
		)
	}
}
