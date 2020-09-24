package flow

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/swapbyt3s/lightflow/command"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
	"github.com/swapbyt3s/lightflow/flow/loops"
	"github.com/swapbyt3s/lightflow/flow/pipes"
	"github.com/swapbyt3s/lightflow/flow/retry"
	"github.com/swapbyt3s/lightflow/flow/tasks"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/variables"
)

type Flow struct {
	Index Index
	Config config.Structure
}

type Index struct {
	Task int
	Loop int
	Pipe int
}

func (f *Flow) Run() {
	f.Config = *config.Load()
	f.Task()
}

func (f *Flow) Task() {
	tsk := tasks.Task{
		Config: f.Config,
	}

	if err := tsk.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	err := tsk.Run(func() {
		f.Index.Task = tsk.Index
		f.Loop()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"Tasks",
			map[string]interface{}{
				"Execution Time": tsk.ExecutionTime,
			})
	}
}

func (f *Flow) Loop() {
	loop := loops.Loop{
		Task: f.Index.Task,
		Config: f.Config,
	}

	if err := loop.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	err := loop.Run(func() {
		f.Index.Loop = loop.Index
		f.Pipes()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"Loop",
			map[string]interface{}{
				"Execution Time": loop.ExecutionTime,
			})
	}
}

func (f *Flow) Pipes() {
	pipe := pipes.Pipe{
		Task: f.Index.Task,
		Loop: f.Index.Loop,
		Config: f.Config,
	}

	if err := pipe.Valid(); err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}


	err := pipe.Run(func() {
		log.Info(pipe.Title, nil)
		f.Index.Pipe =  pipe.Index
		f.AutoIncrement()
	})
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	} else {
		log.Info(
			"Pipe",
			map[string]interface{}{
				"Execution Time": pipe.ExecutionTime,
			})
	}
}

func (f *Flow) AutoIncrement() {
	variables.Load().SetDefaults()
	err := autoincrement.Date(
		getAutoIncrementStartDate(),
		getAutoIncrementEndDate(),
		func(date string){
			variables.Load().SetDate(date)
			f.PopulateVariables()
			f.Retry()
		})
	if err != nil {
		log.Error(err.Error(), nil)
	}
}

func (f *Flow) Retry() {
	retry.Retry(
		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts,
		time.Duration(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait),
		func() bool {
			return f.ExecuteCommand()
		})

	f.PrintDebugVariables()
}

func (f *Flow) ExecuteCommand() bool {
	cmd := f.RenderCommand()

	if common.GetArgVal("dry-run").(bool) {
		fmt.Println(cmd)

		return false
	} else {
		f.Execute(cmd)

		if err := f.ParseStdout(); err != nil {
			log.Error(err.Error(), nil)
		}

		return f.RetryCommand()
	}

	return false
}

func (f *Flow) RenderCommand() string {
	var v = variables.Load()

	var cmd = f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Command

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
	stdout, exit_code := command.Execute(cmd)

	var v = variables.Load()
	v.Set(map[string]interface{}{
		"exit_code": exit_code,
		"stdout": stdout,
	})
}

func (f *Flow) ParseStdout() error {
	var v = variables.Load()

	switch f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format {
	case "TEXT":
		if reg := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Register; len(reg) > 0 {
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


func (f *Flow) PopulateVariables() {
	var v = variables.Load()

	// Set default variables abour flow: task, loop, pipe.
	v.Set(map[string]interface{} {
		"task_name": f.Config.Tasks[f.Index.Task].Name,
		"loop_name": f.Config.Tasks[f.Index.Task].Loops[f.Index.Loop].Name,
		"pipe_name": f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Name,
	})

	// Add variables from Loops:
	v.Set(f.Config.Tasks[f.Index.Task].Loops[f.Index.Loop].Variables)

	// Store config variables in memory:
	v.Set(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables)

	// Render only variables with variables:
	// No me gusta esta parte, hay que mejorarla, los dos for:
	for variable, value := range f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables {
		rendered, err := template.Render(common.TrimNewlines(value.(string)), v.Items)
		if err != nil {
			log.Warning(err.Error(), nil)
		}

		v.Set(map[string]interface{}{variable: rendered})
	}

	for variable, value := range v.Items {
		if value, ok := value.(string); ok {
			rendered, err := template.Render(common.TrimNewlines(value), v.Items)
			if err != nil {
				log.Warning(err.Error(), nil)
			}

			v.Set(map[string]interface{}{variable: rendered})
		}
	}

	// Define default values:
	if format := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format; len(format) == 0 {
		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format = "TEXT"
	}
}

func (f *Flow) PrintDebugVariables() {
	title := fmt.Sprintf(
		"Task[%s] Loop[%s] Pipe[%s]",
		f.Config.Tasks[f.Index.Task].Name,
		f.Config.Tasks[f.Index.Task].Loops[f.Index.Loop].Name,
		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Name,
	)

	for variable, value := range variables.Load().Items {
		switch variable {
		case "exit_code":
			if value.(int) > 0 {
				log.Error(title, map[string]interface{}{
					"Exit Code": value,
				})
			}
			break
		default:
			log.Debug(title, map[string]interface{}{
				variable: value,
			})
		}
	}
}

func (f *Flow) GetRetryError() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Error

	if len(value) == 0 {
		value = "error"
	}

	return value
}

func (f *Flow) GetRetryStatus() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Status

	if len(value) == 0 {
		value = "status"
	}

	return value
}

func (f *Flow) GetRetryDone() string {
	value := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Done

	if len(value) == 0 {
		value = "done"
	}

	return value
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
