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
