package core

import (
	"github.com/debeando/lightflow/config"
)

func (core *Core) Pipes(fn func(pipe config.Pipe) error) error {
	core.Config = *config.Load()
	for _, pipe := range core.Config.Pipes {
		err := fn(pipe)
		if err != nil {
			return err
		}
	}
	return nil
}
