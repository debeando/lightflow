package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Subtask() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Subtask,
		Name: args.Subtask(),
	}

	itr.Run(func() bool {
		f.Index.Pipe = 0
		f.Index.Subtask = itr.Index
		f.Skip = false

		if err := f.AutoIncrement(); err != nil {
			log.Error(err.Error(), nil)
		}
		return false
	})

	log.Info(
		fmt.Sprintf(
			"TASK[%s] SUB TASK ET[%s]", // ET is acronym for execution time.
			f.TaskName(),
			itr.ExecutionTime,
		), nil)
}
