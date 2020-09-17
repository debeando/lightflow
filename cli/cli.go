package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/config/example"
	"github.com/swapbyt3s/lightflow/flow"
)

// USAGE is a const to have help description for CLI.
const USAGE = `lightflow %s.
Usage:

	lightflow [--help | --version | --example ]
	lightflow --task=foo
	lightflow --task=foo --pipe=bar
	lightflow --variables='{"date": "2019-08-01"}'
	lightflow --task=foo --pipe=bar --variables='{"query": "SELECT * FROM foo", "date": "2019-08-01"}'

Options:

  --help       Show this help.
  --example    Print out full sample configuration to stdout.
  --config     Using specific config file.
  --task       Task name to execute.
  --looping    Filter by looping.
  --pipe       Task name to execute.
  --debug      Enable debug mode.
  --variables  Passing variables on tasks.
  --version    Print version numbers.

Default variables:

	This variables take values from Operating System.

	- date
	- year
	- hour

	You can rewrite default variables by passing on --variables with same name.

For more help, plese visit: https://github.com/swapbyt3s/ligthflow/wiki
`

func Run() {
	fHelp := flag.Bool("help", false, "Show this help.")
	fExample := flag.Bool("example", false, "Print out full sample configuration to stdout.")
	fVersion := flag.Bool("version", false, "Show version.")
	fConfig := flag.String("config", "", "Using specific config file.")
	_ = flag.Bool("debug", false, "Enable debug mode.")
	_ = flag.String("task", "", "Task name to execute.")
	_ = flag.String("looping", "", "Filter by looping.")
	_ = flag.String("pipe", "", "Pipe name to execute on task.")
	_ = flag.String("variables", "", "Variables in JSON format.")

	flag.Usage = func() { help(1) }
	flag.Parse()

	switch {
	case *fVersion:
		fmt.Println(Version())
		os.Exit(0)
	case *fHelp:
		help(0)
	case *fExample:
		fmt.Printf(example.GetConfigFile())
	}

	if err := config.Load().Read(*fConfig); err != nil {
		log.Error("Config", map[string]interface{}{"error": err})
		os.Exit(1)
  	}

  	log.Configure()

	flow.Run()
}

func help(rc int) {
	fmt.Printf(USAGE, Version())
	os.Exit(rc)
}