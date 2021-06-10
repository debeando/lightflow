package flow

import (
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/variables"
)

type Flow struct {
	Config    config.Structure
	Index     Index
	Skip      bool
	Variables variables.List
	Attempt   int
	Interval  bool
}

type Index struct {
	Pipe    int
	Task    int
}

func (f *Flow) Run() {
	f.Config = *config.Load()
	f.Variables = *variables.Load()

	f.Tasks()
}

