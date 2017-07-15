package vm

import (
	"io"

	"github.com/covrom/gonec/gonecparser/ast"
	"fmt"
)

// виртуальная машина

// в этой мапе ключом является iota поле Data из scope.Object
type varMap map[int]variant

type VirtMachine struct {
	af       *ast.File
	w        io.Writer
	vars     varMap
	funcInit *ast.FuncDecl
}

func NewVM(af *ast.File, w io.Writer) *VirtMachine {
	return &VirtMachine{
		af:   af,
		w:    w,
		vars: make(varMap),
	}
}

func (v *VirtMachine) Run() error {
	ast.Inspect(v.af, v.enumIdents)
	return nil
}

func (v *VirtMachine) enumIdents(n ast.Node) bool {
	switch x := n.(type) {
	case *ast.FuncDecl:
		if x.Name.Name == "__init__" {
			v.funcInit = x
		}
	case *ast.Ident:
		i, ok := x.Obj.Data.(int)
		if ok {
			v.vars[i] = variant{}
			fmt.Printf("Assign %v to id %v",x.Name,i)
		}
	}
	return true
}
