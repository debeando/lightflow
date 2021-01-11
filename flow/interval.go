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

	if args.IntervalAIDate() {
		fmt.Println("AI")
		start := args.IntervalAIStartDate(f.Variables.GetDate())
		end := args.IntervalAIEndDate(f.Variables.GetDate())

		return date.Increment(
			start,
			end,
			func(date string) {
				f.dateLoop(date)
			})
	} else if args.IntervalADDate() {
		fmt.Println("AD")
		start := args.IntervalADStartDate(f.Variables.GetDate())
		end := args.IntervalADEndDate(f.Variables.GetDate())

		return date.Decrement(
			start,
			end,
			func(date string) {
				f.dateLoop(date)
			})
	} else {
		f.Pipes()
	}

	return nil
}

func (f *Flow) dateLoop(date string) {
	f.Skip = false
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
}
