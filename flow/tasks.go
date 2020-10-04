package flow

import (
	"os"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/tasks"
)

func (f *Flow) Task() {
	task := tasks.Task{
		Config: f.Config,
	}

	if err := task.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	f.Index.Task = 0
	f.Index.Loop = 0
	f.Index.Pipe = 0

	err := task.Run(func() {
		f.Index.Task = task.Index
		f.Loop()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"TASKS",
			map[string]interface{}{
				"Execution Time": task.ExecutionTime,
			})
	}
}
