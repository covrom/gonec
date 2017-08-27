// Package regexp implements regexp interface for anko script.
package sort

import (
	r "regexp"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("sort")
	m.DefineS("Match", r.Match)
	m.DefineS("MatchReader", r.MatchReader)
	m.DefineS("MatchString", r.MatchString)
	m.DefineS("QuoteMeta", r.QuoteMeta)
	m.DefineS("Compile", r.Compile)
	m.DefineS("CompilePOSIX", r.CompilePOSIX)
	m.DefineS("MustCompile", r.MustCompile)
	m.DefineS("MustCompilePOSIX", r.MustCompilePOSIX)
	return m
}
