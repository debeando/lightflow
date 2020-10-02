package flow

import (
	"fmt"
	"errors"

	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/execute"
	"github.com/swapbyt3s/lightflow/flow/template"
)

func (f *Flow) Execute() {
	cmd := f.RenderCommand()

	if common.GetArgVal("dry-run").(bool) {
		fmt.Println(cmd)
	} else {
		f.Retry(func () bool {
			stdout, exit_code := execute.Execute(cmd)

			f.Variables.Set(map[string]interface{}{
				"exit_code": exit_code,
				"stdout": stdout,
			})

			if exit_code > 0 && f.GetRetryExitCode() == 0 {
				log.Error(f.GetTitle(), map[string]interface{}{
					"Exit Code": exit_code,
				})
			}

			if err := f.ParseStdout(); err != nil {
				log.Error(err.Error(), nil)
			}

			for variable, value := range f.Variables.Items {
				log.Debug(f.GetTitle(), map[string]interface{}{
					variable: value,
				})
			}

			return f.EvalRetry()
		})
	}
}

func (f *Flow) RenderCommand() string {
	var cmd = f.GetExecute()

	for _, variable := range template.Variables(cmd) {
		if f.Variables.Exist(variable) == false {
			log.Warning("Variable not defined", map[string]interface{}{
				"Variable Name": variable,
			})
		}
	}

	cmd, err := template.Render(cmd, f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(cmd)
}

func (f *Flow) ParseStdout() error {
	switch f.GetFormat() {
	case config.TEXT:
		if reg := f.GetRegister(); len(reg) > 0 {
			f.Variables.Set(map[string]interface{}{reg: f.GetStdOut()})
		}
	case config.JSON:
		raw, err := common.StringToJSON(common.InterfaceToString(f.GetStdOut()))
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

func (f *Flow) EvalRetry() bool {
	if f.GetRetryWait() == 0 {
		return false
	}

	if f.GetRetryAttempts() == 0 {
		return false
	}

	exit_code := f.Variables.Get("exit_code").(int)
	status := common.InterfaceToString(f.Variables.Get(f.GetRetryStatus()))
	error := common.InterfaceToString(f.Variables.Get(f.GetRetryError()))

	if f.GetRetryExitCode() != exit_code {
		log.Warning(f.GetTitle() + " Retry", map[string]interface{}{
			"Exit Code": exit_code,
			"Expected": f.GetRetryExitCode(),
		})

		return true
	}

	if exit_code == 0 && len(error) == 0 && len(status) > 0 && len(f.GetRetryDone()) > 0 && f.GetRetryDone() != status {
		log.Info(f.GetTitle() + " Retry", map[string]interface{}{
			"Status": status,
		})

		return true
	}

	if len(error) > 0 {
		log.Error(common.InterfaceToString(error), nil)
	}

	return false
}
