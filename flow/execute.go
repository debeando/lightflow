package flow

import (
	"fmt"
	"errors"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/execute"
	"github.com/swapbyt3s/lightflow/flow/template"
)

func (f *Flow) ExecuteCommand() bool {
	cmd := f.RenderCommand()

	if common.GetArgVal("dry-run").(bool) {
		fmt.Println(cmd)

		return false
	} else {
		f.Execute(cmd)

		if err := f.ParseStdout(); err != nil {
			log.Error(err.Error(), nil)
			return false
		}

		return f.RetryCommand()
	}

	return false
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

func (f *Flow) Execute(cmd string) {
	stdout, exit_code := execute.Execute(cmd)

	f.Variables.Set(map[string]interface{}{
		"exit_code": exit_code,
		"stdout": stdout,
	})
}

func (f *Flow) ParseStdout() error {
	switch f.GetFormat() {
	case "TEXT":
		if reg := f.GetRegister(); len(reg) > 0 {
			f.Variables.Set(map[string]interface{}{reg: f.GetStdOut()})
		}
	case "JSON":
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

func (f *Flow) RetryCommand() bool {
	// Log possible error and retry it is true the error:
	if error := f.Variables.Get(f.GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
		log.Error(common.InterfaceToString(error), nil)
		return false
	}

	// Si el status que retorna el stdout es diferente reintenta
	if status := f.Variables.Get(f.GetRetryStatus()); common.InterfaceToString(status) == f.GetRetryDone() {
		return false
	}

	return true
}
