package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/iterator"
)

func (f *Flow) Tasks() {
	itr := iterator.Iterator{
		Items: f.Config.Tasks,
		Name:  args.Task(),
	}

	log.Info(
		fmt.Sprintf(
			"Starting...",
		), nil)

	itr.Run(func() bool {
		f.Index.Pipe = 0
		f.Index.Task = itr.Index
		f.Skip = false

		if err := f.Date(); err != nil {
			log.Error(err.Error(), nil)
		}
		return false
	})

	log.Info(
		fmt.Sprintf(
			"Finished %s", // ET is acronym for execution time.
			itr.ExecutionTime,
		), nil)
}
