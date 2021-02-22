package flow

import (
	"fmt"

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

	if f.when() {
		f.Retry(func() {
			f.unset()
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

func (f *Flow) unset() {
	for _, key := range f.GetPipeUnset() {
		f.Variables.Set(map[string]interface{}{
			key: "",
		})
	}
}

func (f *Flow) execute(cmd string) {
	log.Debug(f.GetTitle(), map[string]interface{}{
		"Execute": cmd,
	})

	stdout, exitCode := execute.Execute(cmd, args.DryRun())

	f.Variables.Set(map[string]interface{}{
		"exit_code": exitCode,
		"stdout":    stdout,
	})
}

// parseStdout verify the formant and store value(s) in variable register.
func (f *Flow) parse() {
	switch f.GetFormat() {
	case config.TEXT:
		if reg := f.GetProperty("Register"); len(reg) > 0 {
			f.Variables.Set(map[string]interface{}{reg: f.GetVariable("stdout")})
		}
	case config.JSON:
		//f.Variables puede tener un metodo para salvar en json de forma automatica?
		raw, err := common.StringToJSON(common.InterfaceToString(f.GetVariable("stdout")))
		if err != nil {
			log.Error(err.Error(), nil)
		}

		for variable, value := range raw {
			f.Variables.Set(map[string]interface{}{variable: value})
		}
	default:
		log.Error("Format option is invalid, please use; TEXT (default) or JSON", nil)
	}
}

// When a condition is true allow execute pipe.
func (f *Flow) when() bool {
	if len(f.GetProperty("When")) == 0 {
		return true
	}

	expression, err := template.Render(f.GetProperty("When"), f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	value := evaluate.Expression(expression)

	debug_vars := make(map[string]interface{})
	debug_vars["Expression"] = f.GetProperty("When")
	debug_vars["Rendered"] = expression
	debug_vars["Result"] = value

	log.Debug(f.GetTitle(), debug_vars)

	if !value {
		log.Info(
			fmt.Sprintf(
				"TASK[%s] SUB TASK[%s] PIPE[%s] !WHEN",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			nil,
		)
	}

	return value
}

// Skip evaluate condition to set skip flag.
func (f *Flow) skip() {
	expression, err := template.Render(f.GetProperty("Skip"), f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	f.Skip = evaluate.Expression(expression)

	f.Variables.Set(map[string]interface{}{
		"skip": f.Skip,
	})
}

// Error evaluate expression to identify any error or suggest error.
func (f *Flow) error() {
	error := f.GetProperty("Error")
	if len(error) == 0 {
		error = "{{ .exit_code }} != 0 || {{ len .error }} > 0"
	}

	expression := f.Render(error)
	vars := template.Variables(error)

	debug_vars := make(map[string]interface{})
	for _, v := range vars {
		debug_vars[v] = f.GetVariable(v)
	}

	debug_vars["Expression"] = error
	debug_vars["Rendered"] = expression
	debug_vars["Stdout"] = f.GetVariable("stdout")

	if evaluate.Expression(expression) {
		log.Error(f.GetTitle(), debug_vars)
	}
}

// Print specific variable with value.
func (f *Flow) print() {
	names := f.GetPipePrint()
	if len(names) > 0 {
		vars := make(map[string]interface{})

		for _, key := range names {
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
	for variable, value := range f.Variables.Items {
		log.Debug(f.GetTitle(), map[string]interface{}{
			variable: value,
		})
	}
}

// Slack send custom message.
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
				"TASK[%s] SUB TASK[%s] PIPE[%s] Send message to slack.",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			nil,
		)
	}
}

// Render a template with variables.
func (f *Flow) Render(s string) string {
	r, err := template.Render(s, f.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
