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
	lightflow --task=foo --pipe=bar --ai-date='{"start": "2019-08-01", "end":"2019-08-31"}'

Options:

  --ai-date    Auto Increment date, not compatible with --variables.
  --config     Using specific config file.
  --debug      Enable debug mode.
  --dry-run    No execute commands.
  --example    Print out full sample configuration to stdout.
  --help       Show this help.
  --loop       Filter by loop name.
  --pipe       Filter by pipe name.
  --task       Filter by task name.
  --variables  Passing variables on tasks, not compatible with --ai-date.
  --version    Print version numbers.

Default variables:

	This variables take values from Operating System.

	- date
	- year
	- month
	- day
	- hour

	You can rewrite default variables by passing in JSON on --variables with same name.

For more help, plese visit: https://github.com/swapbyt3s/ligthflow/wiki
`

func Run() {
	fAIDate := flag.String("ai-date", "", "")
	fConfig := flag.String("config", "", "")
	fDryRun := flag.Bool("dry-run", false, "")
	fExample := flag.Bool("example", false, "")
	fHelp := flag.Bool("help", false, "")
	fVariables := flag.String("variables", "", "")
	fVersion := flag.Bool("version", false, "")
	_ = flag.Bool("debug", false, "")
	_ = flag.String("loop", "", "")
	_ = flag.String("pipe", "", "")
	_ = flag.String("task", "", "")

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
	case len(*fAIDate) > 0 && len(*fVariables) > 0:
		help(0)
	case *fDryRun == true:
		log.Warning("Safe Command in dry-run mode", nil)
	}

	if err := config.Load().Read(*fConfig); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
  	}

	flow.Run()
}

func help(rc int) {
	fmt.Printf(USAGE, Version())
	os.Exit(rc)
}
