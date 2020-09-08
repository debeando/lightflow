package flow

import (
	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/register"
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
		var r = register.Load()
		var v = variables.Load()
		var c string

		name := tasks[task].Pipes[pipe].Name
		cmd := tasks[task].Pipes[pipe].Command
		cmd = common.TrimNewlines(cmd)


		// Add variables from Loops:
		if len(tasks[task].Loops) > 0 {
			for variable, value := range tasks[task].Loops[looping] {
				v.Set(map[string]string{variable: template.Render(string(value))})
			}
		}

		// Store variables in memory:
		v.Set(tasks[task].Pipes[pipe].Variables)

		// Add task and pipe name into variables memory:
		v.Set(map[string]string {
			"task_name": tasks[task].Name,
			"pipe_name": tasks[task].Pipes[pipe].Name,
		})

		// Render variables:
		for variable, value := range tasks[task].Pipes[pipe].Variables {
			v.Set(map[string]string{variable: template.Render(string(value))})
		}

		// Validate to have all variables defined:
		for _, varInTmpl := range template.Variables(cmd) {
			if v.Exist(varInTmpl) == false {
				log.Warning("Template", map[string]interface{}{
					"Message": "This variable is not defined",
					"VariableName": varInTmpl,
				})
			}
		}

		// Render command:
		c = template.Render(cmd)

		// Logging debug details before excute command:
		log.Debug("Pipe", map[string]interface{}{
			"Name": name,
			"Variables": v.Items,
			"Command": cmd,
			"Rendered": c,
		})

		// Execute command:
		stdout, exit_code := command.Execute(c)
		stdout = common.TrimNewlines(stdout)

		// Save the stdout on the register:
		if reg := tasks[task].Pipes[pipe].Register; len(reg) > 0 {
			r.Save(reg, stdout)
		}

		// Print command output logging by error code:
		msg := map[string]interface{}{
			"Name": name,
			"stdout": stdout,
			"Exit Code": exit_code,
		}

		if exit_code == 0 {
			log.Info("Pipe", msg)
		} else {
			log.Error("Pipe", msg)
		}
	}
}
