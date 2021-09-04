package core

import (
	"fmt"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/plugins/template"
)

func (core *Core) GetTitle() string {
	return fmt.Sprintf(
		"%s/%s Starting...",
		core.TaskName(),
		core.PipeName(),
	)
}

func (core *Core) TaskName() string {
	return core.setName("task_name", core.getTaskName())
}

func (core *Core) PipeName() string {
	return core.setName("pipe_name", core.GetProperty("Name"))
}

func (core *Core) GetVariable(name string) interface{} {
	return core.Variables.Items[name]
}

func (core *Core) SetDefaults() {
	core.Variables.Set(args.Variables())
	core.Variables.Set(core.GetVariables())
	core.Variables.Set(core.GetTaskVariables())
	core.Variables.Set(core.GetPipeVariables())
	core.Variables.Set(map[string]interface{}{
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
	core.Variables.SetDate(args.VariableDate())
}

func (core *Core) setName(key string, value string) string {
	core.Variables.Set(map[string]interface{}{
		key: value,
	})

	return value
}

// Render a template with variables.
func (core *Core) Render(s string) string {
	r, err := template.Render(s, core.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
