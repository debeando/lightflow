package flow

/*
 * TODO: Meter dentro del command el retry y no al revez.
 */

import (
	"fmt"
	"errors"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/execute"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/variables"
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
	var v = variables.Load()

	var cmd = f.GetExecute()

	for _, variable := range template.Variables(cmd) {
		if v.Exist(variable) == false {
			log.Warning("Variable not defined", map[string]interface{}{
				"Variable Name": variable,
			})
		}
	}

	cmd, err := template.Render(cmd, v.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(cmd)
}

func (f *Flow) Execute(cmd string) {
	stdout, exit_code := execute.Execute(cmd)

	var v = variables.Load()
	v.Set(map[string]interface{}{
		"exit_code": exit_code,
		"stdout": stdout,
	})
}

func (f *Flow) ParseStdout() error {
	var v = variables.Load()

	switch f.GetFormat() {
	case "TEXT":
		if reg := f.GetRegister(); len(reg) > 0 {
			v.Set(map[string]interface{}{reg: v.Items["stdout"]})
		}
	case "JSON":
		raw, err := common.StringToJSON(common.InterfaceToString(v.Items["stdout"]))
		if err != nil {
			return err
		}

		for variable, value := range raw {
			v.Set(map[string]interface{}{variable: value})
		}
	default:
		return errors.New("Format option is invalid, please use; TEXT (default) or JSON")
	}

	return nil
}

func (f *Flow) RetryCommand() bool {
	// Log possible error and retry it is true the error:
	var v = variables.Load()

	if error := v.Get(f.GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
		log.Error(common.InterfaceToString(error), nil)
		return false
	}

	// Si el status que retorna el stdout es diferente reintenta
	if status := v.Get(f.GetRetryStatus()); common.InterfaceToString(status) == f.GetRetryDone() {
		return false
	}

	return true
}
