package config

// All is a struct to contain all configuration imported or loaded from config file.
type Config struct {
	Path    string
	General struct {
		Debug               bool   `yaml:"debug"`
		AWSRegion           string `yaml:"aws_region"`
		AWSAccessKeyID      string `yaml:"aws_access_key_id"`
		AWSSecretAccessKey  string `yaml:"aws_secret_access_key"`
		Temporary_Directory string `yaml:"tmp_dir"`
	}
	Tasks []struct {
		Name string  `yaml:"name"`                         // Nombre de la task.
		Loops []map[string]string `yaml:"loops"`           // Lista de ciclos que contiene variables, cada ciclo repite la task.
		Pipes []struct {
			Name string `yaml:"name"`                      // Nombre del pipe.
			Command string `yaml:"command"`                // Comando a ejecutar.
			Register string `yaml:"register"`              // Nombre de la variable donde se guarda el stdout/stderr.
			Format string `yaml:"format"`                  // Formato de la variable, si es JSON, un MySQL stdout, CSV, etc... que se anade luego a las variables.
			While string `yaml:"while"`                    // Condición para salir del reintento, se usan las variables, por eso el format.
			Retry string `yaml:"retry"`                    // Cuantas veces se reintenta el comando.
			Wait string `yaml:"wait"`                      // Cuando tiempo debe transcurrir entre reintento.
			Variables map[string]string `yaml:"variables"` // Lista de variables.
		}
	}
}
