package plugins

import (
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/csv"
	"github.com/debeando/lightflow/plugins/template"
	"github.com/debeando/lightflow/variables"
)

type PluginCSV struct {
	Config    csv.CSV
	Variables variables.List
}

func (p *PluginCSV) Load() (bool, error) {
	p.Variables = *variables.Load()
	p.Config.Path = p.Render(p.Config.Path)

	if err := p.Config.IsValid(); err != nil {
		if err.Error() == "File name is empty." {
			return false, nil
		}
		return false, err
	} else if err := p.Config.Create(); err != nil {
		return false, err
	}

	return true, nil
}

func (p *PluginCSV) Write(chIn <-chan []string) error {
	err := p.Config.Write(chIn)
	if err != nil {
		return err
	}

	return nil
}

func (p *PluginCSV) Render(s string) string {
	r, err := template.Render(s, p.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
