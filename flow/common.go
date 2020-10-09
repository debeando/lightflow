package flow

import (
	"fmt"
)

func (f *Flow) GetTitle() string {
	return fmt.Sprintf(
		"TASK[%s] SUB TASK[%s] PIPE[%s]",
		f.GetTaskName(),
		f.GetLoopName(),
		f.GetPipeName(),
	)
}

func (f *Flow) GetTaskName() string {
	return f.Config.Tasks[f.Index.Task].Name
}

func (f *Flow) GetLoopName() string {
	return f.Config.Tasks[f.Index.Task].Subtask[f.Index.Subtask].Name
}

func (f *Flow) GetPipeName() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Name
}
