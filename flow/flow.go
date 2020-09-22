package flow

import (
	"encoding/json"
	"flag"
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
	l := loops.Loop{}
	p := pipes.Pipe{}

	t.Run(func() {
		l.Run(func() {
			p.Run(func(){
				Populate()

				retry.Retry(registry.Load().GetRetryAttempts(), func() bool {
					if flag.Lookup("dry-run") != nil && flag.Lookup("dry-run").Value.(flag.Getter).Get().(bool) {
						var cmd = registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Command

						// Render command with variables:
						cmd = template.Render(cmd)
						cmd = common.TrimNewlines(cmd)

						log.Warning(Title(), map[string]interface{}{
							"Safe Command (dry-run)": cmd,
						})

						return false
					} else {
						// Execute the command:
						log.Info(Title() + " Start", nil)
						diff := common.Duration(func(){
							Execute()
						})
						log.Info(Title() + " End", map[string]interface{}{"ExecutionTime": diff})

						// Log possible error and retry it is true the error:
						var v = variables.Load()
						if error := v.Get(registry.Load().GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
							log.Error(Title(), map[string]interface{}{
								"Message": error,
							})
							return false
						}

						if status := v.Get(registry.Load().GetRetryStatus()); common.InterfaceToString(status) == registry.Load().GetRetryDone() {
							log.Info(Title(), map[string]interface{}{
								"Message": "Finish retry, apparently everything is fine.",
							})
							return false
						}

						return true
					}
				})
			})
		})
	})
}

func Title() string {
	var task string
	var loop string
	var pipe string

	task = fmt.Sprintf("Task[%s]", registry.Load().GetTaskName())
	loop = fmt.Sprintf(" Loop[%s]", registry.Load().GetLoopName())
	pipe = fmt.Sprintf(" Pipe[%s]", registry.Load().GetPipeName())

	return fmt.Sprintf(
		"%s%s%s",
		task,
		loop,
		pipe,
	)
}

func Populate() {
	var v = variables.Load()

	// Set default variables abour flow: task, loop, pipe.
	v.Set(map[string]interface{} {"task_name": registry.Load().GetTaskName()})
	v.Set(map[string]interface{} {"loop_name": registry.Load().GetLoopName()})
	v.Set(map[string]interface{} {"pipe_name": registry.Load().GetPipeName()})

	// Add variables from Loops:
	if len(registry.Load().Config.Tasks[registry.Load().Task].Loops) > 0 {
		for variable, value := range registry.Load().Config.Tasks[registry.Load().Task].Loops[registry.Load().Loop] {
			v.Set(map[string]interface{}{variable: common.TrimNewlines(string(value))})
		}
	}

	// Store config variables in memory:
	v.Set(registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Variables)

	// Render only variables with variables:
	// No me gusta esta parte, hay que mejorarla, los dos for:
	for variable, value := range registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Variables {
		v.Set(map[string]interface{}{variable: template.Render(common.TrimNewlines(value.(string)))})
	}

	for variable, value := range v.Items {
		if value, ok := value.(string); ok {
			v.Set(map[string]interface{}{variable: template.Render(common.TrimNewlines(value))})
		}
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
		"ExitCode": exit_code,
		"Stdout": stdout,
		"Variables": variables.Load().Items,
	})

	// Print command output logging by error code:
	if exit_code > 0 {
		log.Error(Title(), map[string]interface{}{
			"ExitCode": exit_code,
			"Stdout": stdout,
		})
	}
}
