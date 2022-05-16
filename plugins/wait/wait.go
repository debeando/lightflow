package retry

import (
	"time"

	"github.com/debeando/lightflow/plugins/plugin"
)

type Wait struct {}

func init() {
	plugin.Add("Wait", func() plugin.Plugin { return &Wait{} })
}

func (w *Wait) Run(event interface{}) (error, uint8) {
 	var wait uint

	switch t := event.(type) {
    case uint:
    	wait = t
    }

	time.Sleep(time.Duration(wait) * time.Second)

	return nil, false
}
