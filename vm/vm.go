package vm

import (
	"io"

	"fmt"

	"github.com/covrom/gonec/gonecparser/ast"
	"github.com/covrom/gonec/gonecparser/variant"
)

// виртуальная машина

type VirtMachine struct {
	af       *ast.File
	w        io.Writer
	ctx	variant.Context
}

func NewVM(af *ast.File, w io.Writer) *VirtMachine {
	return &VirtMachine{
		af: af,
		w:  w,
	}
}

func (v *VirtMachine) Run() error {
	for _,ur:=range v.af.Unresolved{
		fmt.Printf("Не назначен объект у идентификатора %v\n", ur.Name)
	}

	ast.Inspect(v.af, v.enumIdents)
	return nil
}

func (v *VirtMachine) enumIdents(n ast.Node) bool {
	switch x := n.(type) {
	case *ast.FuncDecl:
		if x.Name.Name == "__init__" {
			v.funcInit = x
			//добавляем контекст окружения в стек вызовов

		}
	case *ast.Ident:
		//экспортируемый идентификатор определяется по суффиксу "экспорт"
		if x.Obj != nil {
			if x.Obj.Var == nil {
				x.Obj.Var = variant.NewVariant()
				fmt.Printf("Resolved, assign new variant to %v in scope %v\n", x.Name, x.Scope)
			}
		} 
	}
	return true
}
