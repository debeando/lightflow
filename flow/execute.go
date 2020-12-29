package flow

import (
	"errors"
	"fmt"
	"strings"

	"github.com/debeando/lightflow/cli/args"
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/config"
	"github.com/debeando/lightflow/flow/evaluate"
	"github.com/debeando/lightflow/flow/execute"
	"github.com/debeando/lightflow/flow/template"
)

func (f *Flow) Execute() {
	cmd := f.RenderCommand()

	if args.DryRun() {
		fmt.Println(cmd)
	} else {
		f.Retry(func() bool {
			stdout, exit_code := execute.Execute(cmd)

			f.Variables.Set(map[string]interface{}{
				"exit_code": exit_code,
				"stdout":    stdout,
			})

			if exit_code > 0 && f.GetRetryExitCode() == 0 {
				log.Error(f.GetTitle(), map[string]interface{}{
					"Exit Code": exit_code,
				})
			}

			if err := f.ParseStdout(); err != nil {
				log.Error(err.Error(), nil)
			}

			f.Print()
			f.Debug()

			if f.EvalSkip() {
				f.Skip = true
				return false
			}

			f.PrintRetry()

			return f.EvalRetry()
		})
	}
}

func (f *Flow) RenderCommand() string {
	var cmd = f.GetExecute()

	for _, variable := range template.Variables(cmd) {
		if f.Variables.Exist(variable) == false {
			log.Warning(fmt.Sprintf("Register empty variable: %s", variable), nil)
			f.Variables.Set(map[string]interface{}{variable: ""})
		}
	}

	cmd, err := template.Render(cmd, f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(cmd)
}

func (f *Flow) ParseStdout() error {
	switch f.GetFormat() {
	case config.TEXT:
		if reg := f.GetRegister(); len(reg) > 0 {
			f.Variables.Set(map[string]interface{}{reg: f.GetStdOut()})
		}
	case config.JSON:
		//f.Variables puede tener un metodo para salvar en json de forma automatica?
		raw, err := common.StringToJSON(common.InterfaceToString(f.GetStdOut()))
		if err != nil {
			return err
		}

		for variable, value := range raw {
			f.Variables.Set(map[string]interface{}{variable: value})
		}
	default:
		return errors.New("Format option is invalid, please use; TEXT (default) or JSON")
	}

	return nil
}

func (f *Flow) EvalRetry() bool {
	if f.GetRetryWait() == 0 {
		return false
	}

	if f.GetRetryAttempts() == 0 {
		return false
	}

	// EvalRetryByExitCode
	exit_code := f.Variables.Get("exit_code").(int)
	status := common.InterfaceToString(f.Variables.Get(f.GetRetryStatus()))
	error := common.InterfaceToString(f.Variables.Get(f.GetRetryError()))

	if f.GetRetryExitCode() != exit_code {
		return true
	}

	// EvalRetryByStatus
	if exit_code == 0 && len(error) == 0 && len(status) > 0 && len(f.GetRetryDone()) > 0 && f.GetRetryDone() != status {
		return true
	}

	// meter esto en el debug variables
	if len(error) > 0 {
		log.Error(common.InterfaceToString(error), nil)
	}

	return false
}

// EvalSkip evaluate condition to set skip flag.
func (f *Flow) EvalSkip() bool {
	expression, err := template.Render(f.GetSkip(), f.Variables.Items)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return evaluate.Expression(expression)
}

// PrintRetry show the retry progress.
func (f *Flow) PrintRetry() {
	if f.GetRetryAttempts() < 1 {
		return
	}

	log.Info(
		fmt.Sprintf(
			"TASK[%s] SUB TASK[%s] PIPE[%s] RETRY[%d/%d]",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
			f.Attempt,
			f.GetRetryAttempts(),
		), nil)
}

// Print specific variable with value.
func (f *Flow) Print() {
	names := f.GetPrint()
	if names != nil {
		vars := make(map[string]interface{})

		for _, name := range names {
			vars[name] = f.Variables.Get(name)
		}

		log.Info(
			fmt.Sprintf(
				"TASK[%s] SUB TASK[%s] PIPE[%s] PRINT:",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			vars,
		)
	}
}

// Debug print all variables in debug mode.
func (f *Flow) Debug() {
	//f.Variables deberia tener un debug.
	for variable, value := range f.Variables.Items {
		log.Debug(f.GetTitle(), map[string]interface{}{
			variable: value,
		})
	}
}

func (f *Flow) GetExecute() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Execute
}

func (f *Flow) GetFormat() config.Format {
	if len(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format) == 0 {
		return config.TEXT
	}
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Format
}

func (f *Flow) GetRegister() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Register
}

func (f *Flow) GetSkip() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Skip
}

func (f *Flow) GetPrint() []string {
	if len(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Print) == 0 {
		return nil
	}

	return strings.Split(f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].Print, ",")
}
