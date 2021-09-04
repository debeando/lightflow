package retry

import (
	"time"

	"github.com/debeando/lightflow/plugins/plugin"
)

type Wait struct {}

func init() {
	plugin.Add("Wait", func() plugin.Plugin { return &Wait{} })
}

func (w *Wait) Run(event interface{}) (error, bool) {
 	var wait uint

	switch t := event.(type) {
    case uint:
    	wait = t
    }

	time.Sleep(time.Duration(wait) * time.Second)

	return nil, false
}
