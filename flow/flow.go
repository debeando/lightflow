package flow

import (
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/variables"
)

type Flow struct {
	Index Index
	Config config.Structure
	Variables variables.List
}

type Index struct {
	Task int
	Loop int
	Pipe int
}

func (f *Flow) Run() {
	f.Config = *config.Load()
	f.Variables = *variables.Load()
	f.Tasks()
}
