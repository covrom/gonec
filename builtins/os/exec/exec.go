// Package exec implements os/exec interface for anko script.
package exec

import (
	e "os/exec"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("exec")
	m.DefineS("ErrNotFound", e.ErrNotFound)
	m.DefineS("LookPath", e.LookPath)
	m.DefineS("Command", e.Command)
	return m
}
