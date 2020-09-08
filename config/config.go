package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

var File Config

// Init does any initialization necessary for the module.
func init() {
	File = Config{
		Path: "/etc/lightflow/lightflow.yaml",
	}
}

func (c *Config) Load(file_name string) error {
	if len(file_name) == 0 {
		file_name = "lightflow.yaml"
	}

	source, err := ioutil.ReadFile(c.Path)
	if err != nil {
		source, err = ioutil.ReadFile(file_name)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"Fail to read config file: %s, %s or %s",
				file_name,
				c.Path,
				"./lightflow.yaml"),
			)
		}
	}

	source = []byte(os.ExpandEnv(string(source)))

	if err := yaml.Unmarshal(source, &c); err != nil {
		errors.New(fmt.Sprintf("Imposible to parse config file - %s", err))
	}

	return nil
}
