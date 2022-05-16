package unset

import (
	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/variables"
)

type Unset struct {}

func init() {
	plugin.Add("Unset", func() plugin.Plugin { return &Unset{} })
}

func (u *Unset) Run(event interface{}) (error, bool) {
	vars := *variables.Load()

	for _, key := range event.([]string) {
		vars.Set(map[string]interface{}{
			key: "",
		})
	}

	return nil, false
}
