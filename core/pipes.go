package core

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/core/iterator"
)

func (core *Core) PipesInArgs() {
	for _, pipeName := range args.Pipes() {
		if core.Valid(pipeName) {
			itr := iterator.Iterator{
				Items: core.Config.Pipes,
				Name:  pipeName,
			}

			itr.Run(func() bool {
				core.Index.Pipe = itr.Index

				if ! core.Skip {
					log.Info(core.GetTitle(), nil)
					core.Chunks()
				}

				return core.Skip
			})

			if ! core.Skip {
				log.Info(
					fmt.Sprintf(
						"%s Finished %s", // ET is acronym for execution time.
						core.TaskName(),
						itr.ExecutionTime,
					), nil)
			}
		}
	}
}

func (core *Core) Valid(pipeName string) bool {
	return (len(args.Pipes()) > 0 && len(pipeName) > 0) || (len(args.Pipe()) == 0)
}
