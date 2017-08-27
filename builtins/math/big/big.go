// Package big implements math/big interface for anko script.
package big

import (
	t "math/big"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewModule("big")
	m.DefineS("Above", t.Above)
	m.DefineS("AwayFromZero", t.AwayFromZero)
	m.DefineS("Below", t.Below)
	m.DefineS("Exact", t.Exact)
	m.DefineS("Jacobi", t.Jacobi)
	m.DefineS("MaxBase", t.MaxBase)
	m.DefineS("MaxExp", t.MaxExp)
	// TODO https://github.com/mattn/anko/issues/49
	//m.DefineS("MaxPrec", t.MaxPrec)
	m.DefineS("MinExp", t.MinExp)
	m.DefineS("NewFloat", t.NewFloat)
	m.DefineS("NewInt", t.NewInt)
	m.DefineS("NewRat", t.NewRat)
	m.DefineS("ParseFloat", t.ParseFloat)
	m.DefineS("ToNearestAway", t.ToNearestAway)
	m.DefineS("ToNearestEven", t.ToNearestEven)
	m.DefineS("ToNegativeInf", t.ToNegativeInf)
	m.DefineS("ToPositiveInf", t.ToPositiveInf)
	m.DefineS("ToZero", t.ToZero)
	return m
}
