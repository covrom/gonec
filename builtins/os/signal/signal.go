// Package signal implements signal interface for anko script.
package signal

import (
	pkg "os/signal"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("os/signal")

	//m.DefineS("Ignore", pkg.Ignore)
	m.DefineS("Notify", pkg.Notify)
	//m.DefineS("Reset", pkg.Reset)
	m.DefineS("Stop", pkg.Stop)
	return m
}
