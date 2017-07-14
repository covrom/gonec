package vm

import (
	"fmt"
	"io"

	"github.com/covrom/gonec/gonecparser/ast"
)

type VirtMachine struct {
	af *ast.File
	w  io.Writer
}

func NewVM(af *ast.File, w io.Writer) *VirtMachine {
	return &VirtMachine{af: af, w: w}
}

func (v *VirtMachine) Run() error {
	ast.Inspect(v.af, v.astInspect)
	return nil
}

func (v *VirtMachine) astInspect(n ast.Node) bool {
	var s string
	switch x := n.(type) {
	case *ast.BasicLit:
		s = x.Value
	case *ast.Ident:
		s = x.Name
	}
	if s != "" {
		fmt.Println(s)
	}
	return true
}
