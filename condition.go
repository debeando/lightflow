package main

import (
	"fmt"
	"strconv"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	var stmt string

	stmt = "1 != 2"

	if e, err := parser.ParseExpr(stmt); err == nil {
		eval(e)
	}
}

func eval(e ast.Expr) int {
  switch n := e.(type) {
    case *ast.BasicLit:
        if n.Kind != token.INT {
		fmt.Println(n.Kind)
        }
        i, _ := strconv.Atoi(n.Value)
        return i

    case *ast.BinaryExpr:
      fmt.Println("BinaryExpr")
      fmt.Println(n.Op)
      fmt.Println(n.X)
      fmt.Println(n.Y)

      x := eval(n.X)
      y := eval(n.Y)

      switch n.Op {
        case token.NEQ:
            fmt.Println(x != y)

      }
  }

  return 0
}
