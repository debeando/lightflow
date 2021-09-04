package core

import (
	"fmt"
	"time"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/execute"
	"github.com/debeando/lightflow/plugins/template"
)

// Execute is a method to render and execute command and save result into
// variables and then are interpreted to take other decisions.
func (core *Core) Execute() {
	cmd := core.renderCommand()

	if core.when() {
		core.Retry(func() {
			core.unset()
			core.wait()
			core.execute(cmd)
			core.mysql()
			core.aws()
			core.parse()
			core.error()
			core.print()
			core.debug()
			core.skip()
			core.slack()
		})
	}
}

func (core *Core) renderCommand() string {
	var cmd = core.GetProperty("Execute")
	var vars = core.Variables.GetItems()

	// Find unknown variables:
	for _, variable := range template.Variables(cmd) {
		if core.Variables.Exist(variable) == false {
			log.Warning(fmt.Sprintf("Register empty variable: %s", variable), nil)
			core.Variables.Set(map[string]interface{}{variable: ""})
		}
	}

	// Find template variables to render:
	for variable, value := range core.Variables.Items {
		value_template := template.Variables(common.InterfaceToString(value))

		if len(value_template) > 0 {
			cmd, err := template.Render(common.InterfaceToString(value), vars)
			if err != nil {
				log.Warning(err.Error(), nil)
			}

			vars[variable] = cmd
		}
	}

	// Render template:
	cmd, err := template.Render(cmd, vars)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(cmd)
}

func (core *Core) execute(cmd string) {
	if len(cmd) == 0 {
		return
	}

	log.Debug(
		fmt.Sprintf(
			"%s/%s",
			core.TaskName(),
			core.PipeName(),
		),
		map[string]interface{}{
			"Execute": cmd,
		},
	)

	stdout, exitCode := execute.Execute(cmd, args.DryRun())

	core.Variables.Set(map[string]interface{}{
		"exit_code": exitCode,
		"stdout":    stdout,
	})
}

func (core *Core) wait() {
	wait := core.Config.Pipes[core.Index.Pipe].Wait
	time.Sleep(time.Duration(wait) * time.Second)
}
