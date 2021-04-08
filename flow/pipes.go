package flow

import (
	"fmt"
	"time"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Pipes() {
	for _, pipeName := range args.Pipes() {
		if f.Valid(pipeName) {
			itr := iterator.Iterator{
				Items: f.Config.Tasks[f.Index.Task].Pipes,
				Name:  pipeName,
			}

			itr.Run(func() bool {
				f.Index.Pipe = itr.Index
				log.Info(f.GetTitle(), nil)
				f.Wait()
				f.Chunks()

				if f.Skip {
					log.Warning(
						fmt.Sprintf(
							"%s/%s/%s",
							f.TaskName(),
							f.SubTaskName(),
							f.PipeName(),
						), nil)
				}

				return f.Skip
			})

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

func (f *Flow) Valid(pipeName string) bool {
	return (len(args.Pipes()) > 0 && len(pipeName) > 0) || (len(args.Pipe()) == 0)
}

func (f *Flow) Wait() {
	wait := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Wait
	time.Sleep(time.Duration(wait) * time.Second)
}
