package flow

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow/template"
)

func (f *Flow) GetTitle() string {
	return fmt.Sprintf(
		"%s/%s Starting...",
		f.TaskName(),
		f.PipeName(),
	)
}

func (f *Flow) TaskName() string {
	return f.setName("task_name", f.getTaskName())
}

func (f *Flow) PipeName() string {
	return f.setName("pipe_name", f.GetProperty("Name"))
}

func (f *Flow) GetVariable(name string) interface{} {
	return f.Variables.Items[name]
}

func (f *Flow) SetDefaults() {
	f.Variables.Set(args.Variables())
	f.Variables.Set(f.GetVariables())
	f.Variables.Set(f.GetTaskVariables())
	f.Variables.Set(f.GetPipeVariables())
	f.Variables.Set(map[string]interface{}{
		"chunk_limit":         0,
		"chunk_offset":        0,
		"chunk_step":          0,
		"chunk_total":         0,
		"chunk_end":           0,
		"error":               "",
		"exit_code":           0,
		"path":                config.Load().General.Temporary_Directory,
		"skip":                false,
		"stdout":              "",
		"aws_s3_objects_size": 0,
	})
	f.Variables.SetDate(args.VariableDate())
}

func (f *Flow) setName(key string, value string) string {
	f.Variables.Set(map[string]interface{}{
		key: value,
	})

	return value
}

// Render a template with variables.
func (f *Flow) Render(s string) string {
	r, err := template.Render(s, f.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
