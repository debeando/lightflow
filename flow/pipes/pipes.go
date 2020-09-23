package pipes

import (
	// "fmt"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/registry"
)

type Pipe struct {
	Index int
	Name string
}

func (p *Pipe) Run(fn func()) {
	if ! p.Validate() {
		return
	}

	if common.IsArgDefined("pipe") && len(common.GetArgVal("pipe").(string)) > 0 {
		// fmt.Println("p")
		if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) >= 1 {
			p.Position()
			p.One(fn)
		}
	} else {
		// fmt.Println("all")
		if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) > 1 {
			p.All(fn)
		} else if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) == 1 {
			p.One(fn)
		}
	}
}

func (p *Pipe) Set(index int) {
	p.Index = index
	registry.Load().Pipe = index
}

func (p *Pipe) Position() {
	for pipe := range registry.Load().Config.Tasks[registry.Load().Task].Pipes {
		if name := registry.Load().Config.Tasks[registry.Load().Task].Pipes[pipe].Name; len(name) > 0 {
			if name == p.Name {
				p.Set(pipe)
				return
			}
		}
	}

	p.Set(-1)
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

func (p *Pipe) Validate() bool {
	if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) == 0 {
		log.Error("Pipe in the manifest is empty or malformed, please verify.", nil)
		return false
	}

	if common.IsArgDefined("pipe") && len(common.GetArgVal("pipe").(string)) > 0 {
		p.Name = common.GetArgVal("pipe").(string)
		if len(p.Name) > 0 && !p.Exist() {
			log.Error("Pipe name passed by argument does not exist in manifest.", nil)
			return false
		}

	}

	return true
}

func (p *Pipe) Exist() bool {
	if len(registry.Load().Config.Tasks[registry.Load().Task].Pipes) >= 1 {
		for pipe := range registry.Load().Config.Tasks[registry.Load().Task].Pipes {
			if registry.Load().Config.Tasks[registry.Load().Task].Pipes[pipe].Name == p.Name {
				return true
			}
		}
	}
	return false
}
