package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) PipesInArgs() {
	for _, pipeName := range args.Pipes() {
		if f.Valid(pipeName) {
			itr := iterator.Iterator{
				Items: f.Config.Tasks[f.Index.Task].Pipes,
				Name:  pipeName,
			}

			itr.Run(func() bool {
				f.Index.Pipe = itr.Index

				if ! f.Skip {
					log.Info(f.GetTitle(), nil)
					f.Chunks()
				}

				return f.Skip
			})

			if ! f.Skip {
				log.Info(
					fmt.Sprintf(
						"%s/%s Finished %s", // ET is acronym for execution time.
						f.TaskName(),
						f.SubTaskName(),
						itr.ExecutionTime,
					), nil)
			}
		}
	}
}

func (f *Flow) Valid(pipeName string) bool {
	return (len(args.Pipes()) > 0 && len(pipeName) > 0) || (len(args.Pipe()) == 0)
}
