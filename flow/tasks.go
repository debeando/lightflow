package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Tasks() {
	name := common.GetArgVal("task").(string)
	itr := iterator.Iterator{
		Items: f.Config.Tasks,
	}

	itr.Run(name, func() {
		f.Index.Loop = 0
		f.Index.Pipe = 0
		f.Index.Task = itr.Index
		f.Loops()
	})

	log.Info(fmt.Sprintf(
		"TASKS ET[%s]", // ET is acronym for execution time.
		itr.ExecutionTime,
	), nil)
}
