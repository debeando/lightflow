package flow

import (
	"fmt"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
	"github.com/swapbyt3s/lightflow/flow/loops"
	"github.com/swapbyt3s/lightflow/flow/pipe"
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
								cmd := pipe.RenderCommand()

								if common.GetArgVal("dry-run").(bool) {
									fmt.Println(cmd)

									return false
								} else {
									log.Info(pipe.Title() + " Start...", nil)
									diff := common.Duration(func(){
										pipe.Execute(cmd)
									})
									log.Info(pipe.Title() + " End!", map[string]interface{}{"Execution Time": diff})

									if err := pipe.ParseStdout(); err != nil {
										log.Error(err.Error(), nil)
									}

									pipe.Debug()

									return pipe.Retry()
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
	if format := registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format; len(format) == 0 {
		registry.Load().Config.Tasks[registry.Load().Task].Pipes[registry.Load().Pipe].Format = "TEXT"
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
