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

func DryRun() bool {
	dryrun := common.GetArgVal("dry-run")

	switch v := dryrun.(type){
	case bool:
		return v
	default:
		return false
	}

	return false
}

func Variables() map[string]interface{} {
	args_vars := common.GetArgVal("variables")

	switch v := args_vars.(type){
	case string:
		vars, _ := common.StringToJSON(v)

		return vars
	default:
		return nil
	}

	return nil
}

func AutoIncrementStartDate(default_date string) string {
	val, _ := common.GetArgValJSON("ai-date", "start")

	if len(val) == 0 {
		return default_date
	}

	return val
}

func AutoIncrementEndDate(default_date string) string {
	val, _ := common.GetArgValJSON("ai-date", "end")

	if len(val) == 0 {
		return default_date
	}

	return val
}

func VariableDate() string {
	if date, _ := common.GetArgValJSON("variables", "date"); len(date) > 0 {
		return date
	}
	return common.CurrentDate()
}
