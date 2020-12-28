package evaluate

import (
	"go/constant"
	"go/token"
	"go/types"
)

// Expression allow evaluate formulas and return true or false value if the result is satisfactory.
func Expression(formula string) bool {
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, formula)
	if err != nil {
		return false
	}
	return constant.BoolVal(tv.Value)
}
