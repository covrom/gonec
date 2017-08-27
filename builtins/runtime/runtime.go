// Package runtime implements runtime interface for anko script.
package runtime

import (
	pkg "runtime"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewModule("runtime")
	//m.DefineS("BlockProfile", pkg.BlockProfile)
	//m.DefineS("Breakpoint", pkg.Breakpoint)
	//m.DefineS("CPUProfile", pkg.CPUProfile)
	//m.DefineS("Caller", pkg.Caller)
	//m.DefineS("Callers", pkg.Callers)
	//m.DefineS("CallersFrames", pkg.CallersFrames)
	//m.DefineS("Compiler", pkg.Compiler)
	//m.DefineS("FuncForPC", pkg.FuncForPC)
	m.DefineS("GC", pkg.GC)
	m.DefineS("GOARCH", pkg.GOARCH)
	m.DefineS("GOMAXPROCS", pkg.GOMAXPROCS)
	m.DefineS("GOOS", pkg.GOOS)
	m.DefineS("GOROOT", pkg.GOROOT)
	//m.DefineS("Goexit", pkg.Goexit)
	//m.DefineS("GoroutineProfile", pkg.GoroutineProfile)
	//m.DefineS("Gosched", pkg.Gosched)
	//m.DefineS("LockOSThread", pkg.LockOSThread)
	//m.DefineS("MemProfile", pkg.MemProfile)
	//m.DefineS("MemProfileRate", pkg.MemProfileRate)
	//m.DefineS("NumCPU", pkg.NumCPU)
	//m.DefineS("NumCgoCall", pkg.NumCgoCall)
	//m.DefineS("NumGoroutine", pkg.NumGoroutine)
	//m.DefineS("ReadMemStats", pkg.ReadMemStats)
	//m.DefineS("ReadTrace", pkg.ReadTrace)
	//m.DefineS("SetBlockProfileRate", pkg.SetBlockProfileRate)
	//m.DefineS("SetCPUProfileRate", pkg.SetCPUProfileRate)
	//m.DefineS("SetFinalizer", pkg.SetFinalizer)
	//m.DefineS("Stack", pkg.Stack)
	//m.DefineS("StartTrace", pkg.StartTrace)
	//m.DefineS("StopTrace", pkg.StopTrace)
	//m.DefineS("ThreadCreateProfile", pkg.ThreadCreateProfile)
	//m.DefineS("UnlockOSThread", pkg.UnlockOSThread)
	//m.DefineS("Version", pkg.Version)
	return m
}
