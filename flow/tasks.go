package flow

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/config"
)

func getTaskIndex() int {
	tasks := config.File.Tasks
	name  := common.GetArgVal("task")

	for index := range tasks {
		if tasks[index].Name == name {
			return index
		}
	}
	return -1
}

func Tasks() {
	tasks := config.File.Tasks

	for index := range tasks {
		Task(index)
	}
}

func Task(index int) {
	tasks := config.File.Tasks

	if index >=0 && index < len(tasks) {
		LoadLoops(index)
	}
}
