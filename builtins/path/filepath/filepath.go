// Package path implements path manipulation interface for anko script.
package filepath

import (
	f "path/filepath"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("filepath")
	m.DefineS("Join", f.Join)
	m.DefineS("Clean", f.Join)
	m.DefineS("Abs", f.Abs)
	m.DefineS("Base", f.Base)
	m.DefineS("Clean", f.Clean)
	m.DefineS("Dir", f.Dir)
	m.DefineS("EvalSymlinks", f.EvalSymlinks)
	m.DefineS("Ext", f.Ext)
	m.DefineS("FromSlash", f.FromSlash)
	m.DefineS("Glob", f.Glob)
	m.DefineS("HasPrefix", f.HasPrefix)
	m.DefineS("IsAbs", f.IsAbs)
	m.DefineS("Join", f.Join)
	m.DefineS("Match", f.Match)
	m.DefineS("Rel", f.Rel)
	m.DefineS("Split", f.Split)
	m.DefineS("SplitList", f.SplitList)
	m.DefineS("ToSlash", f.ToSlash)
	m.DefineS("VolumeName", f.VolumeName)
	return m
}
