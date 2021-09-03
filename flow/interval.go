package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/date"
)

// Date is the handle method for generic Auto Increment or Decrement package.
func (f *Flow) Date() error {
	f.SetDefaults()
	f.Interval = false

	if args.IntervalAIDate() {
		start := args.IntervalAIStartDate(f.Variables.GetDate())
		end := args.IntervalAIEndDate(f.Variables.GetDate())

		return date.Increment(
			start,
			end,
			func(date string) {
				f.dateLoop(date)
			})
	} else if args.IntervalADDate() {
		start := args.IntervalADStartDate(f.Variables.GetDate())
		end := args.IntervalADEndDate(f.Variables.GetDate())

		return date.Decrement(
			start,
			end,
			func(date string) {
				f.dateLoop(date)
			})
	} else {
		f.PipesInArgs()
	}

	return nil
}

func (f *Flow) dateLoop(date string) {
	f.Skip = false
	f.Interval = true
	f.SetDefaults()
	if f.Variables.SetDate(date) {
		log.Info(
			fmt.Sprintf(
				"%s Increment %s",
				f.TaskName(),
				date,
			),
			nil)
	}
	f.PipesInArgs()
}
