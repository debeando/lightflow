package execute

import (
	"errors"
	"os/exec"
	"syscall"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/common/template"
	"github.com/debeando/lightflow/variables"
)

type Execute struct{
	// Comando a ejecutar. Si hay que limpiar el stdout en formato JSON, usar
	// tool jq.
	Command string `yaml:"command"`
	// If true, is eval mode, not execute.
	DryRun  bool   `yaml:"dryrun"`
}

func init() {
	plugin.Add("Execute", func() plugin.Plugin { return &Execute{} })
}

func (e *Execute) Run(event interface{}) (error, bool) {
	var stdout string
	var exitcode int
	var err error
	var out []byte

	execute, ok := event.(Execute)
	if !ok {
		return errors.New("Invalid struct"), false
	}

	vars := *variables.Load()
	execute.Command, err = template.Render(execute.Command, vars.GetItems())
	if err != nil {
		return err, true
	}

	if execute.DryRun {
		out, err = exec.Command("/bin/bash", "-n", "-c", execute.Command).CombinedOutput()
	} else {
		out, err = exec.Command("/bin/bash", "-c", execute.Command).CombinedOutput()
	}

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitcode = ws.ExitStatus()
		}
	}

	stdout = string(out[:])
	stdout = common.TrimNewlines(stdout)

	vars.Set(map[string]interface{}{
		"exit_code": exitcode,
		"stdout": stdout,
	})

	return nil, false
}
