package flow

import (
	"fmt"
	"time"
	"encoding/json"

	"github.com/swapbyt3s/lightflow/config"
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
)

func (f *Flow) GetTitle() string {
	return fmt.Sprintf(
		"TASK[%s] LOOP[%s] PIPE[%s]",
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

func (f *Flow) GetFormat() config.Format {
	if len(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format) == 0 {
		return config.TEXT
	}
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format
}

func (f *Flow) SetFormat(format config.Format) {
	f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format = format
}

func (f *Flow) GetRegister() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Register
}

func (f *Flow) GetRetryAttempts() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts
}

func (f *Flow) GetRetryWait() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait
}

func (f *Flow) GetRetryExitCode() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.ExitCode
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

func (f *Flow) GetArgAutoIncrementStartDate() string {
	val, _ := common.GetArgValJSON("ai-date", "start")

	if len(val) == 0 {
		return f.GetDefaultDate()
	}

	return val
}

func (f *Flow) GetArgAutoIncrementEndDate() string {
	val, _ := common.GetArgValJSON("ai-date", "end")

	if len(val) == 0 {
		return f.GetDefaultDate()
	}

	return val
}

func (f *Flow) GetDefaultDate() string {
	return common.InterfaceToString(f.Variables.Get("date"))
}

func (f *Flow) SetDefaults() {
	f.SetFormat("TEXT")
	f.Variables.Items["path"] = config.Load().General.Temporary_Directory

	if date, _ := common.GetArgValJSON("variables", "date"); len(date) > 0 {
		f.Variables.SetDate(date)
	} else {
		f.Variables.SetDate(time.Now().Format("2006-01-02"))
	}

	f.Variables.Set(map[string]interface{} {
		"task_name": f.GetTaskName(),
		"loop_name": f.GetLoopName(),
		"pipe_name": f.GetPipeName(),
	})

	f.Variables.Set(f.GetGlobalVariables())
	f.Variables.Set(f.GetLoopVariables())
	f.Variables.Set(f.GetPipeVariables())

	args_vars := common.GetArgVal("variables").(string)

	if len(args_vars) >= 2 {
		err := json.Unmarshal([]byte(args_vars), &f.Variables.Items)
		if err != nil {
			log.Warning("Variables", map[string]interface{}{"Message": err})
		}
	}

	f.Variables.Items["error"] = ""
	f.Variables.Items["exit_code"] = 0
	f.Variables.Items["status"] = ""
	f.Variables.Items["stdout"] = ""
}

func (f *Flow) GetStdOut() interface{} {
	return f.Variables.Items["stdout"]
}
