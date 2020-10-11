package flow

import (
	"fmt"
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
