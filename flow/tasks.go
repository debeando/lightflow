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

func (f *Flow) Tasks() {
	for _, inc := range f.Config.TasksInclude {
		source, err := ioutil.ReadFile(inc)
		if err != nil {
			log.Error(err.Error(), nil)
		} else {
			log.Info(
				fmt.Sprintf(
					"Include Task: %s",
					inc,
				), nil)
		}

		var task config.Task

		source = []byte(os.ExpandEnv(string(source)))

		if err := yaml.Unmarshal(source, &task); err != nil {
			log.Error(
				fmt.Sprintf(
					"Imposible to parse config file: %s",
					inc,
				), nil)
			os.Exit(1)
		}

		if ! reflect.DeepEqual(task, config.Task{}) {
			f.Config.Tasks = append(f.Config.Tasks, task)
		}
	}

	for _, inc := range f.Config.PipesInclude {
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
			f.Config.Pipes = append(f.Config.Pipes, pipe)
		}
	}

	itr := iterator.Iterator{
		Items: f.Config.Tasks,
		Name:  args.Task(),
	}

	itr.Run(func() bool {
		f.Index.Pipe = 0
		f.Index.Task = itr.Index
		f.Skip = false

		if err := f.Date(); err != nil {
			log.Error(err.Error(), nil)
		}
		return false
	})

	log.Info(
		fmt.Sprintf(
			"Finished %s", // ET is acronym for execution time.
			itr.ExecutionTime,
		), nil)
}
