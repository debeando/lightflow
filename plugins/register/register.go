package register

type Format string

const (
	TEXT Format = "TEXT"
	JSON        = "JSON"
)

type Register struct{
	Name   string `yaml:"register"` // Nombre de la variable donde se guarda el stdout/stderr solo cuando el formato es TEXT, se usa para guardar un valor de un pipe y usarlo en otro pipe.
	Format Format `yaml:"format"`   // Formato de la variable, por defecto TEXT, si es JSON, un MySQL stdout, CSV, etc... que se anade luego a las variables.
}
