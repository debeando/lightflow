package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Tasks() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks,
		Name:  args.Task(),
	}

	itr.Run(func() bool {
		f.Index.Subtask = 0
		f.Index.Pipe = 0
		f.Index.Task = itr.Index
		f.Subtask()

		return false
	})

	log.Info(fmt.Sprintf(
		"TASKS ET[%s]", // ET is acronym for execution time.
		itr.ExecutionTime,
	), nil)
}
