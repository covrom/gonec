// Package errors implements errors interface for anko script.
package errors

import (
	pkg "errors"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewModule("errors")
	m.DefineS("New", pkg.New)
	return m
}
