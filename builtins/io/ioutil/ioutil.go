// Package ioutil implements I/O interface for anko script.
package ioutil

import (
	u "io/ioutil"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("iotuil")
	m.DefineS("ReadAll", u.ReadAll)
	m.DefineS("ReadDir", u.ReadDir)
	m.DefineS("ReadFile", u.ReadFile)
	m.DefineS("WriteFile", u.WriteFile)
	return m
}
