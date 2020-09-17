package pipes

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/registry"
)

type Pipe struct {
	Index int
}

func (p *Pipe) Run(fn func()) {
	if common.IsArgDefined("pipe") {
		p.Position()
		p.One(fn)
	} else {
		if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) > 0 {
			p.All(fn)
		} else {
			p.One(fn)
		}
	}
}

func (p *Pipe) Set(index int) {
	p.Index = index
	registry.Load().Pipe = index
}

func (p *Pipe) Position() {
	arg := common.GetArgVal("pipe")

	for pipe := range registry.Load().Config.Tasks[registry.Load().Task].Pipes {
		if name := registry.Load().Config.Tasks[registry.Load().Task].Pipes[pipe].Name; len(name) > 0 {
			if name == arg {
				p.Set(pipe)
				return
			}
		}
	}

	p.Set(0)
}

func (p *Pipe) All(fn func()) {
	for index := range registry.Load().Config.Tasks[registry.Load().Task].Pipes {
		p.Set(index)
		p.One(fn)
	}
}

func (p *Pipe) One(fn func()) {
 	fn()
}
