package flow

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow/autoincrement"
)

func (f *Flow) AutoIncrement() error {
	f.SetDefaults()

	return autoincrement.Date(
		f.GetArgAutoIncrementStartDate(),
		f.GetArgAutoIncrementEndDate(),
		func(date string) {
			f.SetDefaults()
			if f.Variables.SetDate(date) {
				log.Info(
					fmt.Sprintf(
						"TASK[%s] SUB TASK[%s] PIPES AI DATE[%s]",
						f.GetTaskName(),
						f.GetLoopName(),
						date,
					),
					nil)
			}
			f.Pipes()
		})
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
	f.Variables.Items["path"] = config.Load().General.Temporary_Directory

	if date, _ := common.GetArgValJSON("variables", "date"); len(date) > 0 {
		f.Variables.SetDate(date)
	} else {
		f.Variables.SetDate(time.Now().Format("2006-01-02"))
	}

	f.Variables.Set(map[string]interface{}{
		"task_name": f.GetTaskName(),
		"subtask_name": f.GetLoopName(),
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

func (f *Flow) GetGlobalVariables() map[string]interface{} {
	return f.Config.Variables
}

func (f *Flow) GetLoopVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Subtask[f.Index.Subtask].Variables
}

func (f *Flow) GetPipeVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables
}
