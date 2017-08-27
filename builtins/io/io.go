// Package io implements io interface for anko script.
package io

import (
	pkg "io"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("io")
	m.DefineS("Copy", pkg.Copy)
	m.DefineS("CopyN", pkg.CopyN)
	m.DefineS("EOF", pkg.EOF)
	m.DefineS("ErrClosedPipe", pkg.ErrClosedPipe)
	m.DefineS("ErrNoProgress", pkg.ErrNoProgress)
	m.DefineS("ErrShortBuffer", pkg.ErrShortBuffer)
	m.DefineS("ErrShortWrite", pkg.ErrShortWrite)
	m.DefineS("ErrUnexpectedEOF", pkg.ErrUnexpectedEOF)
	m.DefineS("LimitReader", pkg.LimitReader)
	m.DefineS("MultiReader", pkg.MultiReader)
	m.DefineS("MultiWriter", pkg.MultiWriter)
	m.DefineS("NewSectionReader", pkg.NewSectionReader)
	m.DefineS("Pipe", pkg.Pipe)
	m.DefineS("ReadAtLeast", pkg.ReadAtLeast)
	m.DefineS("ReadFull", pkg.ReadFull)
	m.DefineS("TeeReader", pkg.TeeReader)
	m.DefineS("WriteString", pkg.WriteString)
	return m
}
