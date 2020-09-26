package flow

import (
	"os"

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

	err := loop.Run(func() {
		f.Index.Loop = loop.Index
		f.Pipes()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"Loop",
			map[string]interface{}{
				"Execution Time": loop.ExecutionTime,
			})
	}
}
