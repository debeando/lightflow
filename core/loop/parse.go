package core

// import (
// 	"github.com/debeando/lightflow/common"
// 	"github.com/debeando/lightflow/common/log"
// 	"github.com/debeando/lightflow/config"
// )

// // parseStdout verify the formant and store value(s) in variable register.
// func (core *Core) parse() {
// 	switch core.GetFormat() {
// 	case config.TEXT:
// 		if reg := core.GetProperty("Register"); len(reg) > 0 {
// 			if reg == "date" && core.Interval == false {
// 				core.Variables.SetDate(common.InterfaceToString(core.GetVariable("stdout")))
// 			}

// 			core.Variables.Set(map[string]interface{}{reg: core.GetVariable("stdout")})
// 		}
// 	case config.JSON:
// 		//core.Variables puede tener un metodo para salvar en json de forma automatica?
// 		raw, err := common.StringToJSON(common.InterfaceToString(core.GetVariable("stdout")))
// 		if err != nil {
// 			log.Error(err.Error(), nil)
// 		}

// 		for variable, value := range raw {
// 			core.Variables.Set(map[string]interface{}{variable: value})
// 		}
// 	default:
// 		log.Error("Format option is invalid, please use; TEXT (default) or JSON", nil)
// 	}
// }
