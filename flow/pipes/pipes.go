package pipes

import (
	"fmt"
	"errors"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

type Pipe struct {
	Task int
	Loop int
	Index int
	Title string
	ExecutionTime string
	Config config.Structure
}

func (p *Pipe) Valid() (err error) {
	if common.IsArgStringDefined("pipe") && ! p.Exist(common.GetArgVal("pipe").(string)) {
		return errors.New("Pipe name passed by argument does not exist in manifest, please verify.")
	}

	return err
}

func (p *Pipe) Exist(name string) bool {
	if len(p.Config.Tasks[p.Task].Pipes) >= 1 {
		for index := range p.Config.Tasks[p.Task].Pipes {
			if loop_name := p.Config.Tasks[p.Task].Pipes[index].Name; len(loop_name) > 0 && loop_name == name {
				p.Index = index
				return true
			}
		}
	}
	return false
}

func (p *Pipe) Run(fn func()) error {
	p.ExecutionTime = common.Duration(func(){
		if common.IsArgStringDefined("pipe") && p.Exist(common.GetArgVal("pipe").(string)) {
			p.One(fn)
		} else {
			p.All(fn)
		}
	})

	return nil
}

func (p *Pipe) All(fn func()) {
	for index := range p.Config.Tasks[p.Task].Pipes {
		p.Index = index
		p.One(fn)
	}
}

func (p *Pipe) One(fn func()) {
	p.Title = fmt.Sprintf(
		"Task[%s] Loop[%s] Pipe[%s]",
		p.Config.Tasks[p.Task].Name,
		p.Config.Tasks[p.Task].Loops[p.Loop].Name,
		p.Config.Tasks[p.Task].Pipes[p.Index].Name,
	)

	fn()
}
