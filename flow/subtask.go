package flow

import (
	"fmt"
	"os"
	"reflect"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"

	"io/ioutil"
	"github.com/go-yaml/yaml"
)

func (f *Flow) Subtask() {
	for _, inc := range f.Config.Tasks[f.Index.Task].SubtasksInclude {
		source, err := ioutil.ReadFile(inc)
		if err != nil {
			log.Error(err.Error(), nil)
		} else {
			log.Info(
				fmt.Sprintf(
					"Include Subtask: %s",
					inc,
				), nil)
		}

		var subtask config.Subtask

		source = []byte(os.ExpandEnv(string(source)))

		if err := yaml.Unmarshal(source, &subtask); err != nil {
			log.Error(
				fmt.Sprintf(
					"Imposible to parse config file: %s",
					inc,
				), nil)
			os.Exit(1)
		}

		if ! reflect.DeepEqual(subtask, config.Subtask{}) {
			f.Config.Tasks[f.Index.Task].Subtasks = append(f.Config.Tasks[f.Index.Task].Subtasks, subtask)
		}
	}





	for _, inc := range f.Config.Tasks[f.Index.Task].PipesInclude {
		source, err := ioutil.ReadFile(inc)
		if err != nil {
			log.Error(err.Error(), nil)
		} else {
			log.Info(
				fmt.Sprintf(
					"Include Pipe: %s",
					inc,
				), nil)
		}

		var pipe config.Pipe

		source = []byte(os.ExpandEnv(string(source)))

		if err := yaml.Unmarshal(source, &pipe); err != nil {
			log.Error(
				fmt.Sprintf(
					"Imposible to parse config file: %s",
					inc,
				), nil)
			os.Exit(1)
		}

		if ! reflect.DeepEqual(pipe, config.Pipe{}) {
			f.Config.Tasks[f.Index.Task].Pipes = append(f.Config.Tasks[f.Index.Task].Pipes, pipe)
		}
	}










	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Subtasks,
		Name:  args.Subtask(),
	}

	itr.Run(func() bool {
		f.Index.Pipe = 0
		f.Index.Subtask = itr.Index
		f.Skip = false

		if err := f.Date(); err != nil {
			log.Error(err.Error(), nil)
		}
		return false
	})

	log.Info(
		fmt.Sprintf(
			"%s Finished %s", // ET is acronym for execution time.
			f.TaskName(),
			itr.ExecutionTime,
		), nil)
}
