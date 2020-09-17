package config

// All is a struct to contain all configuration imported or loaded from config file.
type Structure struct {
	Path    string
	General struct {
		Debug               bool   `yaml:"debug"`
		AWSRegion           string `yaml:"aws_region"`
		AWSAccessKeyID      string `yaml:"aws_access_key_id"`
		AWSSecretAccessKey  string `yaml:"aws_secret_access_key"`
		Temporary_Directory string `yaml:"tmp_dir"`
	}
	Tasks []struct {
		Name string  `yaml:"name"`                              // Nombre de la task.
		Loops []map[string]string `yaml:"loops"`                // Lista de ciclos que contiene variables, cada ciclo repite la task.
		Pipes []struct {
			Name string `yaml:"name"`                           // Nombre del pipe.
			Command string `yaml:"command"`                     // Comando a ejecutar. Si hay que limpiar el stdout en formato JSON, usar tool jq.
			Register string `yaml:"register"`                   // Nombre de la variable donde se guarda el stdout/stderr.
			Format string `yaml:"format"`                       // Formato de la variable, si es JSON, un MySQL stdout, CSV, etc... que se anade luego a las variables.
			Variables map[string]interface{} `yaml:"variables"` // Lista de variables.

			While string `yaml:"while"`                         // Condición para salir del reintento, se usan las variables, por eso el format.
			Retry int `yaml:"retry"`                            // Cuantas veces se reintenta el comando.
			Wait int `yaml:"wait"`                              // Cuando tiempo debe transcurrir entre reintento.
		}
	}
}

//	Retry:
// 		While: // Conditions:       (status == "done") si es == true sales.
// 		Wait:
// 		Attempts:
//		Error: "errors" // Variable que indica el error.
