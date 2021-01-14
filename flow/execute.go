package flow

import (
	"errors"
	"fmt"
	"strings"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow/evaluate"
	"github.com/debeando/lightflow/flow/execute"
	"github.com/debeando/lightflow/flow/slack"
	"github.com/debeando/lightflow/flow/template"
)

// Execute is a method to render and execute command and save result into
// variables and then are interpreted to take other decisions.
func (f *Flow) Execute() {
	cmd := f.renderCommand()

	if args.DryRun() {
		fmt.Println(cmd)
	} else {
		f.Retry(func() {
			f.execute(cmd)
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
	stdout, exitCode := execute.Execute(cmd)

	f.Variables.Set(map[string]interface{}{
		"exit_code": exitCode,
		"stdout":    stdout,
	})
}

func (f *Flow) parse() {
	if err := f.parseStdout(); err != nil {
		log.Error(err.Error(), nil)
	}
}

func (f *Flow) parseStdout() error {
	switch f.GetFormat() {
	case config.TEXT:
		if reg := f.GetProperty("Register"); len(reg) > 0 {
			f.Variables.Set(map[string]interface{}{reg: f.GetVariable("stdout")})
		}
	case config.JSON:
		//f.Variables puede tener un metodo para salvar en json de forma automatica?
		raw, err := common.StringToJSON(common.InterfaceToString(f.GetVariable("stdout")))
		if err != nil {
			return err
		}

		for variable, value := range raw {
			f.Variables.Set(map[string]interface{}{variable: value})
		}
	default:
		return errors.New("Format option is invalid, please use; TEXT (default) or JSON")
	}

	return nil
}

// Skip evaluate condition to set skip flag.
func (f *Flow) skip() {
	expression, err := template.Render(f.GetProperty("Skip"), f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	f.Skip = evaluate.Expression(expression)
}

// Error evaluate expression to identify any error or suggest error.
func (f *Flow) error() {
	error := f.GetProperty("Error")
	if len(error) == 0 {
		error = "{{ .exit_code }} != 0"
	}

	expression := f.Render(error)
	vars := template.Variables(error)

	debug_vars := make(map[string]interface{})
	for _,v := range vars {
    	debug_vars[v] = f.GetVariable(v)
	}

    debug_vars["Expression"] = error
    debug_vars["Rendered"]   = expression

	if evaluate.Expression(expression) {
		log.Error(f.GetTitle(), debug_vars)
	}
}

// Print specific variable with value.
func (f *Flow) print() {
	names := f.GetProperty("Print")
	if len(names) > 0 {
		vars := make(map[string]interface{})

		for _, name := range strings.Split(names, ",") {
			key := strings.Trim(name, " ")
			vars[key] = f.Variables.Get(key)
		}

		log.Info(
			fmt.Sprintf(
				"TASK[%s] SUB TASK[%s] PIPE[%s] PRINT:",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			vars,
		)
	}
}

// Debug print all variables in debug mode.
func (f *Flow) debug() {
	//f.Variables deberia tener un debug.
	for variable, value := range f.Variables.Items {
		log.Debug(f.GetTitle(), map[string]interface{}{
			variable: value,
		})
	}
}

func (f *Flow) slack() {
	expression := f.Render(f.GetSlackExpression())

	if evaluate.Expression(expression) {
		title := f.Render(f.GetSlackTitle())
		message := f.Render(f.GetSlackMessage())

		slack.Token = f.Config.General.Slack.Token
		slack.Send(
			f.GetSlackChannel(),
			title,
			message,
			f.GetSlackColor(),
		)

		log.Info(
			fmt.Sprintf(
				"TASK[%s] SUB TASK[%s] PIPE[%s] SM2S!",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			nil,
		)
	}
}

func (f *Flow) Render(s string) string {
	r, err := template.Render(s, f.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
