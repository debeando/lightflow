package flow

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

func LoadLoops(index int) {
	tasks := config.File.Tasks

	if len(tasks[index].Loops) > 0 {
		if common.IsArgDefined("looping") {
			Looping(getLoopingIndex(index))
		} else {
			Loops(index)
		}
	} else {
		LoadPipes(index, -1)
	}
}

func Loops(index int) {
	tasks := config.File.Tasks

	for looping := range tasks[index].Loops {
		LoadPipes(index, looping)
	}
}

func Looping(index int, looping int) {
	tasks := config.File.Tasks

	if looping >=0 && looping < len(tasks[index].Loops) {
		LoadPipes(index, looping)
	}
}

func getLoopingIndex(index int) (int, int) {
	tasks := config.File.Tasks
	arg   := common.GetArgVal("looping")

	for looping := range tasks[index].Loops {
		if name := tasks[index].Loops[looping]["name"]; len(name) > 0 {
			if name == arg {
				return index, looping
			}
		}
	}

	return index, -1
}
