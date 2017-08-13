// Package json implements json interface for anko script.
package json

import (
	"encoding/json"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("json")
	m.DefineS("Marshal", json.Marshal)
	m.DefineS("Unmarshal", json.Unmarshal)
	return m
}
