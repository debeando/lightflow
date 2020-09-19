package registry

import (
	"time"

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

func (r *Registry) GetRetryAttempts() int {
	return r.Config.Tasks[r.Task].Pipes[r.Pipe].Retry.Attempts
}

func (r *Registry) GetRetryWait() time.Duration {
	return time.Duration(r.Config.Tasks[r.Task].Pipes[r.Pipe].Retry.Wait)
}

func (r *Registry) GetRetryError() string {
	value := r.Config.Tasks[r.Task].Pipes[r.Pipe].Retry.Error

	if len(value) == 0 {
		value = "error"
	}

	return value
}

func (r *Registry) GetRetryWhile() string {
	return r.Config.Tasks[r.Task].Pipes[r.Pipe].Retry.While
}
