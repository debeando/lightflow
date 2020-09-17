package tasks

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/registry"
)

type Task struct {
	Index int
}

func (t *Task) Run(fn func()) {
	if common.IsArgDefined("task") {
		t.Position()
		t.One(fn)
	} else {
		if len(registry.Load().Config.Tasks) > 0 {
			t.All(fn)
		} else {
			t.One(fn)
		}
	}
}

func (t *Task) Set(index int) {
	t.Index = index
	registry.Load().Task = index
}

func (t *Task) Position() {
	name := common.GetArgVal("task")

	for index := range registry.Load().Config.Tasks {
		if registry.Load().Config.Tasks[index].Name == name {
			t.Set(index)
			return
		}
	}
	t.Set(0)
}

func (t *Task) All(fn func()) {
	for index := range registry.Load().Config.Tasks {
		t.Set(index)
		t.One(fn)
	}
}

func (t *Task) One(fn func()) {
	fn()
}
