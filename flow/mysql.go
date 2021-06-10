package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins"
)

func (f *Flow) mysql() {
	var chData = make(chan []string)

	defer close(chData)

	c := plugins.PluginCSV{
		Config: f.Config.Pipes[f.Index.Pipe].CSV,
	}
	writeToCSV, err := c.Load()
	if err != nil {
		log.Error(
			fmt.Sprintf(
				"%s/%s %s",
				f.TaskName(),
				f.PipeName(),
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
						f.TaskName(),
						f.PipeName(),
						err,
					),
					nil,
				)
			}
		}()
	}

	p := plugins.PluginMySQL{
		Config: f.Config.Pipes[f.Index.Pipe].MySQL,
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
