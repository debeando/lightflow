package flow

import (
	"fmt"
	"encoding/json"

	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/variables"
)

func LoadPipes(index int, looping int) {
	if common.IsArgDefined("pipe") {
		Pipe(getPipeIndex(index, looping))
	} else {
		Pipes(index, looping)
	}
}

func getPipeIndex(task int, looping int) (int, int, int) {
	tasks := config.File.Tasks
	name  := common.GetArgVal("pipe")

	for index := range tasks[task].Pipes {
		if tasks[task].Pipes[index].Name == name {
			return task, index, looping
		}
	}
	return task, -1, looping
}

func Pipes(task int, looping int) {
	tasks := config.File.Tasks

	for pipe := range tasks[task].Pipes {
		Pipe(task, pipe, looping)
	}
}

func Pipe(task int, pipe int, looping int) {
	tasks := config.File.Tasks

	if pipe >=0 && pipe < len(tasks[task].Pipes) {
		var v = variables.Load()
		var c string
		var title string
		var looping_name string
		var looping_log_name string

		name := tasks[task].Pipes[pipe].Name
		cmd := tasks[task].Pipes[pipe].Command
		cmd = common.TrimNewlines(cmd)

		// Define default value for format:
		if format := tasks[task].Pipes[pipe].Format; len(format) == 0 {
			tasks[task].Pipes[pipe].Format = "TEXT"
		}

		// Define looping name:
		if looping >=0 && looping < len(tasks[task].Loops) {
			if name := tasks[task].Loops[looping]["name"]; len(name) > 0 {
				looping_name = name
				looping_log_name = fmt.Sprintf("Looping[%s] ", name)
			}
		}

		// Add variables from Loops:
		if len(tasks[task].Loops) > 0 {
			for variable, value := range tasks[task].Loops[looping] {
				v.Set(map[string]interface{}{variable: template.Render(string(value))})
			}
		}

		// Store variables in memory:
		v.Set(tasks[task].Pipes[pipe].Variables)

		// Add task and pipe name into variables memory:
		v.Set(map[string]interface{} {
			"task_name": tasks[task].Name,
			"looping_name": looping_name,
			"pipe_name": tasks[task].Pipes[pipe].Name,
		})

		// Render only variables with variables:
		for variable, value := range tasks[task].Pipes[pipe].Variables {
			v.Set(map[string]interface{}{variable: template.Render(value.(string))})
		}

		// Validate to have all template variables defined:
		for _, varInTmpl := range template.Variables(cmd) {
			if v.Exist(varInTmpl) == false {
				log.Warning(title, map[string]interface{}{
					"Message": "This variable is not defined",
					"VariableName": varInTmpl,
				})
			}
		}

		// Render command with variables:
		c = template.Render(cmd)

		// Define title for logs:
		title = fmt.Sprintf(
			"Task[%s] %sPipe[%s]",
			tasks[task].Name,
			looping_log_name,
			tasks[task].Pipes[pipe].Name,
		)

		// Execute command:
		stdout, exit_code := command.Execute(c)
		stdout = common.TrimNewlines(stdout)

		// Save the stdout on the register:
		switch tasks[task].Pipes[pipe].Format {
		case "TEXT":
			if reg := tasks[task].Pipes[pipe].Register; len(reg) > 0 {
				v.Set(map[string]interface{}{reg: stdout})
			}
		case "JSON":
			var raw map[string]interface{}
			if err := json.Unmarshal([]byte(stdout), &raw); err != nil {
				log.Error(title, map[string]interface{}{
					"Message": "Format JSON invalid or Format type is invalid",
					"Error": err,
				})
			} else {
				for variable, value := range raw {
					v.Set(map[string]interface{}{variable: value})
				}
			}
		default:
			log.Warning(title, map[string]interface{}{
				"Message": "Format option is invalid, please use; TEXT (default) or JSON",
				"Format": tasks[task].Pipes[pipe].Format,
			})
		}

		// Logging debug details before excute command:
		log.Debug(title, map[string]interface{}{
			"Variables": v.Items,
		})

		// Print command output logging by error code:
		msg := map[string]interface{}{
			"Name": name,
			"stdout": stdout,
			"Exit Code": exit_code,
		}

		if exit_code == 0 {
			log.Info(title, msg)
		} else {
			log.Error(title, msg)
		}
	}
}
