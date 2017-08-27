// Package json implements json interface for anko script.
package json

import (
	"encoding/json"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("json")
	m.DefineS("Marshal", json.Marshal)
	m.DefineS("Unmarshal", json.Unmarshal)
	return m
}
