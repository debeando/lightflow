package skip

// import (
// 	"fmt"

// 	"github.com/debeando/lightflow/common/log"
// 	"github.com/debeando/lightflow/plugins/evaluate"
// )

// // Skip evaluate condition to set skip flag.
// func (core *Core) skip() {
// 	if len(core.GetProperty("Skip")) == 0 {
// 		return
// 	}

// 	expression := core.Render(core.GetProperty("Skip"))

// 	core.Skip = evaluate.Expression(expression)
// 	debug_vars := make(map[string]interface{})
// 	debug_vars["Expression"] = core.GetProperty("Skip")
// 	debug_vars["Rendered"] = expression
// 	debug_vars["Result"] = core.Skip

// 	core.Variables.Set(map[string]interface{}{
// 		"skip": core.Skip,
// 	})

// 	log.Info(
// 		fmt.Sprintf(
// 			"%s/%s Skip: %#v",
// 			core.TaskName(),
// 			core.PipeName(),
// 			debug_vars["Result"],
// 		),
// 		nil,
// 	)

// 	log.Debug(
// 		fmt.Sprintf(
// 			"%s/%s Skip: %s => %s => %#v",
// 			core.TaskName(),
// 			core.PipeName(),
// 			debug_vars["Expression"],
// 			debug_vars["Rendered"],
// 			debug_vars["Result"],
// 		),
// 		nil,
// 	)
// }
