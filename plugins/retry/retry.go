package retry

import (
	"time"
	"errors"

	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/variables"
)

// Retry save the common variables:
type Retry struct {
	Attempt    uint   `yaml:"attempts"`   // Current attempt
	Wait       uint   `yaml:"wait"`       // Wait between attempt
	Expression string `yaml:"expression"` // Expression to evaluate condition and retry.
}

func init() {
	plugin.Add("Retry", func() plugin.Plugin { return &Retry{} })
}

func (r *Retry) Run(event interface{}) (error, uint8) {
	retry, ok := event.(Retry)
	if !ok {
		return errors.New("Invalid struct"), false
	}

	if retry.Attempt == 0 {
		return errors.New("Retry attempt should be greater equal than 1."), false
	}

	vars := *variables.Load()

	attempt_tmp := vars.Get("retry_attempt")
	switch t := attempt_tmp.(type) {
    case uint:
    	retry.Attempt = t
    }
	retry.Attempt = retry.Attempt - 1

	vars.Set(map[string]interface{}{
		"retry_attempt": retry.Attempt,
	})

	time.Sleep(time.Duration(retry.Wait) * time.Second)

	if retry.Attempt == 0 {
		return nil, false
	}

	return nil, true
}
