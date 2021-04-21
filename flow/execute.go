package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/execute"
	"github.com/debeando/lightflow/flow/template"
)

// Execute is a method to render and execute command and save result into
// variables and then are interpreted to take other decisions.
func (f *Flow) Execute() {
	cmd := f.renderCommand()

	if f.when() {
		f.Retry(func() {
			f.unset()
			f.execute(cmd)
			f.mysql()
			f.parse()
			f.error()
			f.print()
			f.debug()
			f.skip()
			f.slack()
		})
	}
}

func (f *Flow) renderCommand() string {
	var cmd = f.GetProperty("Execute")
	var vars = f.Variables.GetItems()

	// Find unknown variables:
	for _, variable := range template.Variables(cmd) {
		if f.Variables.Exist(variable) == false {
			log.Warning(fmt.Sprintf("Register empty variable: %s", variable), nil)
			f.Variables.Set(map[string]interface{}{variable: ""})
		}
	}

	// Find template variables to render:
	for variable, value := range f.Variables.Items {
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

func (f *Flow) execute(cmd string) {
	if len(cmd) == 0 {
		return
	}

	log.Debug(
		fmt.Sprintf(
			"%s/%s/%s",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
		),
		map[string]interface{}{
			"Execute": cmd,
		},
	)

	stdout, exitCode := execute.Execute(cmd, args.DryRun())

	f.Variables.Set(map[string]interface{}{
		"exit_code": exitCode,
		"stdout":    stdout,
	})
}
