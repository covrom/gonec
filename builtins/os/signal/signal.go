// Package signal implements signal interface for anko script.
package signal

import (
	pkg "os/signal"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("os/signal")

	//m.DefineS("Ignore", pkg.Ignore)
	m.DefineS("Notify", pkg.Notify)
	//m.DefineS("Reset", pkg.Reset)
	m.DefineS("Stop", pkg.Stop)
	return m
}
