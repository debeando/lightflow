package flow

import (
	"os"
	"fmt"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/loops"
)

func (f *Flow) Loop() {
	loop := loops.Loop{
		Task: f.Index.Task,
		Config: f.Config,
	}

	if err := loop.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	f.Index.Loop = 0
	f.Index.Pipe = 0

	err := loop.Run(func() {
		f.Index.Loop = loop.Index
		f.AutoIncrement()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			fmt.Sprintf(
				"TASK[%s] LOOPS",
				f.GetTaskName(),
			),
			map[string]interface{}{
				"Execution Time": loop.ExecutionTime,
			})
	}
}
