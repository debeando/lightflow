package flow

import (
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
					f.GetTitle(),
					map[string]interface{}{
						"AutoIncrement Date": date,
				})
			}

			f.PopulateVariables()
			f.Chunks()
		})
}
