package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow/autoincrement"
)

func (f *Flow) AutoIncrement() error {
	f.SetDefaults()

	return autoincrement.Date(
		args.AutoIncrementStartDate(f.Variables.GetDate()),
		args.AutoIncrementEndDate(f.Variables.GetDate()),
		func(date string) {
			f.SetDefaults()
			if f.Variables.SetDate(date) {
				log.Info(
					fmt.Sprintf(
						"TASK[%s] SUB TASK[%s] PIPES AI DATE[%s]",
						f.TaskName(),
						f.SubTaskName(),
						date,
					),
					nil)
			}
			f.Pipes()
		})
}

func (f *Flow) SetDefaults() {
	f.Variables.SetDate(args.VariableDate())
	f.Variables.Set(f.GetGlobalVariables())
	f.Variables.Set(f.GetSubTaskVariables())
	f.Variables.Set(f.GetPipeVariables())
	f.Variables.Set(args.Variables())

	f.Variables.Items["path"] = config.Load().General.Temporary_Directory
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

func (f *Flow) GetSubTaskVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Subtask[f.Index.Subtask].Variables
}

func (f *Flow) GetPipeVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables
}
