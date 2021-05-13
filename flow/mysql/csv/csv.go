package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type CSV struct {
	Path string
	Extension string
	Separator rune
}

func (c *CSV) IsValidPath() error {
	if len(c.Path) == 0 {
		return errors.New("File name is empty.")
	}

	if len(c.Path) > 0 && len(c.Path) < 5 {
		return errors.New(fmt.Sprintf("File name is short: %s", c.Path))
	}

	if filepath.Ext(c.Path) != c.Extension {
		return errors.New(fmt.Sprintf("File extension ins't equal to %s: %s", c.Extension, c.Path))
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
	w.Comma = c.Separator

	for row := range chIn {
		w.Write(row)
	}

	w.Flush()
	return nil
}
