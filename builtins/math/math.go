// Package math implements math interface for anko script.
package math

import (
	t "math"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("math")
	m.Define("Abs", t.Abs)
	m.Define("Acos", t.Acos)
	m.Define("Acosh", t.Acosh)
	m.Define("Asin", t.Asin)
	m.Define("Asinh", t.Asinh)
	m.Define("Atan", t.Atan)
	m.Define("Atan2", t.Atan2)
	m.Define("Atanh", t.Atanh)
	m.Define("Cbrt", t.Cbrt)
	m.Define("Ceil", t.Ceil)
	m.Define("Copysign", t.Copysign)
	m.Define("Cos", t.Cos)
	m.Define("Cosh", t.Cosh)
	m.Define("Dim", t.Dim)
	m.Define("Erf", t.Erf)
	m.Define("Erfc", t.Erfc)
	m.Define("Exp", t.Exp)
	m.Define("Exp2", t.Exp2)
	m.Define("Expm1", t.Expm1)
	m.Define("Float32bits", t.Float32bits)
	m.Define("Float32frombits", t.Float32frombits)
	m.Define("Float64bits", t.Float64bits)
	m.Define("Float64frombits", t.Float64frombits)
	m.Define("Floor", t.Floor)
	m.Define("Frexp", t.Frexp)
	m.Define("Gamma", t.Gamma)
	m.Define("Hypot", t.Hypot)
	m.Define("Ilogb", t.Ilogb)
	m.Define("Inf", t.Inf)
	m.Define("IsInf", t.IsInf)
	m.Define("IsNaN", t.IsNaN)
	m.Define("J0", t.J0)
	m.Define("J1", t.J1)
	m.Define("Jn", t.Jn)
	m.Define("Ldexp", t.Ldexp)
	m.Define("Lgamma", t.Lgamma)
	m.Define("Log", t.Log)
	m.Define("Log10", t.Log10)
	m.Define("Log1p", t.Log1p)
	m.Define("Log2", t.Log2)
	m.Define("Logb", t.Logb)
	m.Define("Max", t.Max)
	m.Define("Min", t.Min)
	m.Define("Mod", t.Mod)
	m.Define("Modf", t.Modf)
	m.Define("NaN", t.NaN)
	m.Define("Nextafter", t.Nextafter)
	m.Define("Pow", t.Pow)
	m.Define("Pow10", t.Pow10)
	m.Define("Remainder", t.Remainder)
	m.Define("Signbit", t.Signbit)
	m.Define("Sin", t.Sin)
	m.Define("Sincos", t.Sincos)
	m.Define("Sinh", t.Sinh)
	m.Define("Sqrt", t.Sqrt)
	m.Define("Tan", t.Tan)
	m.Define("Tanh", t.Tanh)
	m.Define("Trunc", t.Trunc)
	m.Define("Y0", t.Y0)
	m.Define("Y1", t.Y1)
	m.Define("Yn", t.Yn)
	return m
}
