package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/config/example"
	"github.com/debeando/lightflow/flow"
)

// USAGE is a const to have help description for CLI.
const USAGE = `lightflow %s.
Usage:

	lightflow [--help | --version | --example ]
	lightflow --task=foo
	lightflow --task=foo --pipes=bar
	lightflow --variables='{"date": "2019-08-01"}'
	lightflow --task=foo --pipes=bar --variables='{"query": "SELECT * FROM foo", "date": "2019-08-01"}'
	lightflow --task=foo --pipes=bar --ai-date='{"start": "2019-08-01", "end":"2019-08-31"}'

Options:

  --ai-date    Auto Increment date, not compatible with --variables.
  --config     Using specific config file.
  --debug      Enable debug mode.
  --dry-run    No execute commands.
  --example    Print out full sample configuration to stdout.
  --help       Show this help.
  --list       List tasks, subtask and pipes.
  --pipes      Filter by pipe name.
  --subtask    Filter by subtask name.
  --task       Filter by task name.
  --variables  Passing variables on tasks, not compatible with --ai-date.
  --version    Print version numbers.

Default variables:

	This variables take values from Operating System.

	- date
	- year
	- month
	- day

	You can rewrite default variables by passing in JSON on --variables with same name.

For more help, plese visit: https://github.com/debeando/ligthflow/wiki
`

func Run() {
	fAIDate := flag.String("ai-date", "", "")
	fConfig := flag.String("config", "", "")
	fDryRun := flag.Bool("dry-run", false, "")
	fExample := flag.Bool("example", false, "")
	fHelp := flag.Bool("help", false, "")
	fList := flag.Bool("list", false, "")
	fVariables := flag.String("variables", "", "")
	fVersion := flag.Bool("version", false, "")
	_ = flag.Bool("debug", false, "")
	_ = flag.String("pipes", "", "")
	_ = flag.String("subtask", "", "")
	_ = flag.String("task", "", "")

	flag.Usage = func() { help(1) }
	flag.Parse()

	if err := config.Load().Read(*fConfig); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	switch {
	case *fVersion:
		fmt.Println(Version())
		os.Exit(0)
	case *fHelp:
		help(0)
	case *fExample:
		fmt.Printf(example.GetConfigFile())
	case len(*fAIDate) > 0 && len(*fVariables) > 0:
		help(0)
	case *fList:
		for task_index := range config.Load().Tasks {
			fmt.Println("* Task:", config.Load().Tasks[task_index].Name)
			for subtask_index := range config.Load().Tasks[task_index].Subtask {
				fmt.Println("  > Subtask:", config.Load().Tasks[task_index].Subtask[subtask_index].Name)
				for pipe_index := range config.Load().Tasks[task_index].Pipes {
					fmt.Println("    - Pipe:", config.Load().Tasks[task_index].Pipes[pipe_index].Name)
				}
			}
		}

		os.Exit(0)
	case *fDryRun == true:
		log.Warning("Running in safe mode, not execute any commands, only print commands.", nil)
	}

	flow := flow.Flow{}
	flow.Run()
}

func help(rc int) {
	fmt.Printf(USAGE, Version())
	os.Exit(rc)
}
