// +build appengine

// Package url implements url interface for anko script.
package url

import (
	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	panic("can't import 'url'")
	return nil
}
