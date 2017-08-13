// Package errors implements errors interface for anko script.
package errors

import (
	pkg "errors"
	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewModule("errors")
	m.DefineS("New", pkg.New)
	return m
}
