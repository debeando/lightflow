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
	f.setTaskName()
	return f.getTaskName()
}

func (f *Flow) getTaskName() string {
	return f.Config.Tasks[f.Index.Task].Name
}

func (f *Flow) setTaskName() {
	f.Variables.Set(map[string]interface{}{
		"task_name": f.getTaskName(),
	})
}

func (f *Flow) SubTaskName() string {
	f.setSubTaskName()
	return f.getSubTaskName()
}

func (f *Flow) setSubTaskName() {
	f.Variables.Set(map[string]interface{}{
		"subtask_name": f.getSubTaskName(),
	})
}

func (f *Flow) getSubTaskName() string {
	return f.Config.Tasks[f.Index.Task].Subtask[f.Index.Subtask].Name
}

func (f *Flow) PipeName() string {
	f.setPipeName()
	return f.getPipeName()
}

func (f *Flow) setPipeName() {
	f.Variables.Set(map[string]interface{}{
		"pipe_name": f.getPipeName(),
	})
}

func (f *Flow) getPipeName() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Name
}

func (f *Flow) GetStdOut() interface{} {
	return f.Variables.Items["stdout"]
}

func (f *Flow) GetGlobalVariables() map[string]interface{} {
	return f.Config.Variables
}

func (f *Flow) GetSubTaskVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Subtask[f.Index.Subtask].Variables
}

func (f *Flow) GetPipeVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables
}

func (f *Flow) SetDefaults() {
	f.Variables.Set(args.Variables())
	f.Variables.Set(f.GetGlobalVariables())
	f.Variables.Set(f.GetPipeVariables())
	f.Variables.Set(f.GetSubTaskVariables())
	f.Variables.SetDate(args.VariableDate())

	f.Variables.Items["error"] = ""
	f.Variables.Items["exit_code"] = 0
	f.Variables.Items["limit"] = 0
	f.Variables.Items["offset"] = 0
	f.Variables.Items["path"] = config.Load().General.Temporary_Directory
	f.Variables.Items["status"] = ""
	f.Variables.Items["stdout"] = ""
}
