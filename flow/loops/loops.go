package loops

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/registry"
)

type Loop struct {
	Index int
}

func (l *Loop) Run(fn func()) {
	if common.IsArgDefined("loop") {
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

func (l *Loop) Set(index int) {
	l.Index = index
	registry.Load().Loop = index
}

func (l *Loop) Position() {
	arg := common.GetArgVal("loop")

	for loop := range registry.Load().Config.Tasks[registry.Load().Task].Loops {
		if name := registry.Load().Config.Tasks[registry.Load().Task].Loops[loop]["name"]; len(name) > 0 {
			if name == arg {
				l.Set(loop)
				return
			}
		}
	}

	l.Set(0)
}

func (l *Loop) All(fn func()) {
	for index := range registry.Load().Config.Tasks[registry.Load().Task].Loops {
		l.Set(index)
		l.One(fn)
	}
}

func (l *Loop) One(fn func()) {
 	fn()
}
