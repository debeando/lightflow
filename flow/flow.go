package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/variables"
)

type Flow struct {
	Config    config.Structure
	Index     Index
	Skip      bool
	Variables variables.List
	Attempt   int
}

type Index struct {
	Pipe    int
	Subtask int
	Task    int
}

func (f *Flow) Run() {
	f.Config = *config.Load()
	f.Variables = *variables.Load()

	PrintArgumentVariables()

	if args.List() {
		f.List()
	} else {
		f.Tasks()
	}
}

func (f *Flow) List() {
	for task_index := range f.Config.Tasks {
		fmt.Println("* Task:", f.Config.Tasks[task_index].Name)
		for subtask_index := range f.Config.Tasks[task_index].Subtask {
			fmt.Println("  > Subtask:", f.Config.Tasks[task_index].Subtask[subtask_index].Name)
			for pipe_index := range f.Config.Tasks[task_index].Pipes {
				fmt.Println("    - Pipe:", f.Config.Tasks[task_index].Pipes[pipe_index].Name)
			}
		}
	}
}

// PrintArgumentVariables print the variables passed via arguments on the terminal.
func PrintArgumentVariables() {
	if len(args.Variables()) > 0 {
		log.Info("Argument variables", args.Variables())
	}
}
