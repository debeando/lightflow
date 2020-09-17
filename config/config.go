package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

var config *Structure

func Load() *Structure {
	if config == nil {
		config = &Structure{
			Path: "/etc/lightflow/lightflow.yaml",
		}
	}
	return config
}

func (s *Structure) Read(file_name string) error {
	if len(file_name) == 0 {
		file_name = "lightflow.yaml"
	}

	source, err := ioutil.ReadFile(s.Path)
	if err != nil {
		source, err = ioutil.ReadFile(file_name)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"Fail to read config file: %s, %s or %s",
				file_name,
				s.Path,
				"./lightflow.yaml"),
			)
		}
	}

	source = []byte(os.ExpandEnv(string(source)))

	if err := yaml.Unmarshal(source, &s); err != nil {
		errors.New(fmt.Sprintf("Imposible to parse config file - %s", err))
	}

	return nil
}
