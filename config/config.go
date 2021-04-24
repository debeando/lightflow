package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

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
	// estas son las cosas que usando un metodo que valida y retorna el valor default es más elegante.
	if len(file_name) == 0 {
		file_name = "lightflow.yaml"
	}

	// creo que estos dos ReadFile anidados se pueden optimizar
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
		return errors.New(fmt.Sprintf(
			"Imposible to parse config file: %s",
			err,
			),
		)
	}

	return s.Validate()
}

func (s *Structure) Validate() error {
	var re = regexp.MustCompile(`^[0-9A-Za-z\-\_]+$`)

	for task_index := range s.Tasks {
		if !re.MatchString(s.Tasks[task_index].Name) {
			return errors.New(
				fmt.Sprintf(
					"Invalid task name for '%s', only allow 0-9, A-Z, a-z, - and _.",
					s.Tasks[task_index].Name,
				))
		}

		for subtask_index := range s.Tasks[task_index].Subtasks {
			if !re.MatchString(s.Tasks[task_index].Subtasks[subtask_index].Name) {
				return errors.New(
					fmt.Sprintf(
						"Invalid sub task name for '%s', only allow 0-9, A-Z, a-z, - and _.",
						s.Tasks[task_index].Subtasks[subtask_index].Name,
					))
			}
		}

		for pipe_index := range s.Tasks[task_index].Pipes {
			if !re.MatchString(s.Tasks[task_index].Pipes[pipe_index].Name) {
				return errors.New(
					fmt.Sprintf(
						"Invalid pipe name for '%s', only allow 0-9, A-Z, a-z, - and _.",
						s.Tasks[task_index].Pipes[pipe_index].Name,
					))
			}
		}
	}

	return nil
}
