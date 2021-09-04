package core

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/date"
)

// Date is the handle method for generic Auto Increment or Decrement package.
func (core *Core) Date() error {
	core.SetDefaults()
	core.Interval = false

	if args.IntervalAIDate() {
		start := args.IntervalAIStartDate(core.Variables.GetDate())
		end := args.IntervalAIEndDate(core.Variables.GetDate())

		return date.Increment(
			start,
			end,
			func(date string) {
				core.dateLoop(date)
			})
	} else if args.IntervalADDate() {
		start := args.IntervalADStartDate(core.Variables.GetDate())
		end := args.IntervalADEndDate(core.Variables.GetDate())

		return date.Decrement(
			start,
			end,
			func(date string) {
				core.dateLoop(date)
			})
	} else {
		core.PipesInArgs()
	}

	return nil
}

func (core *Core) dateLoop(date string) {
	core.Skip = false
	core.Interval = true
	core.SetDefaults()
	if core.Variables.SetDate(date) {
		log.Info(
			fmt.Sprintf(
				"%s Increment %s",
				core.TaskName(),
				date,
			),
			nil)
	}
	core.PipesInArgs()
}
