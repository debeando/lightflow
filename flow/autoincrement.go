package flow

import (
	"fmt"

	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
)

func (f *Flow) AutoIncrement() error {
	f.Variables.SetDefaults()

	return autoincrement.Date(
		f.GetAutoIncrementStartDate(),
		f.GetAutoIncrementEndDate(),
		func(date string){
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

			f.SetDefaults()
			f.Pipes()
		})
}
