package flow

import (
	"os"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/pipes"
)

func (f *Flow) Pipes() {
	pipe := pipes.Pipe{
		Task: f.Index.Task,
		Loop: f.Index.Loop,
		Config: f.Config,
	}

	if err := pipe.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	err := pipe.Run(func() {
		log.Info(pipe.Title, nil)
		f.Index.Pipe =  pipe.Index
		f.AutoIncrement()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"Pipe",
			map[string]interface{}{
				"Execution Time": pipe.ExecutionTime,
			})
	}
}
