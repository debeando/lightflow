package flow

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

type Flow struct{
	Tasks config.Config
	Task int
	Item int
	Pipe int
}

func Run() {
	if common.IsArgDefined("task") {
		Task(getTaskIndex())
	} else {
		Tasks()
	}
}
