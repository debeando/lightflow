package registry

import (
	"github.com/swapbyt3s/lightflow/config"
)

type Registry struct{
	Config *config.Structure
	Task int
	Looping int
	Pipe int
}

var registry *Registry

func Load() *Registry {
	if registry == nil {
		registry = &Registry{
			Config: config.Load(),
		}
	}
	return registry
}
