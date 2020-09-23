package pipe

import (
	"fmt"
	"errors"

	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/registry"
	"github.com/swapbyt3s/lightflow/variables"
	"github.com/swapbyt3s/lightflow/flow/template"
)

func Title() string {
	return fmt.Sprintf(
		"%s%s%s",
		fmt.Sprintf("Task[%s]", registry.Load().GetTaskName()),
		fmt.Sprintf(" Loop[%s]", registry.Load().GetLoopName()),
		fmt.Sprintf(" Pipe[%s]", registry.Load().GetPipeName()),
	)
}

func Execute(cmd string) {
	stdout, exit_code := command.Execute(cmd)

	var v = variables.Load()
	v.Set(map[string]interface{}{"exit_code": exit_code})
	v.Set(map[string]interface{}{"stdout": stdout})
}

func ParseStdout() error {
	var v = variables.Load()

	switch registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format {
	case "TEXT":
		if reg := registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Register; len(reg) > 0 {
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

func Debug() {
	for variable, value := range variables.Load().Items {
		switch variable {
		case "exit_code":
			if value.(int) > 0 {
				log.Error(Title(), map[string]interface{}{
					"Exit Code": value,
				})
			}
			break
		default:
			log.Debug(Title(), map[string]interface{}{
				variable: value,
			})
		}
	}
}

func RenderCommand() string {
	var v = variables.Load()

	var cmd = registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Command

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

func Retry() bool {
	// Log possible error and retry it is true the error:
	var v = variables.Load()
	if error := v.Get(registry.Load().GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
		log.Error(common.InterfaceToString(error), nil)
		return false
	}

	// Si el status que retorna el stdout es diferente reintenta
	if status := v.Get(registry.Load().GetRetryStatus()); common.InterfaceToString(status) == registry.Load().GetRetryDone() {
		return false
	}

	return true
}
