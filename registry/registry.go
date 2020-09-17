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

func (r *Registry) GetTaskName() string {
	if name := r.Config.Tasks[r.Task].Name; len(name) > 0 {
		return name
	}
	return ""
}

func (r *Registry) GetLoopingName() string {
	if len (r.Config.Tasks[r.Task].Loops) > 0 {
		if name := r.Config.Tasks[r.Task].Loops[r.Looping]["name"]; len(name) > 0 {
			return name
		}
	}
	return ""
}

func (r *Registry) GetPipeName() string {
	if name := r.Config.Tasks[r.Task].Pipes[r.Pipe].Name; len(name) > 0 {
		return name
	}
	return ""
}
