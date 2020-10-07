package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Tasks() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks,
		Name: common.GetArgVal("task").(string),
	}

	itr.Run(func() {
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
