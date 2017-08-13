// Package flag implements flag interface for anko script.
package flag

import (
	pkg "flag"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("flag")
	m.DefineS("Arg", pkg.Arg)
	m.DefineS("Args", pkg.Args)
	m.DefineS("Bool", pkg.Bool)
	m.DefineS("BoolVar", pkg.BoolVar)
	m.DefineS("CommandLine", pkg.CommandLine)
	m.DefineS("ContinueOnError", pkg.ContinueOnError)
	m.DefineS("Duration", pkg.Duration)
	m.DefineS("DurationVar", pkg.DurationVar)
	m.DefineS("ErrHelp", pkg.ErrHelp)
	m.DefineS("ExitOnError", pkg.ExitOnError)
	m.DefineS("Float64", pkg.Float64)
	m.DefineS("Float64Var", pkg.Float64Var)
	m.DefineS("Int", pkg.Int)
	m.DefineS("Int64", pkg.Int64)
	m.DefineS("Int64Var", pkg.Int64Var)
	m.DefineS("IntVar", pkg.IntVar)
	m.DefineS("Lookup", pkg.Lookup)
	m.DefineS("NArg", pkg.NArg)
	m.DefineS("NFlag", pkg.NFlag)
	m.DefineS("NewFlagSet", pkg.NewFlagSet)
	m.DefineS("PanicOnError", pkg.PanicOnError)
	m.DefineS("Parse", pkg.Parse)
	m.DefineS("Parsed", pkg.Parsed)
	m.DefineS("PrintDefaults", pkg.PrintDefaults)
	m.DefineS("Set", pkg.Set)
	m.DefineS("String", pkg.String)
	m.DefineS("StringVar", pkg.StringVar)
	m.DefineS("Uint", pkg.Uint)
	m.DefineS("Uint64", pkg.Uint64)
	m.DefineS("Uint64Var", pkg.Uint64Var)
	m.DefineS("UintVar", pkg.UintVar)
	m.DefineS("Usage", pkg.Usage)
	m.DefineS("Var", pkg.Var)
	m.DefineS("Visit", pkg.Visit)
	m.DefineS("VisitAll", pkg.VisitAll)
	return m
}
