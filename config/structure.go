package config

import (
	"github.com/debeando/lightflow/plugins/aws"
	// "github.com/debeando/lightflow/plugins/chunk"
	"github.com/debeando/lightflow/plugins/csv"
	"github.com/debeando/lightflow/plugins/evaluate"
	"github.com/debeando/lightflow/plugins/execute"
	"github.com/debeando/lightflow/plugins/mysql"
	"github.com/debeando/lightflow/plugins/register"
	// "github.com/debeando/lightflow/plugins/retry"
	"github.com/debeando/lightflow/plugins/slack"
)

// All is a struct to contain all configuration imported or loaded from config file.
type Structure struct {
	Path string
	Debug bool   `yaml:"debug"`
	Variables map[string]interface{} `yaml:"variables"` // Lista global de variables para usar en los pipes.
	TasksInclude []string `yaml:"tasks_include"`        // Path list of subtasks.
	Tasks        []Task   `yaml:"tasks"`                // Subtask is a common collection of pipes.
	PipesInclude []string `yaml:"pipes_include"`        // Path list of pipes.
	Pipes        []Pipe   `yaml:"pipes"`                // Pipes is a common collection for subtask.
}

type Task struct {
	Name      string                 `yaml:"name"`
	Ignore    bool                   `yaml:"ignore"`
	Variables map[string]interface{} `yaml:"variables"` // Lista de variables para usar en los pipes.
}

// The elemnts into pipe struct define the order execution.
type Pipe struct { // Pipes is a common collection for subtask.
	Name      string                 `yaml:"name"`      // Pipe name.
	Ignore    bool                   `yaml:"ignore"`    // Ignore this pipe.
	Unset     []string               `yaml:"unset"`     // List of variables to unset every pipe loop.
	// When      string                 `yaml:"when"`      // Evaluate expression before execute pipe, require true to run.
	Evaluate  []evaluate.Evaluate
	Execute   execute.Execute
	Register  register.Register
	Print     []string               `yaml:"print"`     // List of variables to print.
	Wait      uint                   `yaml:"wait"`      // Sleep for N seconds pipe before start.
	// Variables map[string]interface{} `yaml:"variables"` // Lista de variables.
	// Skip      string                 `yaml:"skip"`      // Skip own pipe block when specific expression condition, use the variable definied in the Register to compare. First run pipe and them evaluate skip condition.
	Error     string                 `yaml:"error"`     // Show error when specific expression condition, use the variable definied in the Register to compare. By default is exit_code != 0.
	// Retry     retry.Retry
	// Chunk     chunk.Chunk
	Slack     slack.Slack
	MySQL     mysql.MySQL
	CSV       csv.CSV
	AWS       aws.AWS
}
