package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/debeando/lightflow/common"
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
	lightflow --task=foo --pipes=bar,baz
	lightflow --variables='{"date": "2019-08-01"}'
	lightflow --task=foo --pipes=bar --variables='{"query": "SELECT * FROM foo", "date": "2019-08-01"}'
	lightflow --task=foo --pipes=bar --ai-date='{"start": "2019-08-01", "end":"2019-08-31"}'

Options:

  --ai-date    Auto Increment date, not compatible with --variables.
  --ad-date    Auto Decrement date, not compatible with --variables.
  --config     Using specific config file.
  --debug      Enable debug mode.
  --dry-run    No execute commands.
  --example    Print out full sample configuration to stdout.
  --help       Show this help.
  --list       List tasks, subtask and pipes.
  --pipes      Filter by one or many pipe name.
  --subtask    Filter by subtask name.
  --task       Filter by task name.
  --variables  Passing variables in JSON format, is not compatible with --ai-date.
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
	_ = flag.Bool("debug", false, "")
	_ = flag.Bool("list", false, "")
	_ = flag.String("pipes", "", "")
	_ = flag.String("subtask", "", "")
	_ = flag.String("task", "", "")

	fAIDate := flag.String("ai-date", "", "")
	fADDate := flag.String("ad-date", "", "")
	fConfig := flag.String("config", "", "")
	fDryRun := flag.Bool("dry-run", false, "")
	fExample := flag.Bool("example", false, "")
	fHelp := flag.Bool("help", false, "")
	fVariables := flag.String("variables", "", "")
	fVersion := flag.Bool("version", false, "")

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
		os.Exit(0)
	case len(*fVariables) > 0:
		if ok, err := isValidJSON("variables"); ok == false {
			log.Error("Problem to parse JSON in argument --variables", nil)
			log.Error(fmt.Sprintf("%s: %s", err, common.GetArgVal("variables")), nil)
			os.Exit(1)
		}
	case len(*fAIDate) > 0 && len(*fVariables) > 0:
		help(0)
	case len(*fADDate) > 0 && len(*fVariables) > 0:
		help(0)
	case len(*fAIDate) > 0 && len(*fADDate) > 0:
		help(0)
	case len(*fAIDate) > 0 && len(*fADDate) > 0:
		os.Exit(1)
	case len(*fAIDate) > 0:
		if ok, err := isValidJSON("ai-date"); ok == false {
			log.Error("Problem to parse JSON in argument --ai-date", nil)
			log.Error(fmt.Sprintf("%s: %s", err, common.GetArgVal("ai-date")), nil)
			os.Exit(1)
		}
	case len(*fADDate) > 0:
		if ok, err := isValidJSON("ad-date"); ok == false {
			log.Error("Problem to parse JSON in argument --ad-date", nil)
			log.Error(fmt.Sprintf("%s: %s", err, common.GetArgVal("ad-date")), nil)
			os.Exit(1)
		}
	case *fDryRun == true:
		log.Warning("Running in safe mode, no execute commands.", nil)
	}

	if err := config.Load().Read(*fConfig); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	flow := flow.Flow{}
	flow.Run()
}

func help(rc int) {
	fmt.Printf(USAGE, Version())
	os.Exit(rc)
}

func isValidJSON(name string) (bool, error) {
	args_vars := common.GetArgVal(name)

	switch v := args_vars.(type) {
	case string:
		c := strings.Trim(v, "'")

		_, err := common.StringToJSON(c)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
