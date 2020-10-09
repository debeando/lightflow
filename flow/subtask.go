package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Subtask() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Subtask,
		Name: common.GetArgVal("subtask").(string),
	}

	itr.Run(func() {
		f.Index.Pipe = 0
		f.Index.Subtask = itr.Index
		f.AutoIncrement()
	})

	log.Info(
		fmt.Sprintf(
			"TASK[%s] SUB TASK ET[%s]", // ET is acronym for execution time.
			f.GetTaskName(),
			itr.ExecutionTime,
		), nil)
}
