// +build appengine

// Package net implements net interface for anko script.
package net

import (
	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	panic("can't import 'net'")
	return nil
}
