package flow

import (
	"encoding/json"
	"fmt"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/flow/loops"
	"github.com/swapbyt3s/lightflow/flow/pipes"
	"github.com/swapbyt3s/lightflow/flow/retry"
	"github.com/swapbyt3s/lightflow/flow/tasks"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/registry"
	"github.com/swapbyt3s/lightflow/variables"
)

func Run() {
	t := tasks.Task{}
	l := loops.Looping{}
	p := pipes.Pipe{}

	t.Run(func() {
		l.Run(func() {
			p.Run(func(){
				Populate()

				retry.Retry(registry.Load().GetRetryAttempts(), func() bool {
					// Execute the command:
					Execute()


					// Log possible error and retry it is true the error:
					var v = variables.Load()
					if error := v.Get(registry.Load().GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
						log.Error(Title(), map[string]interface{}{
							"Message": error,
						})
						return false
					}

					if status := v.Get(registry.Load().GetRetryStatus()); common.InterfaceToString(status) == registry.Load().GetRetryDone() {
						return false
					}

					return true
				})
			})
		})
	})
}

func Title() string {
	var task string
	var looping string
	var pipe string

	task = fmt.Sprintf("Task[%s]", registry.Load().GetTaskName())
	looping = fmt.Sprintf(" Looping[%s]", registry.Load().GetLoopingName())
	pipe = fmt.Sprintf(" Pipe[%s]", registry.Load().GetPipeName())

	return fmt.Sprintf(
		"%s%s%s",
		task,
		looping,
		pipe,
	)
}

func Populate() {
	var v = variables.Load()

	// Set default variables abour flow: task, looping, pipe.
	v.Set(map[string]interface{} {"task_name": registry.Load().GetTaskName()})
	v.Set(map[string]interface{} {"looping_name": registry.Load().GetLoopingName()})
	v.Set(map[string]interface{} {"pipe_name": registry.Load().GetPipeName()})

	// Add variables from Loops:
	if len(registry.Load().Config.Tasks[registry.Load().Task].Loops) > 0 {
		for variable, value := range registry.Load().Config.Tasks[registry.Load().Task].Loops[registry.Load().Looping] {
			v.Set(map[string]interface{}{variable: string(value)})
		}
	}

	// Store config variables in memory:
	v.Set(registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Variables)

	// Render only variables with variables:
	for variable, value := range registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Variables {
		v.Set(map[string]interface{}{variable: template.Render(value.(string))})
	}
}

func Execute() {
	var v = variables.Load()
	var cmd = registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Command

	cmd  = common.TrimNewlines(cmd)

	// Define default values:
	if format := registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format; len(format) == 0 {
		registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format = "TEXT"
	}

 	// Validate to have all template variables defined:
 	for _, variable := range template.Variables(cmd) {
 		if v.Exist(variable) == false {
 			log.Warning(Title(), map[string]interface{}{
 				"Message": "This variable is not defined",
 				"VariableName": variable,
 			})
 		}
 	}

	// Render command with variables:
	var c = template.Render(cmd)

	// Execute command:
	stdout, exit_code := command.Execute(c)

	v.Set(map[string]interface{}{"exit_code": exit_code})

	// Save the stdout on the register:
	switch registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format {
	case "TEXT":
		if reg := registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Register; len(reg) > 0 {
			v.Set(map[string]interface{}{reg: stdout})
		}
	case "JSON":
		var raw map[string]interface{}
		if err := json.Unmarshal([]byte(stdout), &raw); err != nil {
			log.Error(Title(), map[string]interface{}{
				"Message": "Format JSON invalid or Format type is invalid",
				"Error": err,
			})
		} else {
			for variable, value := range raw {
				v.Set(map[string]interface{}{variable: value})
				log.Debug(Title(), map[string]interface{}{
					variable: value,
				})
			}
		}
	default:
		log.Warning(Title(), map[string]interface{}{
			"Message": "Format option is invalid, please use; TEXT (default) or JSON",
			"Format": registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format,
		})
	}

	// Log all variables:
	log.Debug(Title(), map[string]interface{}{
		"Variables": variables.Load().Items,
	})

	// Print command output logging by error code:
	msg := map[string]interface{}{
		"Name": Title(),
		"stdout": stdout,
		"Exit Code": exit_code,
	}

	if exit_code == 0 {
		log.Info(Title(), msg)
	} else {
		log.Error(Title(), msg)
	}
}
