package flow

import (
	"github.com/debeando/lightflow/plugins"
)

func (f *Flow) aws() {
	aws := plugins.PluginAWS{
		Config: f.Config.Pipes[f.Index.Pipe].AWS,
	}
	aws.Load()
}
