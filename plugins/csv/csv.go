package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Extension string
type Separator string

const (
	dotCSV Extension = ".csv"
	dotTSV           = ".tsv"
	TAB   Separator = "TAB"
	COMMA           = "COMMA"
)

var Extensions = map[string]bool {
    ".csv": true,
    ".tsv": true,
}

type CSV struct {
	Header    bool      `yaml:"header"`
	Path      string    `yaml:"path"` // Path and filename with extension to save results.
	Extension Extension `yaml:"extension"`
	Separator Separator `yaml:"separator"` // comma (default) or tab.
}

func (c *CSV) IsValid() error {
	if len(c.Path) == 0 {
		return errors.New("File name is empty.")
	}

	if len(c.Path) > 0 && len(c.Path) < 5 {
		return errors.New(
			fmt.Sprintf(
				"File name is short: %s",
				c.Path))
	}

	if ! Extensions[string(c.Extension)] {
		return errors.New(
			fmt.Sprintf(
				"File extension is not valid: %s",
				c.Extension))
	}

	if filepath.Ext(c.Path) != string(c.Extension) {
		return errors.New(
			fmt.Sprintf(
				"File extension ins't equal to %s: %s",
				c.Extension,
				c.Path))
	}

	return nil
}

func (c *CSV) Create() error {
	f, err := os.Create(c.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *CSV) Write(chIn <-chan []string) error {
	f, err := os.OpenFile(c.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	if c.Separator == TAB {
		w.Comma = '\t'
	} else {
		w.Comma = ','
	}

	for row := range chIn {
		w.Write(row)
	}

	w.Flush()
	return nil
}
