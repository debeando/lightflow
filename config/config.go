package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"reflect"

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

	s.ReadInclude()

	return s.IsValid()
}

func (s *Structure) ReadInclude() {
	for _, inc := range s.TasksInclude {
		source, err := ioutil.ReadFile(inc)
		if err != nil {
			errors.New(err.Error())
		}

		var task Task

		source = []byte(os.ExpandEnv(string(source)))

		if err := yaml.Unmarshal(source, &task); err != nil {
			errors.New(
				fmt.Sprintf(
					"Imposible to parse config file: %s",
					inc,
				))
		}

		if ! reflect.DeepEqual(task, Task{}) {
			s.Tasks = append(s.Tasks, task)
		}
	}

	for _, inc := range s.PipesInclude {
		source, err := ioutil.ReadFile(inc)
		if err != nil {
			errors.New(err.Error())
		}

		var pipe Pipe

		source = []byte(os.ExpandEnv(string(source)))

		if err := yaml.Unmarshal(source, &pipe); err != nil {
			errors.New(
				fmt.Sprintf(
					"Imposible to parse config file: %s",
					inc,
				))
		}

		if ! reflect.DeepEqual(pipe, Pipe{}) {
			s.Pipes = append(s.Pipes, pipe)
		}
	}
}

func (s *Structure) IsValid() error {
	var re = regexp.MustCompile(`^[0-9A-Za-z\-\_]+$`)

	for task_index := range s.Tasks {
		if !re.MatchString(s.Tasks[task_index].Name) {
			return errors.New(
				fmt.Sprintf(
					"Invalid task name for '%s', only allow 0-9, A-Z, a-z, - and _.",
					s.Tasks[task_index].Name,
				))
		}
	}

	for pipe_index := range s.Pipes {
		if !re.MatchString(s.Pipes[pipe_index].Name) {
			return errors.New(
				fmt.Sprintf(
					"Invalid pipe name for '%s', only allow 0-9, A-Z, a-z, - and _.",
					s.Pipes[pipe_index].Name,
				))
		}
	}

	return nil
}
