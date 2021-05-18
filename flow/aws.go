package flow

import (
	"github.com/debeando/lightflow/plugins"
)

func (f *Flow) aws() {
	aws := plugins.PluginAWS {
		Config: f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].AWS,
	}
	aws.Load()
}
