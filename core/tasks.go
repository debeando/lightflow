package core

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/core/iterator"
)

func (core *Core) Tasks() {
	itr := iterator.Iterator{
		Items: core.Config.Tasks,
		Name:  args.Task(),
	}

	log.Info(
		fmt.Sprintf(
			"Starting...",
		), nil)

	itr.Run(func() bool {
		core.Index.Pipe = 0
		core.Index.Task = itr.Index
		core.Skip = false

		if err := core.Date(); err != nil {
			log.Error(err.Error(), nil)
		}
		return false
	})

	log.Info(
		fmt.Sprintf(
			"Finished %s", // ET is acronym for execution time.
			itr.ExecutionTime,
		), nil)
}
