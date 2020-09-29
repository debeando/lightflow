package flow

import (
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
	"github.com/swapbyt3s/lightflow/variables"
)

func (f *Flow) AutoIncrement() error {
	variables.Load().SetDefaults()
	return autoincrement.Date(
		getAutoIncrementStartDate(),
		getAutoIncrementEndDate(),
		func(date string){
			if variables.Load().SetDate(date) {
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
