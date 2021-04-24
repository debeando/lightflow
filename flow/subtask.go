package flow

import (
	"fmt"
	"os"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"

	"io/ioutil"
	"github.com/go-yaml/yaml"
)

func (f *Flow) Subtask() {
	for _, inc := range f.Config.Tasks[f.Index.Task].SubtasksInclude {
		log.Info(
			fmt.Sprintf(
				"Include Subtask: %s",
				inc,
			), nil)

		source, err := ioutil.ReadFile(inc)
		if err != nil {

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

		f.Config.Tasks[f.Index.Task].Subtasks = append(f.Config.Tasks[f.Index.Task].Subtasks, subtask)
	}


	itr := iterator.Iterator{
		Items: f.Config.Tasks[f.Index.Task].Subtasks,
		Name:  args.Subtask(),
	}

	itr.Run(func() bool {
		f.Index.Pipe = 0
		f.Index.Subtask = itr.Index
		f.Skip = false
		f.When = true
		f.In = false

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
