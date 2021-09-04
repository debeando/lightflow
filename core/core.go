package core

import (
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/variables"
)

type Core struct {
	Config config.Structure
	Variables variables.List
}

func (core *Core) Run() {
	core.Config = *config.Load()
	core.Variables = *variables.Load()

	err := core.Tasks(func() error {
		return core.Pipes(func(pipe config.Pipe) error {
			return core.Plugins(pipe)
		})
	})
	if err != nil {
		log.Error(err.Error(), nil)
	}
}
