package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Pipes() {
	for _, pipe_name := range args.Pipes() {
		if len(pipe_name) > 0 {
			itr := iterator.Iterator{
				Items: f.Config.Tasks[f.Index.Task].Pipes,
				Name: pipe_name,
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
	}
}
