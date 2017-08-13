// Package fmt implements json interface for anko script.
package fmt

import (
	pkg "fmt"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("fmt")
	m.DefineS("Errorf", pkg.Errorf)
	m.DefineS("Fprint", pkg.Fprint)
	m.DefineS("Fprintf", pkg.Fprintf)
	m.DefineS("Fprintln", pkg.Fprintln)
	m.DefineS("Fscan", pkg.Fscan)
	m.DefineS("Fscanf", pkg.Fscanf)
	m.DefineS("Fscanln", pkg.Fscanln)
	m.DefineS("Print", pkg.Print)
	m.DefineS("Printf", pkg.Printf)
	m.DefineS("Println", pkg.Println)
	m.DefineS("Scan", pkg.Scan)
	m.DefineS("Scanf", pkg.Scanf)
	m.DefineS("Scanln", pkg.Scanln)
	m.DefineS("Sprint", pkg.Sprint)
	m.DefineS("Sprintf", pkg.Sprintf)
	m.DefineS("Sprintln", pkg.Sprintln)
	m.DefineS("Sscan", pkg.Sscan)
	m.DefineS("Sscanf", pkg.Sscanf)
	m.DefineS("Sscanln", pkg.Sscanln)
	return m
}
