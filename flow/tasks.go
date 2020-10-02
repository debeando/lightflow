package flow

import (
	"os"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/tasks"
)

func (f *Flow) Task() {
	tsk := tasks.Task{
		Config: f.Config,
	}

	if err := tsk.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	f.Index.Task = 0
	f.Index.Loop = 0
	f.Index.Pipe = 0

	err := tsk.Run(func() {
		f.Index.Task = tsk.Index
		f.Loop()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"TASKS",
			map[string]interface{}{
				"Execution Time": tsk.ExecutionTime,
			})
	}
}
