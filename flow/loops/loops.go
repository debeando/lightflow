package loops

import (
	"errors"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

type Loop struct {
	Task int
	Index int
	Title string
	ExecutionTime string
	Config config.Structure
}

func (l *Loop) Valid() (err error) {
	if len(l.Config.Tasks[l.Task].Loops) == 0 {
		return errors.New("Loop in the manifest is empty or malformed, please verify.")
	}

	if common.IsArgStringDefined("loop") && ! l.Exist(common.GetArgVal("loop").(string)) {
		return errors.New("Loop name passed by argument does not exist in manifest, please verify.")
	}

	return err
}

func (l *Loop) Exist(name string) bool {
	if len(l.Config.Tasks[l.Task].Loops) >= 1 {
		for index := range l.Config.Tasks[l.Task].Loops {
			if loop_name := l.Config.Tasks[l.Task].Loops[index].Name; len(loop_name) > 0 && loop_name == name {
				l.Index = index
				return true
			}
		}
	}
	return false
}

func (l *Loop) Run(fn func()) error {
	l.ExecutionTime = common.Duration(func(){
		if common.IsArgStringDefined("loop") && l.Exist(common.GetArgVal("loop").(string)) {
			l.One(fn)
		} else {
			l.All(fn)
		}
	})

	return nil
}

func (l *Loop) All(fn func()) {
	for index := range l.Config.Tasks[l.Task].Loops {
		l.Index = index
		l.One(fn)
	}
}

func (l *Loop) One(fn func()) {
	fn()
}
