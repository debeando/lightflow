package print

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/variables"
)

type Print struct{}

func init() {
	plugin.Add("Print", func() plugin.Plugin { return &Print{} })
}

func (p *Print) Run(event interface{}) (error, bool) {
	vals := make(map[string]interface{})
	vars := *variables.Load()

 	for _, key := range event.([]string) {
 		if vars.Get(key) != nil {
 			vals[key] = vars.Get(key)
 		} else {
 			log.Info(fmt.Sprintf("Print %s", key), nil)
 		}
 	}

 	if len(vals) > 0 {
		log.Info(
			"Print",
			vals,
		)
 	}

	return nil, false
}
