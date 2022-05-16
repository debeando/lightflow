package core

import (
	"github.com/debeando/lightflow/plugins"
)

func (core *Core) aws() {
	aws := plugins.PluginAWS{
		Config: core.Config.Pipes[core.Index.Pipe].AWS,
	}
	aws.Load()
}
