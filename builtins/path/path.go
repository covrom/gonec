// Package path implements path interface for anko script.
package path

import (
	pkg "path"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("path")
	m.DefineS("Base", pkg.Base)
	m.DefineS("Clean", pkg.Clean)
	m.DefineS("Dir", pkg.Dir)
	m.DefineS("ErrBadPattern", pkg.ErrBadPattern)
	m.DefineS("Ext", pkg.Ext)
	m.DefineS("IsAbs", pkg.IsAbs)
	m.DefineS("Join", pkg.Join)
	m.DefineS("Match", pkg.Match)
	m.DefineS("Split", pkg.Split)
	return m
}
