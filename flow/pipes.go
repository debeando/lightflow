package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Pipes() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Pipes,
		Name: common.GetArgVal("pipe").(string),
	}

	itr.Run(func() {
		f.Index.Pipe = itr.Index
		log.Info(f.GetTitle(), nil)
		f.Chunks()
	})

	log.Info(
		fmt.Sprintf(
			"TASK[%s] SUB TASK[%s] PIPES ET[%s]", // ET is acronym for execution time.
			f.GetTaskName(),
			f.GetLoopName(),
			itr.ExecutionTime,
		), nil)
}
