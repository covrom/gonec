// +build !appengine

// Package url implements url interface for anko script.
package url

import (
	u "net/url"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("url")
	m.DefineTypeS("Values", make(u.Values))
	m.DefineS("Parse", u.Parse)
	return m
}
