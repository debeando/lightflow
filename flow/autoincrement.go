package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
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
