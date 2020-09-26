package flow

import (
	"github.com/swapbyt3s/lightflow/config"
)

type Flow struct {
	Index Index
	Config config.Structure
}

type Index struct {
	Task int
	Loop int
	Pipe int
}

func (f *Flow) Run() {
	f.Config = *config.Load()
	f.Task()
}
