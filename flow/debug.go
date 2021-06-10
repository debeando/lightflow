package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
)

// Debug print all variables in debug mode.
func (f *Flow) debug() {
	log.Debug(
		fmt.Sprintf(
			"%s/%s",
			f.TaskName(),
			f.PipeName(),
		),
		f.Variables.Items,
	)
}
