package loops

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/registry"
)

type Looping struct {
	Index int
}

func (l *Looping) Run(fn func()) {
	if common.IsArgDefined("looping") {
		l.Position()
		l.One(fn)
	} else {
		if len(registry.Load().Config.Tasks[registry.Load().Task].Loops) > 0 {
			l.All(fn)
		} else {
			l.One(fn)
		}
	}
}

func (l *Looping) Set(index int) {
	l.Index = index
	registry.Load().Looping = index
}

func (l *Looping) Position() {
	arg := common.GetArgVal("looping")

	for looping := range registry.Load().Config.Tasks[registry.Load().Task].Loops {
		if name := registry.Load().Config.Tasks[registry.Load().Task].Loops[looping]["name"]; len(name) > 0 {
			if name == arg {
				l.Set(looping)
				return
			}
		}
	}

	l.Set(0)
}

func (l *Looping) All(fn func()) {
	for index := range registry.Load().Config.Tasks[registry.Load().Task].Loops {
		l.Set(index)
		l.One(fn)
	}
}

func (l *Looping) One(fn func()) {
 	fn()
}
