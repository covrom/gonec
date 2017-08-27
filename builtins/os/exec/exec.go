// Package exec implements os/exec interface for anko script.
package exec

import (
	e "os/exec"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("exec")
	m.DefineS("ErrNotFound", e.ErrNotFound)
	m.DefineS("LookPath", e.LookPath)
	m.DefineS("Command", e.Command)
	return m
}
