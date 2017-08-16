package ast

import (
	"fmt"
	"strings"
	"sync"
)

// уникальные названия переменных, индекс используется в AST-дереве
type EnvNames struct {
	sync.RWMutex
	names   map[string]int
	handles map[int]string
	iter    int
}

func NewEnvNames() *EnvNames {
	en := EnvNames{
		names:   make(map[string]int, 200),
		handles: make(map[int]string, 200),
		iter:    1,
	}
	return &en
}

func (en *EnvNames) Set(n string) int {
	ns := strings.ToLower(n)
	en.RLock()
	if i, ok := en.names[ns]; ok {
		en.RUnlock()
		return i
	}
	en.RUnlock()
	en.Lock()
	i := en.iter
	en.names[ns] = i
	en.handles[i] = n
	en.iter++
	en.Unlock()
	return i
}

func (en *EnvNames) Get(i int) string {
	en.RLock()
	defer en.RUnlock()
	if s, ok := en.handles[i]; ok {
		return s
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

// все переменные
var UniqueNames = NewEnvNames()

// Expr provides all of interfaces for expression.
type Expr interface {
	Pos
	expr()
}

// ExprImpl provide commonly implementations for Expr.
type ExprImpl struct {
	PosImpl // ExprImpl provide Pos() function.
}

// expr provide restraint interface.
func (x *ExprImpl) expr() {}

// NumberExpr provide Number expression.
type NumberExpr struct {
	ExprImpl
	Lit string
}

// StringExpr provide String expression.
type StringExpr struct {
	ExprImpl
	Lit string
}

// ArrayExpr provide Array expression.
type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

// PairExpr provide one of Map key/value pair.
type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

// MapExpr provide Map expression.
type MapExpr struct {
	ExprImpl
	MapExpr map[string]Expr
}

// IdentExpr provide identity expression.
type IdentExpr struct {
	ExprImpl
	Lit string
	Id  int
}

// UnaryExpr provide unary minus expression. ex: -1, ^1, ~1.
type UnaryExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

// AddrExpr provide referencing address expression.
type AddrExpr struct {
	ExprImpl
	Expr Expr
}

// DerefExpr provide dereferencing address expression.
type DerefExpr struct {
	ExprImpl
	Expr Expr
}

// ParenExpr provide parent block expression.
type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

// BinOpExpr provide binary operator expression.
type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type TernaryOpExpr struct {
	ExprImpl
	Expr Expr
	Lhs  Expr
	Rhs  Expr
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

// AnonCallExpr provide anonymous calling expression. ex: func(){}().
type AnonCallExpr struct {
	ExprImpl
	Expr     Expr
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

// MemberExpr provide expression to refer menber.
type MemberExpr struct {
	ExprImpl
	Expr Expr
	Name int //string
}

// ItemExpr provide expression to refer Map/Array item.
type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

// SliceExpr provide expression to refer slice of Array.
type SliceExpr struct {
	ExprImpl
	Value Expr
	Begin Expr
	End   Expr
}

// FuncExpr provide function expression.
type FuncExpr struct {
	ExprImpl
	Name   int //string
	Stmts  []Stmt
	Args   []int //string
	VarArg bool
}

// LetExpr provide expression to let variable.
type LetExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

// LetsExpr provide multiple expression of let.
type LetsExpr struct {
	ExprImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

// AssocExpr provide expression to assoc operation.
type AssocExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

// NewExpr provide expression to make new instance.
type NewExpr struct {
	ExprImpl
	Type int //string
}

// ConstExpr provide expression for constant variable.
type ConstExpr struct {
	ExprImpl
	Value string
}

type ChanExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

type Type struct {
	Name int //string
}

type TypeCast struct {
	ExprImpl
	Type int
	CastExpr Expr
}

type MakeExpr struct {
	ExprImpl
	Type int //string
}

type MakeChanExpr struct {
	ExprImpl
	// Type     int //string
	SizeExpr Expr
}

type MakeArrayExpr struct {
	ExprImpl
	// Type    int //string
	LenExpr Expr
	CapExpr Expr
}
