package core

import (
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/variables"
)

type Core struct {
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

func (core *Core) Run() {
	core.Config = *config.Load()
	core.Variables = *variables.Load()

	core.Tasks()
}

