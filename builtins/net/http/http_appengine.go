// +build appengine

// Package net implements http interface for anko script.
package net

import (
	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	panic("can't import 'net/http'")
	return nil
}
