package plugins

import (
	"github.com/debeando/lightflow/plugins/plugin"

	_ "github.com/debeando/lightflow/plugins/execute"
	_ "github.com/debeando/lightflow/plugins/print"
	_ "github.com/debeando/lightflow/plugins/register"
	_ "github.com/debeando/lightflow/plugins/retry"
	_ "github.com/debeando/lightflow/plugins/wait"
)

func Load(name string, event interface{}) (error, bool) {
	for key := range plugin.Plugins {
		if key == name {
			if creator, ok := plugin.Plugins[key]; ok {
				c := creator()
				return c.Run(event)
			}
		}
	}
	return nil, false
}
