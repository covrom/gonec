package vm

import (
	"fmt"
	"io"

	"github.com/covrom/gonec/gonecparser/ast"
)

// виртуальная машина

type VirtMachine struct {
	af     *ast.File
	w      io.Writer
	scopes map[string]scope
}

func NewVM(af *ast.File, w io.Writer) *VirtMachine {
	return &VirtMachine{
		af:     af,
		w:      w,
		scopes: make(map[string]scope),
	}
}

func (v *VirtMachine) Run() error {
	ast.Inspect(v.af, v.astInspect)
	return nil
}

func (v *VirtMachine) astInspect(n ast.Node) bool {
	var s string
	switch x := n.(type) {
	// TODO: исполнение __init__
	case *ast.GenDecl:

		// s = x.Value

	case *ast.Ident:
		s = x.Name
	}
	if s != "" {
		fmt.Println(s)
		return true
	}
	return false
}
