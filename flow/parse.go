package flow

import (
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
)

// parseStdout verify the formant and store value(s) in variable register.
func (f *Flow) parse() {
	switch f.GetFormat() {
	case config.TEXT:
		if reg := f.GetProperty("Register"); len(reg) > 0 {
			if reg == "date" && f.Interval == false {
				f.Variables.SetDate(common.InterfaceToString(f.GetVariable("stdout")))
			}

			f.Variables.Set(map[string]interface{}{reg: f.GetVariable("stdout")})
		}
	case config.JSON:
		//f.Variables puede tener un metodo para salvar en json de forma automatica?
		raw, err := common.StringToJSON(common.InterfaceToString(f.GetVariable("stdout")))
		if err != nil {
			log.Error(err.Error(), nil)
		}

		for variable, value := range raw {
			f.Variables.Set(map[string]interface{}{variable: value})
		}
	default:
		log.Error("Format option is invalid, please use; TEXT (default) or JSON", nil)
	}
}
