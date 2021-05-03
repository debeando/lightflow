package flow

import (
	"reflect"

	"github.com/debeando/lightflow/config"
)

// GetProperty get value string by key name from config struct.
func (f *Flow) GetProperty(name string) string {
	rv := reflect.ValueOf(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe])

	if rv.Kind() != reflect.Struct {
		return ""
	}

	fv := rv.FieldByName(name)

	if !fv.IsValid() {
		return ""
	}

	val, ok := fv.Interface().(string)

	if !ok {
		return ""
	}

	return val
}

func (f *Flow) GetFormat() config.Format {
	if len(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format) == 0 {
		return "TEXT"
	}
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format
}

func (f *Flow) GetSubTaskVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Subtasks[f.Index.Subtask].Variables
}

func (f *Flow) GetPipeVariables() map[string]interface{} {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Variables
}

func (f *Flow) getSubTaskName() string {
	return f.Config.Tasks[f.Index.Task].Subtasks[f.Index.Subtask].Name
}

func (f *Flow) getTaskName() string {
	return f.Config.Tasks[f.Index.Task].Name
}

func (f *Flow) GetRetryAttempts() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Attempts
}

func (f *Flow) GetRetryWait() int {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Wait
}

func (f *Flow) GetRetryExpression() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Retry.Expression
}

func (f *Flow) GetSlackChannel() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Slack.Channel
}

func (f *Flow) GetSlackTitle() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Slack.Title
}

func (f *Flow) GetSlackMessage() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Slack.Message
}

func (f *Flow) GetSlackColor() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Slack.Color
}

func (f *Flow) GetSlackExpression() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Slack.Expression
}

func (f *Flow) GetPipePrint() []string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Print
}

func (f *Flow) GetPipeUnset() []string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Unset
}
