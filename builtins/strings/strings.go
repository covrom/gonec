// Package strings implements strings interface for anko script.
package strings

import (
	pkg "strings"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("strings")
	m.DefineS("Contains", pkg.Contains)
	m.DefineS("ContainsAny", pkg.ContainsAny)
	m.DefineS("ContainsRune", pkg.ContainsRune)
	m.DefineS("Count", pkg.Count)
	m.DefineS("EqualFold", pkg.EqualFold)
	m.DefineS("Fields", pkg.Fields)
	m.DefineS("FieldsFunc", pkg.FieldsFunc)
	m.DefineS("HasPrefix", pkg.HasPrefix)
	m.DefineS("HasSuffix", pkg.HasSuffix)
	m.DefineS("Index", pkg.Index)
	m.DefineS("IndexAny", pkg.IndexAny)
	m.DefineS("IndexByte", pkg.IndexByte)
	m.DefineS("IndexFunc", pkg.IndexFunc)
	m.DefineS("IndexRune", pkg.IndexRune)
	m.DefineS("Join", pkg.Join)
	m.DefineS("LastIndex", pkg.LastIndex)
	m.DefineS("LastIndexAny", pkg.LastIndexAny)
	m.DefineS("LastIndexFunc", pkg.LastIndexFunc)
	m.DefineS("Map", pkg.Map)
	m.DefineS("NewReader", pkg.NewReader)
	m.DefineS("NewReplacer", pkg.NewReplacer)
	m.DefineS("Repeat", pkg.Repeat)
	m.DefineS("Replace", pkg.Replace)
	m.DefineS("Split", pkg.Split)
	m.DefineS("SplitAfter", pkg.SplitAfter)
	m.DefineS("SplitAfterN", pkg.SplitAfterN)
	m.DefineS("SplitN", pkg.SplitN)
	m.DefineS("Title", pkg.Title)
	m.DefineS("ToLower", pkg.ToLower)
	m.DefineS("ToLowerSpecial", pkg.ToLowerSpecial)
	m.DefineS("ToTitle", pkg.ToTitle)
	m.DefineS("ToTitleSpecial", pkg.ToTitleSpecial)
	m.DefineS("ToUpper", pkg.ToUpper)
	m.DefineS("ToUpperSpecial", pkg.ToUpperSpecial)
	m.DefineS("Trim", pkg.Trim)
	m.DefineS("TrimFunc", pkg.TrimFunc)
	m.DefineS("TrimLeft", pkg.TrimLeft)
	m.DefineS("TrimLeftFunc", pkg.TrimLeftFunc)
	m.DefineS("TrimPrefix", pkg.TrimPrefix)
	m.DefineS("TrimRight", pkg.TrimRight)
	m.DefineS("TrimRightFunc", pkg.TrimRightFunc)
	m.DefineS("TrimSpace", pkg.TrimSpace)
	m.DefineS("TrimSuffix", pkg.TrimSuffix)
	return m
}
