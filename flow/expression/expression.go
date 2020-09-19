package expression

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/swapbyt3s/lightflow/variables"
)

func Evaluate(stmt string) bool {
	if e, err := parser.ParseExpr(stmt); err == nil {
		return Eval(e).(bool)
	}
	return false
}

func Eval(e ast.Expr) interface{} {
	switch n := e.(type) {
	case *ast.Ident:
		// Detecta una variable, nosotros la definimos:
		return Eval(&ast.BasicLit {
			ValuePos: n.NamePos,
			Kind: token.STRING,
			Value: fmt.Sprintf("\"%v\"", variables.Load().Get(n.Name)),
		})
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
