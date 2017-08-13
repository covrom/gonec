// Package rand implements math/rand interface for anko script.
package rand

import (
	t "math/rand"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("rand")
	m.DefineS("ExpFloat64", t.ExpFloat64)
	m.DefineS("Float32", t.Float32)
	m.DefineS("Float64", t.Float64)
	m.DefineS("Int", t.Int)
	m.DefineS("Int31", t.Int31)
	m.DefineS("Int31n", t.Int31n)
	m.DefineS("Int63", t.Int63)
	m.DefineS("Int63n", t.Int63n)
	m.DefineS("Intn", t.Intn)
	m.DefineS("NormFloat64", t.NormFloat64)
	m.DefineS("Perm", t.Perm)
	m.DefineS("Seed", t.Seed)
	m.DefineS("Uint32", t.Uint32)
	return m
}
