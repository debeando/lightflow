package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Loops() {
	name := common.GetArgVal("loop").(string)
	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Loops,
	}

	if len(name) > 0 && ! itr.Exist(name) {
		return
	}

	itr.Run(name, func() {
		f.Index.Pipe = 0
		f.Index.Loop = itr.Index
		f.AutoIncrement()
	})

	log.Info(
		fmt.Sprintf(
			"TASK[%s] LOOPS ET[%s]", // ET is acronym for execution time.
			f.GetTaskName(),
			itr.ExecutionTime,
		),nil )
}
