package flow

import (
	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/flow/autoincrement"
	"github.com/swapbyt3s/lightflow/flow/template"
	"github.com/swapbyt3s/lightflow/variables"
)

func (f *Flow) AutoIncrement() {
	variables.Load().SetDefaults()
	err := autoincrement.Date(
		getAutoIncrementStartDate(),
		getAutoIncrementEndDate(),
		func(date string){
			if variables.Load().SetDate(date) {
				log.Info(
					f.GetTitle(),
					map[string]interface{}{
						"AutoIncrement Date": date,
				})
			}

			f.PopulateVariables()
			f.Chunks()
		})
	if err != nil {
		log.Error(err.Error(), nil)
	}
}

func getAutoIncrementStartDate() string {
	val, _ := common.GetArgValJSON("ai-date", "start")

	if len(val) == 0 {
		return GetDefaultDate()
	}

	return val
}

func getAutoIncrementEndDate() string {
	val, _ := common.GetArgValJSON("ai-date", "end")

	if len(val) == 0 {
		return GetDefaultDate()
	}

	return val
}

func GetDefaultDate() string {
	return common.InterfaceToString(variables.Load().Get("date"))
}

func (f *Flow) PopulateVariables() {
	var v = variables.Load()

	// Set default variables abour flow: task, loop, pipe.
	v.Set(map[string]interface{} {
		"task_name": f.GetTaskName(),
		"loop_name": f.GetLoopName(),
		"pipe_name": f.GetPipeName(),
	})

	// Add variables from Loops:
	v.Set(f.GetLoopVariables())

	// Store config variables in memory:
	v.Set(f.GetPipeVariables())

	// ----
	// Se mete o no en el get de la variable especifica?
//	total := v.Get("total")
//
//	if total.(int) > 0 {
//		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Total = total.(int)
//	}
//
//	limit := v.Get("limit")
//
//	if limit.(int) > 0 {
//		f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Chunk.Limit = limit.(int)
//	}
	// ----

	// hay que quitar el render de variables, evitar usarlo, solo usar en el execute
	for variable, value := range v.Items {
		rendered, err := template.Render(common.TrimNewlines(common.InterfaceToString(value)), v.Items)
		if err != nil {
			log.Warning(err.Error(), nil)
		}

		v.Set(map[string]interface{}{variable: rendered})
	}

	// Render only variables with variables:
	// No me gusta esta parte, hay que mejorarla, los dos for:
	for variable, value := range f.GetPipeVariables() {
		// fmt.Println(variable)

		rendered, err := template.Render(common.TrimNewlines(value.(string)), v.Items)
		if err != nil {
			log.Warning(err.Error(), nil)
		}

		v.Set(map[string]interface{}{variable: rendered})
	}

	// Define default values:
	if format := f.GetFormat(); len(format) == 0 {
		f.SetFormat("TEXT")
	}
}
