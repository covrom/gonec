package parser

import (
	"github.com/covrom/gonec/ast"
)

func ConstFolding(inast ast.Stmts) ast.Stmts {
	for i := range inast {
		inast[i].Simplify()
	}
	return inast
}

