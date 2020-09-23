package flow

import (
	"encoding/json"
	// "flag"
	"fmt"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
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
				variables.Load().SetDefaults()

				err := autoincrement.Date(
					getAutoIncrementStartDate(),
					getAutoIncrementEndDate(),
					func(date string){
						variables.Load().SetDate(date)

						Populate()

						retry.Retry(
							registry.Load().GetRetryAttempts(),
							registry.Load().GetRetryWait(),
							func() bool {
								cmd := RenderCommand()

								if common.GetArgVal("dry-run").(bool) {
									fmt.Println(cmd)

									return false
								} else {
									// Execute the command:
									log.Info(Title() + " Start...", nil)
									diff := common.Duration(func(){
										Execute(cmd)
									})
									log.Info(Title() + " End!", map[string]interface{}{"Execution Time": diff})

									// LogError
									// Log possible error and retry it is true the error:
									var v = variables.Load()
									if error := v.Get(registry.Load().GetRetryError()); error != nil && len(common.InterfaceToString(error)) > 0 {
										log.Error(Title(), map[string]interface{}{
											"Message": error,
										})
										return false
									}

									// RetryCondition?
									// Si el status que retorna el stdout es diferente reintenta
									if status := v.Get(registry.Load().GetRetryStatus()); common.InterfaceToString(status) == registry.Load().GetRetryDone() {
										return false
									}

									// Meter todo lo anterior en un RetryPipe() bool { ... return ...}
									return true
								}

								return true
							})
					})

					if err != nil {
						log.Error(err.Error(), nil)
					}
			})
		})
	})
}

func Title() string {
	return fmt.Sprintf(
		"%s%s%s",
		fmt.Sprintf("Task[%s]", registry.Load().GetTaskName()),
		fmt.Sprintf(" Loop[%s]", registry.Load().GetLoopName()),
		fmt.Sprintf(" Pipe[%s]", registry.Load().GetPipeName()),
	)
}

func Populate() {
	var v = variables.Load()
	// v.SetDefaults()

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
		v.Set(map[string]interface{}{variable: template.Render(common.TrimNewlines(value.(string)), v.Items)})
	}

	for variable, value := range v.Items {
		if value, ok := value.(string); ok {
			v.Set(map[string]interface{}{variable: template.Render(common.TrimNewlines(value), v.Items)})
		}
	}

	// Define default values:
	if format := registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format; len(format) == 0 {
		registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format = "TEXT"
	}
}

func Execute(cmd string) {
	stdout, exit_code := command.Execute(cmd)

	var v = variables.Load()
	v.Set(map[string]interface{}{"exit_code": exit_code})
	v.Set(map[string]interface{}{"stdout": stdout})

	// ParseStdOut
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
	Debug()

	// Print command output logging by error code:
	if exit_code > 0 {
		log.Error(Title(), map[string]interface{}{
			"ExitCode": exit_code,
			"Stdout": stdout,
		})
	}
}

func Debug() {
	for variable, value := range variables.Load().Items {
		log.Debug(Title() + " Variable", map[string]interface{}{
			variable: value,
		})
	}
}

func getAutoIncrementStartDate() string {
	val, _ := common.GetArgValJSON("ai-date", "start")

	if len(val) == 0 {
		return GetDefaultDate()
	}

	return val
}

func getAutoIncrementEndDate() string {
	val, _ := common.GetArgValJSON("ai-date", "end")

	if len(val) == 0 {
		return GetDefaultDate()
	}

	return val
}

func GetDefaultDate() string {
	return common.InterfaceToString(variables.Load().Get("date"))
}

func RenderCommand() string {
	var v = variables.Load()

	// Hay que simplificar esto registry.Load().GetCommand()
	var cmd = registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Command

	// Validate to have all template variables defined:
	for _, variable := range template.Variables(cmd) {
		if v.Exist(variable) == false {
			log.Warning(Title(), map[string]interface{}{
				"Message": "This variable is not defined",
				"VariableName": variable,
			})
		}
	}

	cmd = template.Render(cmd, v.Items)
	cmd = common.TrimNewlines(cmd) // <- meterlo dentro del render?

	return cmd
}
