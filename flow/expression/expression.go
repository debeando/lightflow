package expression

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func Render() {

}


func Evaluate(stmt string) interface{} {
	if e, err := parser.ParseExpr(stmt); err == nil {
		return Eval(e)
	}
	return 0
}

func Eval(e ast.Expr) interface{} {
	switch n := e.(type) {
		case *ast.BasicLit:
			// Extrae el valor numerico, alfabetico, símbolos y alfanumerico, menos los bool.
			return n.Value
		case *ast.BinaryExpr:
			x := Eval(n.X)
			y := Eval(n.Y)

			switch n.Op {
				case token.NEQ:
					if x != y {
						return true
					}
				case token.EQL:
					if x == y {
						return true
					}
			}
		case *ast.ParenExpr:
			// Aquí entran todo lo que esta entre paréntesis.
			return Eval(n.X)
	}

	return false
}
