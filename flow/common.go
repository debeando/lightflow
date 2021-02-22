package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/config"
)

func (f *Flow) GetTitle() string {
	return fmt.Sprintf(
		"TASK[%s] SUB TASK[%s] PIPE[%s]",
		f.TaskName(),
		f.SubTaskName(),
		f.PipeName(),
	)
}

func (f *Flow) TaskName() string {
	return f.setName("task_name", f.getTaskName())
}

func (f *Flow) SubTaskName() string {
	return f.setName("subtask_name", f.getSubTaskName())
}

func (f *Flow) PipeName() string {
	return f.setName("pipe_name", f.GetProperty("Name"))
}

func (f *Flow) GetVariable(name string) interface{} {
	return f.Variables.Items[name]
}

func (f *Flow) GetGlobalVariables() map[string]interface{} {
	return f.Config.Variables
}

func (f *Flow) SetDefaults() {
	f.Variables.Set(args.Variables())
	f.Variables.Set(f.GetGlobalVariables())
	f.Variables.Set(f.GetPipeVariables())
	f.Variables.Set(f.GetSubTaskVariables())
	f.Variables.Set(map[string]interface{}{
		"error":	 "",
		"exit_code": 0,
		"limit":     0,
		"offset":    0,
		"path":      config.Load().General.Temporary_Directory,
		"skip":      false,
		"stdout":    "",
	})
	f.Variables.SetDate(args.VariableDate())
}

func (f *Flow) setName(key string, value string) string {
	f.Variables.Set(map[string]interface{}{
		key: value,
	})

	return value
}
