package parser

import (
	"github.com/covrom/gonec/ast"
)

func ConstFolding(inast ast.Stmts) ast.Stmts {

	// num := 4
	// ch := make(chan ast.Stmt, 20)
	// done := make(chan bool, 20)
	// ast.StartStmtSimplifyWorkers(ch, done, num)

	for i := range inast {
		inast[i].Simplify()
		// ch <- inast[i]
	}
	// for i := 0; i < num; i++ {
	// 	done <- true
	// }

	return inast
}
