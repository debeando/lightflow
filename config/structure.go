package config

type Format string

const (
	TEXT Format = "TEXT"
	JSON        = "JSON"
)

// All is a struct to contain all configuration imported or loaded from config file.
type Structure struct {
	Path    string
	General struct {
		Debug               bool   `yaml:"debug"`
		Temporary_Directory string `yaml:"tmp_dir"`
	}
	Variables map[string]interface{} `yaml:"variables"` // Lista global de variables para usar en los pipes.
	Tasks     []struct {
		Name    string     `yaml:"name"` // Nombre de la task.
		Subtask []struct { // Subtask
			Name      string                 `yaml:"name"`      // Nombre del loop, cada loop ejecuta el grupo de pipes.
			Ignore    bool                   `yaml:"ignore"`    // Ignore this subtask.
			Variables map[string]interface{} `yaml:"variables"` // Lista de variables para usar en los pipes.
		}
		Pipes []struct { // Pipes is a common collection for subtask.
			Name      string                 `yaml:"name"`      // Nombre del pipe.
			Ignore    bool                   `yaml:"ignore"`    // Ignore this pipe.
			Execute   string                 `yaml:"execute"`   // Comando a ejecutar. Si hay que limpiar el stdout en formato JSON, usar tool jq.
			Register  string                 `yaml:"register"`  // Nombre de la variable donde se guarda el stdout/stderr solo cuando el formato es TEXT, se usa para guardar un valor de un pipe y usarlo en otro pipe.
			Format    Format                 `yaml:"format"`    // Formato de la variable, por defecto TEXT, si es JSON, un MySQL stdout, CSV, etc... que se anade luego a las variables.
			Print     string                 `yaml:"print"`     // Print specific variables, each variables are comma separated.
			Variables map[string]interface{} `yaml:"variables"` // Lista de variables.
			Skip      struct {               // Skip own pipe block when specific condition, use the variable definied in the Register to compare, by default use stdout variable to compare. And only compare string value.
				Variable string `yaml:"variable"` // Variable name to take value and compare.
				Equals   string `yaml:"equals"`   // Condition to evaluate equal value registered to allow skip.
			}
			Retry struct { // Retry execution command when it fail, retry found inside Chunk.
				Status   string `yaml:"status"`    //
				Done     string `yaml:"done"`      //
				Attempts int    `yaml:"attempts"`  // Cuantas veces se reintenta el comando.
				Wait     int    `yaml:"wait"`      // Cuando tiempo debe transcurrir entre reintentos.
				Error    string `yaml:"error"`     // Variable que indica la existencia de un error, incluso se usa para volver hacer el reintento.
				ExitCode int    `yaml:"exit_code"` // Reintenta hasta que el exit code sea diferente al definido.
			}
			Chunk struct { // Loop own command by chunk logic.
				Limit int `yaml:"limit"` // Número máximo de elementos por chunk.
				Total int `yaml:"total"` // Número total de elementos.
			}
		}
	}
}
