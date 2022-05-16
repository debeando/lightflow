package evaluate

import (
	"fmt"
	"errors"

	"go/constant"
	"go/token"
	"go/types"

	"github.com/debeando/lightflow/plugins/plugin"
	"github.com/debeando/lightflow/common/template"
	"github.com/debeando/lightflow/variables"
)

type Action string

const (
	None Action = "None"  // No action.
	Skip        = "Skip"  // Skip curren execution pipe.
	When        = "When"  // Evaluate expression before execute pipe.
	Retry       = "Retry" // Retry current pipe until satify expression.
)

type Evaluate struct{
	Expression string // 
	Action     Action // Default: None. Options: None, Skip, When, Retry
}

func init() {
	plugin.Add("Evaluate", func() plugin.Plugin { return &Evaluate{} })
}

func (r *Evaluate) Run(event interface{}) (error, bool) {
	var err error
	vars := *variables.Load()

	items, ok := event.([]Evaluate)
	if !ok {
		return errors.New("Invalid struct"), false
	}

	for _, item := range items {
		item.Expression, err = template.Render(item.Expression, vars.GetItems())
		if err != nil {
			return err, true
		}

		if Expression(item.Expression) {
			fmt.Println(item.Expression, "Si!...")
		}
	}

	return nil, false
}

// Expression allow evaluate formulas and return true or false value if the
// result is satisfactory.
func Expression(formula string) bool {
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, formula)
	if err != nil {
		return false
	}
	return constant.BoolVal(tv.Value)
}
