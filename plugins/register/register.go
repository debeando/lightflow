package register

import (
	// "fmt"
	"errors"

	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/variables"
)

type Format string

const (
	TEXT Format = "TEXT"
	JSON        = "JSON"
	INT         = "INT"
)

type Register struct{
	// Nombre de la variable donde se guarda el stdout/stderr solo cuando el
	// formato es TEXT, se usa para guardar un valor de un pipe y usarlo en
	// otro pipe.
	Name   string `yaml:"name"`
	// Formato de la variable, por defecto TEXT, si es JSON, un MySQL stdout,
	// CSV, etc... que se anade luego a las variables.
	Format Format `yaml:"format"`
}

func init() {
	plugin.Add("Register", func() plugin.Plugin { return &Register{} })
}

func (r *Register) Run(event interface{}) (error, uint8) {
	register, ok := event.(Register)
	if !ok {
		return errors.New("Invalid struct"), false
	}

	vars := *variables.Load()
	vars.Set(map[string]interface{}{
		register.Name: vars.Get("stdout"),
	})

	return nil, false
}
