package plugins

// import (
// 	"fmt"

// 	"github.com/debeando/lightflow/common/log"
// 	"github.com/debeando/lightflow/plugins/evaluate"
// )

// // When a condition is true allow execute pipe.
// func (core *Core) when() bool {
// 	if len(core.GetProperty("When")) == 0 {
// 		return true
// 	}

// 	expression := core.Render(core.GetProperty("When"))
// 	value := evaluate.Expression(expression)

// 	debug_vars := make(map[string]interface{})
// 	debug_vars["Expression"] = core.GetProperty("When")
// 	debug_vars["Rendered"] = expression
// 	debug_vars["Result"] = value

// 	log.Info(
// 		fmt.Sprintf(
// 			"%s/%s When: %#v",
// 			core.TaskName(),
// 			core.PipeName(),
// 			debug_vars["Result"],
// 		),
// 		nil,
// 	)

// 	log.Debug(
// 		fmt.Sprintf(
// 			"%s/%s When: %s => %s => %#v",
// 			core.TaskName(),
// 			core.PipeName(),
// 			debug_vars["Expression"],
// 			debug_vars["Rendered"],
// 			debug_vars["Result"],
// 		),
// 		nil,
// 	)

// 	return value
// }
