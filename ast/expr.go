package ast

import (
	"strings"

	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/builtins"
	"github.com/covrom/gonec/pos"
)

// Expr provides all of interfaces for expression.
type Expr interface {
	pos.Pos
	expr()
	Simplify() Expr
	BinTo(*binstmt.BinStmts, int, *int, bool)
}

// ExprImpl provide commonly implementations for Expr.
type ExprImpl struct {
	pos.PosImpl // ExprImpl provide Pos() function.
}

// expr provide restraint interface.
func (x *ExprImpl) expr() {}

// NumberExpr provide Number expression.
type NumberExpr struct {
	ExprImpl
	Lit string
}

func (x *NumberExpr) Simplify() Expr {
	var rv core.VMValuer
	if strings.ContainsAny(x.Lit, ".e") {
		v := &core.VMDecimal{}
		if err := v.Parse(x.Lit); err != nil {
			return x
		}
		rv = *v
	} else {
		v := new(core.VMInt)
		if err := v.Parse(x.Lit); err != nil {
			return x
		}
		rv = *v
	}
	return &NativeExpr{Value: rv}
}

// StringExpr provide String expression.
type StringExpr struct {
	ExprImpl
	Lit string
}

func (x *StringExpr) Simplify() Expr {
	return &NativeExpr{Value: core.VMString(x.Lit)}
}

// ArrayExpr provide Array expression.
type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

func (x *ArrayExpr) Simplify() Expr {
	waserrors := false
	a := make(core.VMSlice, len(x.Exprs))
	for i := range x.Exprs {
		x.Exprs[i] = x.Exprs[i].Simplify()
		if v, ok := x.Exprs[i].(*NativeExpr); ok {
			a[i] = v.Value
		} else {
			waserrors = true
		}
	}
	if waserrors {
		return x
	} else {
		return &NativeExpr{Value: a}
	}
}

// PairExpr provide one of Map key/value pair.
type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

func (x *PairExpr) Simplify() Expr {
	x.Value = x.Value.Simplify()
	return x
}

// MapExpr provide Map expression.
type MapExpr struct {
	ExprImpl
	MapExpr map[string]Expr
}

func (x *MapExpr) Simplify() Expr {
	waserrors := false
	m := make(core.VMStringMap)
	for k, v := range x.MapExpr {
		vv := v.Simplify()
		x.MapExpr[k] = vv
		if arg, ok := vv.(*NativeExpr); ok {
			m[k] = arg.Value
		} else {
			waserrors = true
		}
	}
	if waserrors {
		return x
	} else {
		return &NativeExpr{Value: m}
	}
}

// IdentExpr provide identity expression.
type IdentExpr struct {
	ExprImpl
	Lit string
	Id  int
}

func (x *IdentExpr) Simplify() Expr { return x }

// UnaryExpr provide unary minus expression. ex: -1, ^1, ~1.
type UnaryExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

func (x *UnaryExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	if v, ok := x.Expr.(*NativeExpr); ok {
		if vv := v.Value.(core.VMUnarer); ok {
			oper := core.OperMap[x.Operator]
			rv, err := vv.EvalUnOp(oper)
			if err == nil {
				return &NativeExpr{Value: rv}
			}
		}
	}
	return x
}

// AddrExpr provide referencing address expression.
type AddrExpr struct {
	ExprImpl
	Expr Expr
}

func (x *AddrExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	return x
}

// DerefExpr provide dereferencing address expression.
type DerefExpr struct {
	ExprImpl
	Expr Expr
}

func (x *DerefExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	return x
}

// ParenExpr provide parent block expression.
type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

func (x *ParenExpr) Simplify() Expr {
	x.SubExpr = x.SubExpr.Simplify()
	if arg, ok := x.SubExpr.(*NativeExpr); ok {
		return arg
	}
	return x
}

// BinOpExpr provide binary operator expression.
type BinOpExpr struct {
	ExprImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

func (x *BinOpExpr) Simplify() Expr {
	allnative := true
	for i := range x.Lhss {
		x.Lhss[i] = x.Lhss[i].Simplify()
		if _, ok := x.Lhss[i].(*NativeExpr); !ok {
			allnative = false
		}
	}
	for i := range x.Rhss {
		x.Rhss[i] = x.Rhss[i].Simplify()
		if _, ok := x.Rhss[i].(*NativeExpr); !ok {
			allnative = false
		}
	}
	if len(x.Lhss) == 1 && len(x.Rhss) == 1 && allnative {
		if x1, ok := x.Lhss[0].(*NativeExpr).Value.(core.VMOperationer); ok {
			if x2, ok := x.Rhss[0].(*NativeExpr).Value.(core.VMOperationer); ok {
				oper := core.OperMap[x.Operator]
				rv, err := x1.EvalBinOp(oper, x2)
				if err == nil {
					return &NativeExpr{Value: rv}
				}
			}
		}
	}
	return x
}

type TernaryOpExpr struct {
	ExprImpl
	Expr Expr
	Lhs  Expr
	Rhs  Expr
}

func (x *TernaryOpExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	x.Lhs = x.Expr.Simplify()
	x.Rhs = x.Expr.Simplify()
	if v, ok := x.Expr.(*NativeExpr); ok {
		if b, ok := v.Value.(core.VMBooler); ok {
			if b.Bool() {
				return x.Lhs
			} else {
				return x.Rhs
			}
		}
	}
	return x
}

// CallExpr provide calling expression.
type CallExpr struct {
	ExprImpl
	Func     interface{}
	Name     int //string
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

func (x *CallExpr) Simplify() Expr {
	for i := range x.SubExprs {
		x.SubExprs[i] = x.SubExprs[i].Simplify()
	}
	return x
}

// AnonCallExpr provide anonymous calling expression. ex: func(){}().
type AnonCallExpr struct {
	ExprImpl
	Expr     Expr
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

func (x *AnonCallExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	for i := range x.SubExprs {
		x.SubExprs[i] = x.SubExprs[i].Simplify()
	}
	return x
}

// MemberExpr provide expression to refer menber.
type MemberExpr struct {
	ExprImpl
	Expr Expr
	Name int //string
}

func (x *MemberExpr) Simplify() Expr {
	x.Expr = x.Expr.Simplify()
	return x
}

// ItemExpr provide expression to refer Map/Array item.
type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

func (x *ItemExpr) Simplify() Expr {
	x.Value = x.Value.Simplify()
	x.Index = x.Index.Simplify()
	if v, ok := x.Value.(*NativeExpr); ok {
		if i, ok := x.Index.(*NativeExpr); ok {
			if vv, ok := v.Value.(core.VMSlicer); ok {
				if ii, ok := i.Value.(core.VMInt); ok {
					return &NativeExpr{Value: vv.Slice()[ii.Int()]}
				}
			}
			if vv, ok := v.Value.(core.VMStringMaper); ok {
				if ii, ok := i.Value.(core.VMString); ok {
					return &NativeExpr{Value: vv.StringMap()[ii.String()]}
				}
			}
		}
	}
	return x
}

// SliceExpr provide expression to refer slice of Array.
type SliceExpr struct {
	ExprImpl
	Value Expr
	Begin Expr
	End   Expr
}

func (x *SliceExpr) Simplify() Expr {
	x.Value = x.Value.Simplify()
	x.Begin = x.Begin.Simplify()
	x.End = x.End.Simplify()
	if v, ok := x.Value.(*NativeExpr); ok {
		if ib, ok := x.Begin.(*NativeExpr); ok {
			if ie, ok := x.End.(*NativeExpr); ok {
				if vv, ok := v.Value.(core.VMSlicer); ok {
					if iib, ok := ib.Value.(core.VMInt); ok {
						if iie, ok := ie.Value.(core.VMInt); ok {
							return &NativeExpr{Value: vv.Slice()[iib.Int():iie.Int()]}
						}
					}
				}
			}
		}
	}
	return x
}

// FuncExpr provide function expression.
type FuncExpr struct {
	ExprImpl
	Name   int //string
	Stmts  Stmts
	Args   []int //string
	VarArg bool
}

func (x *FuncExpr) Simplify() Expr {
	for i := range x.Stmts {
		x.Stmts[i].Simplify()
	}
	return x
}

// LetExpr provide expression to let variable.
type LetExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

func (x *LetExpr) Simplify() Expr {
	x.Lhs = x.Lhs.Simplify()
	x.Rhs = x.Rhs.Simplify()
	return x
}

// LetsExpr provide multiple expression of let.
// type LetsExpr struct {
// 	ExprImpl
// 	Lhss     []Expr
// 	Operator string
// 	Rhss     []Expr
// }

// AssocExpr provide expression to assoc operation.
type AssocExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

func (x *AssocExpr) Simplify() Expr {
	x.Lhs = x.Lhs.Simplify()
	x.Rhs = x.Rhs.Simplify()
	return x
}

// NewExpr provide expression to make new instance.
// type NewExpr struct {
// 	ExprImpl
// 	Type int //string
// }

// ConstExpr provide expression for constant variable.
type ConstExpr struct {
	ExprImpl
	Value string
}

func (x *ConstExpr) Simplify() Expr {
	switch strings.ToLower(x.Value) {
	case "истина", "true":
		return &NativeExpr{Value: core.VMBool(true)}
	case "ложь", "false":
		return &NativeExpr{Value: core.VMBool(false)}
	case "null":
		return &NativeExpr{Value: core.VMNullVar}
	}
	return x
}

type ChanExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

func (x *ChanExpr) Simplify() Expr {
	x.Lhs = x.Lhs.Simplify()
	x.Rhs = x.Rhs.Simplify()
	return x
}

type Type struct {
	Name int //string
}

type TypeCast struct {
	ExprImpl
	Type     int
	TypeExpr Expr // должен быть строкой
	CastExpr Expr
}

func (x *TypeCast) Simplify() Expr {
	x.TypeExpr = x.TypeExpr.Simplify()
	x.CastExpr = x.CastExpr.Simplify()
	return x
}

type MakeExpr struct {
	ExprImpl
	Type     int  //string
	TypeExpr Expr // должен быть строкой
}

func (x *MakeExpr) Simplify() Expr {
	x.TypeExpr = x.TypeExpr.Simplify()
	return x
}

type MakeChanExpr struct {
	ExprImpl
	// Type     int //string
	SizeExpr Expr
}

func (x *MakeChanExpr) Simplify() Expr {
	x.SizeExpr = x.SizeExpr.Simplify()
	return x
}

type MakeArrayExpr struct {
	ExprImpl
	// Type    int //string
	LenExpr Expr
	CapExpr Expr
}

func (x *MakeArrayExpr) Simplify() Expr {
	x.LenExpr = x.LenExpr.Simplify()
	x.CapExpr = x.CapExpr.Simplify()
	return x
}

// хранит реальное значение, рассчитанное на этапе оптимизации AST
type NativeExpr struct {
	ExprImpl
	Value core.VMValuer
}

func (x *NativeExpr) Simplify() Expr {
	return x
}
