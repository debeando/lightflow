package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/autoincrement"
)

func (f *Flow) AutoIncrement() error {
	f.SetDefaults()

	return autoincrement.Date(
		f.GetArgAutoIncrementStartDate(),
		f.GetArgAutoIncrementEndDate(),
		func(date string){
			f.SetDefaults()
			if f.Variables.SetDate(date) {
				log.Info(
					fmt.Sprintf(
						"TASK[%s] LOOP[%s] PIPES-AI-DATE[%s]",
						f.GetTaskName(),
						f.GetLoopName(),
						date,
					),
					nil)
			}
			f.Pipes()
		})
}
