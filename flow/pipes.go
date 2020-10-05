package flow

import (
	"fmt"
	"os"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/pipes"
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

	f.Index.Pipe = 0

	err := pipe.Run(func() {
		f.Index.Pipe = pipe.Index
		log.Info(f.GetTitle(), nil)
		f.Chunks()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			fmt.Sprintf(
				"TASK[%s] LOOP[%s] PIPES-ET[%s]", // ET is acronym for Execution Time.
				f.GetTaskName(),
				f.GetLoopName(),
				pipe.ExecutionTime,
			), nil)
	}
}
