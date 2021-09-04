package core

// import (
// 	"reflect"

// 	"github.com/debeando/lightflow/config"
// )

// // GetProperty get value string by key name from config struct.
// func (core *Core) GetProperty(name string) string {
// 	rv := reflect.ValueOf(core.Config.Pipes[core.Index.Pipe])

// 	if rv.Kind() != reflect.Struct {
// 		return ""
// 	}

// 	fv := rv.FieldByName(name)

// 	if !fv.IsValid() {
// 		return ""
// 	}

// 	val, ok := fv.Interface().(string)

// 	if !ok {
// 		return ""
// 	}

// 	return val
// }

// func (core *Core) GetFormat() config.Format {
// 	if len(core.Config.Pipes[core.Index.Pipe].Format) == 0 {
// 		return "TEXT"
// 	}
// 	return core.Config.Pipes[core.Index.Pipe].Format
// }

// func (core *Core) GetVariables() map[string]interface{} {
// 	return core.Config.Variables
// }

// func (core *Core) GetTaskVariables() map[string]interface{} {
// 	return core.Config.Tasks[core.Index.Task].Variables
// }

// func (core *Core) GetPipeVariables() map[string]interface{} {
// 	return core.Config.Pipes[core.Index.Pipe].Variables
// }

// func (core *Core) getTaskName() string {
// 	return core.Config.Tasks[core.Index.Task].Name
// }

// func (core *Core) GetRetryAttempts() int {
// 	return core.Config.Pipes[core.Index.Pipe].Retry.Attempts
// }

// func (core *Core) GetRetryWait() int {
// 	return core.Config.Pipes[core.Index.Pipe].Retry.Wait
// }

// func (core *Core) GetRetryExpression() string {
// 	return core.Config.Pipes[core.Index.Pipe].Retry.Expression
// }

// func (core *Core) GetSlackChannel() string {
// 	return core.Config.Pipes[core.Index.Pipe].Slack.Channel
// }

// func (core *Core) GetSlackTitle() string {
// 	return core.Config.Pipes[core.Index.Pipe].Slack.Title
// }

// func (core *Core) GetSlackMessage() string {
// 	return core.Config.Pipes[core.Index.Pipe].Slack.Message
// }

// func (core *Core) GetSlackColor() string {
// 	return core.Config.Pipes[core.Index.Pipe].Slack.Color
// }

// func (core *Core) GetSlackExpression() string {
// 	return core.Config.Pipes[core.Index.Pipe].Slack.Expression
// }

// func (core *Core) GetPipePrint() []string {
// 	return core.Config.Pipes[core.Index.Pipe].Print
// }

// func (core *Core) GetPipeUnset() []string {
// 	return core.Config.Pipes[core.Index.Pipe].Unset
// }
