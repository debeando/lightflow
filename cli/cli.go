package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/core"
)

// USAGE is a const to have help description for CLI.
const USAGE = `lightflow %s.
Usage:

	lightflow [--help | --version ]
	lightflow --tasks=foo
	lightflow --tasks=foo --pipes=bar,baz
	lightflow --variables='{"date": "2019-08-01"}'
	lightflow --tasks=foo --pipes=bar --variables='{"query": "SELECT * FROM foo", "date": "2019-08-01"}'

Options:

  --config     Using specific config file.
  --debug      Enable debug mode.
  --dry-run    No execute commands.
  --help       Show this help.
  --pipes      Filter by one or many pipe name.
  --tasks      Filter by task name.
  --variables  Passing variables in JSON format.
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
	_ = flag.String("pipes", "", "")
	_ = flag.String("tasks", "", "")

	fConfig := flag.String("config", "", "")
	fDryRun := flag.Bool("dry-run", false, "")
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
	case len(*fVariables) > 0:
		if ok, err := isValidJSON("variables"); ok == false {
			log.Error("Problem to parse JSON in argument --variables", nil)
			log.Error(fmt.Sprintf("%s: %s", err, common.GetArgVal("variables")), nil)
			os.Exit(1)
		}
	case *fDryRun == true:
		log.Warning("Running in safe mode, no execute commands.", nil)
	}

	if err := config.Load().Read(*fConfig); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	core := core.Core{}
	core.Run()
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
