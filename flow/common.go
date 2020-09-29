package flow

import (
	"fmt"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/variables"
)

func (f *Flow) GetTitle() string {
	return fmt.Sprintf(
		"Task[%s] Loop[%s] Pipe[%s]",
		f.GetTaskName(),
		f.GetLoopName(),
		f.GetPipeName(),
	)
}

func (f *Flow) GetTaskName() string {
	return f.Config.Tasks[f.Index.Task].Name
}

func (f *Flow) GetLoopName() string {
	return f.Config.Tasks[f.Index.Task].Loops[f.Index.Loop].Name
}

func (f *Flow) GetPipeName() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Name
}

func (f *Flow) GetGlobalVariables() map[string]interface{} {
	return f.Config.Variables
}

func (f *Flow) GetLoopVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Loops[f.Index.Loop].Variables
}

func (f *Flow) GetPipeVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables
}

func (f *Flow) GetExecute() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Execute
}

func (f *Flow) GetFormat() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format
}

func (f *Flow) SetFormat(format string) {
	f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format = format
}

func (f *Flow) GetRegister() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Register
}

func (f *Flow) GetChunkTotal() int {
	var v = variables.Load()
	total := v.Get("chunk.total")

	if total != nil {
		if value, ok := total.(string); ok {
			if value := common.StringToInt(value); value > 0 {
				f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total = value
			}
		}
	}

	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total
}

func (f *Flow) GetChunkLimit() int {
	// validar que sea int y hacer un warn cuando no es un ent

	var v = variables.Load()
	limit := v.Get("chunk.limit")

	if limit != nil {
		if value, ok := limit.(string); ok {
			if value := common.StringToInt(value); value > 0 {
				f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit = value
			}
		}
	}

	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit
}

func (f *Flow) GetRetryAttempts() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts
}

func (f *Flow) GetRetryWait() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait
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

func (f *Flow) IsValidChunk() bool {
	if f.GetChunkTotal() < 2 {
		return false
	}
	if f.GetChunkLimit() < 2 {
		return false
	}

	return true
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

func (f *Flow) PopulateVariables() {
	var v = variables.Load()

	// Set default variables abour flow: task, loop, pipe.
	v.Set(map[string]interface{} {
		"task_name": f.GetTaskName(),
		"loop_name": f.GetLoopName(),
		"pipe_name": f.GetPipeName(),
	})

	// Add global variables:
	v.Set(f.GetGlobalVariables())

	// Add variables from Loops:
	v.Set(f.GetLoopVariables())

	// Store config variables in memory:
	v.Set(f.GetPipeVariables())

	// ----
	// Se mete o no en el get de la variable especifica?
//	total := v.Get("total")
//
//	if total.(int) > 0 {
//		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total = total.(int)
//	}
//
//	limit := v.Get("limit")
//
//	if limit.(int) > 0 {
//		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit = limit.(int)
//	}
	// ----

	// hay que quitar el render de variables, evitar usarlo, solo usar en el execute
	for variable, value := range v.Items {
		rendered, err := template.Render(common.TrimNewlines(common.InterfaceToString(value)), v.Items)
		if err != nil {
			log.Warning(err.Error(), nil)
		}

		v.Set(map[string]interface{}{variable: rendered})
	}

	// Render only variables with variables:
	// No me gusta esta parte, hay que mejorarla, los dos for:
	for variable, value := range f.GetPipeVariables() {
		// fmt.Println(variable)

		rendered, err := template.Render(common.TrimNewlines(value.(string)), v.Items)
		if err != nil {
			log.Warning(err.Error(), nil)
		}

		v.Set(map[string]interface{}{variable: rendered})
	}

	// Define default values:
	if format := f.GetFormat(); len(format) == 0 {
		f.SetFormat("TEXT")
	}
}
