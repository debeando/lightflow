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
		Slack               struct {
			Token string `yaml:"token"`
		}
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
			Name      string                 `yaml:"name"`      // Pipe name.
			Ignore    bool                   `yaml:"ignore"`    // Ignore this pipe.
			When      string                 `yaml:"when"`      // Evaluate expression before execute pipe, require true to run.
			Execute   string                 `yaml:"execute"`   // Comando a ejecutar. Si hay que limpiar el stdout en formato JSON, usar tool jq.
			Register  string                 `yaml:"register"`  // Nombre de la variable donde se guarda el stdout/stderr solo cuando el formato es TEXT, se usa para guardar un valor de un pipe y usarlo en otro pipe.
			Format    Format                 `yaml:"format"`    // Formato de la variable, por defecto TEXT, si es JSON, un MySQL stdout, CSV, etc... que se anade luego a las variables.
			Print     []string               `yaml:"print"`     // List of variables to print.
			Unset     []string               `yaml:"unset"`     // List of variables to unset every pipe loop.
			Variables map[string]interface{} `yaml:"variables"` // Lista de variables.
			Wait      uint64                 `yaml:"wait"`      // Sleep for N seconds pipe before start.
			Skip      string                 `yaml:"skip"`      // Skip own pipe block when specific expression condition, use the variable definied in the Register to compare. First run pipe and them evaluate skip condition.
			Error     string                 `yaml:"error"`     // Show error when specific expression condition, use the variable definied in the Register to compare. By default is exit_code != 0.
			Retry     struct {               // Retry execution command when it fail, retry found inside Chunk.
				Attempts   int    `yaml:"attempts"`   // Cuantas veces se reintenta el comando.
				Wait       int    `yaml:"wait"`       // Cuando tiempo debe transcurrir entre reintentos.
				Expression string `yaml:"expression"` // Expression to evaluate condition and retry.
			}
			Chunk struct { // Loop own command by chunk logic.
				Limit int `yaml:"limit"` // Número máximo de elementos por chunk.
				Total int `yaml:"total"` // Número total de elementos.
			}
			Slack struct { // Send message to slack
				Channel    string `yaml:"channel"`
				Color      string `yaml:"color"`      // Can either be one of good (green), warning (yellow), danger (red), or any hex color code (eg. #439FE0).
				Expression string `yaml:"expression"` // Expression to evaluate condition and send message.
				Message    string `yaml:"message"`
				Title      string `yaml:"title"`
			}
			MySQL struct { // Connect to MySQL server.
				Host       string `yaml:"host"`
				Port       int    `yaml:"port"`
				User       string `yaml:"user"`
				Password   string `yaml:"password"`
				Schema     string `yaml:"schema"`
				Query      string `yaml:"query"`
				Header   []string `yaml:"header"`
				Path       string `yaml:"path"` // Path and filename to save result into file.
			}
		}
	}
}
