package core

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins"
)

func (core *Core) mysql() {
	var chData = make(chan []string)

	defer close(chData)

	c := plugins.PluginCSV{
		Config: core.Config.Pipes[core.Index.Pipe].CSV,
	}
	writeToCSV, err := c.Load()
	if err != nil {
		log.Error(
			fmt.Sprintf(
				"%s/%s %s",
				core.TaskName(),
				core.PipeName(),
				err,
			),
			nil,
		)
	}

	if writeToCSV {
		go func() {
			err := c.Write(chData)
			if err != nil {
				log.Error(
					fmt.Sprintf(
						"%s/%s %s",
						core.TaskName(),
						core.PipeName(),
						err,
					),
					nil,
				)
			}
		}()
	}

	p := plugins.PluginMySQL{
		Config: core.Config.Pipes[core.Index.Pipe].MySQL,
	}

	p.Retrieve(func(rowCount int, columns []string, row []string) bool {
		if writeToCSV {
			if rowCount == 1 && c.Config.Header {
				chData <- columns
			}
			chData <- row
		}
		return false
	})
}
