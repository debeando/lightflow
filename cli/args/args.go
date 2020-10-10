package args

import (
	"strings"

	"github.com/debeando/lightflow/common"
)

func Task() string {
	return common.GetArgVal("task").(string)
}

func Subtask() string {
	return common.GetArgVal("subtask").(string)
}

func Pipe() string {
	return common.GetArgVal("pipes").(string)
}

func Pipes() []string {
	return strings.Split(Pipe(), ",")
}
