package vm

import (
	"errors"
	"io"

	"fmt"

	"github.com/covrom/gonec/gonecparser/ast"
	"github.com/covrom/gonec/gonecparser/variant"
	"github.com/covrom/gonec/gonecparser/token"
)

var (
	InitError = errors.New("Отсутствует исполняемый код модуля")
)

// виртуальная машина

type VirtMachine struct {
	af          *ast.File
	w           io.Writer
	fset        *token.FileSet
	mainContext *variant.Context
}

func NewVM(af *ast.File, w io.Writer, fst *token.FileSet) *VirtMachine {
	return &VirtMachine{
		af:          af,
		w:           w,
		fset:        fst,
		mainContext: nil,
	}
}

func (v *VirtMachine) Run() error {
	for _, ur := range v.af.Unresolved {
		fmt.Printf("Не назначен объект у идентификатора %v\n", ur.Name)
	}
	v.mainContext = variant.NewMainContext()
	defer v.mainContext.Destroy()

	ini := v.af.Scope.Lookup("__init__")
	if ini == nil {
		return InitError
	}
	
	ast.Inspect(ini.Decl.(*ast.FuncDecl).Body, v.walker)
	return nil
}

func (v *VirtMachine) walker(n ast.Node) bool {

	// ast.Print(v.fset,n)

	switch x := n.(type) {
	case *ast.BlockStmt:
	case *ast.AssignStmt:
		ast.Print(v.fset,x.Lhs)
		ast.Print(v.fset,x.Rhs)
		
	}
	return true
}
