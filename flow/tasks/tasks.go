package tasks

import (
	"errors"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

type Task struct {
	Index int
	Title string
	ExecutionTime string
	Config config.Structure
}

func (t *Task) Valid() (err error) {
	if len(t.Config.Tasks) == 0 {
		return errors.New("Task in the manifest is empty or malformed, please verify.")
	}

	if common.IsArgStringDefined("task") && ! t.Exist(common.GetArgVal("task").(string)) {
		return errors.New("Task name passed by argument does not exist in manifest, please verify.")
	}

	return err
}

func (t *Task) Exist(name string) bool {
	if len(t.Config.Tasks) >= 1 {
		for index := range t.Config.Tasks {
			if task_name := t.Config.Tasks[index].Name; len(task_name) > 0 && task_name == name {
				t.Index = index
				return true
			}
		}
	}
	return false
}

func (t *Task) Run(fn func()) error {
	t.ExecutionTime = common.Duration(func(){
		if common.IsArgStringDefined("task") && t.Exist(common.GetArgVal("task").(string)) {
			t.One(fn)
		} else {
			t.All(fn)
		}
	})

	return nil
}

func (t *Task) All(fn func()) {
	for index := range t.Config.Tasks {
		t.Index = index
		t.One(fn)
	}
}

func (t *Task) One(fn func()) {
	fn()
}
