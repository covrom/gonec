package ast

import (
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/pos"
)

// Stmt provides all of interfaces for statement.
type Stmt interface {
	pos.Pos
	stmt()
	Simplify()
	BinTo(*binstmt.BinStmts, int, *int)
}

// StmtImpl provide commonly implementations for Stmt..
type StmtImpl struct {
	pos.PosImpl // StmtImpl provide Pos() function.
}

// stmt provide restraint interface.
func (x *StmtImpl) stmt() {}

// ExprStmt provide expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

func (x *ExprStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
}

func (s *ExprStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	s.Expr.BinTo(bins, reg, lid, true)
	// *bins = append(*bins, addBinExpr(s.Expr, reg, lid, true)...)
}

// IfStmt provide "if/else" statement.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   []Stmt
	ElseIf []Stmt // This is array of IfStmt
	Else   []Stmt
}

func (x *IfStmt) Simplify() {
	x.If = x.If.Simplify()
	for _, st := range x.Then {
		st.Simplify()
	}
	for _, st := range x.ElseIf {
		st.Simplify()
	}
	for _, st := range x.Else {
		st.Simplify()
	}
}

// TryStmt provide "try/catch/finally" statement.
type TryStmt struct {
	StmtImpl
	Try []Stmt
	// Var     string
	Catch []Stmt
	// Finally []Stmt
}

func (x *TryStmt) Simplify() {
	for _, st := range x.Try {
		st.Simplify()
	}
	for _, st := range x.Catch {
		st.Simplify()
	}
}

// ForStmt provide "for in" expression statement.
type ForStmt struct {
	StmtImpl
	Var   int //string
	Value Expr
	Stmts []Stmt
}

func (x *ForStmt) Simplify() {
	x.Value = x.Value.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// NumForStmt name = expr1 to expr2
type NumForStmt struct {
	StmtImpl
	Name  int //string
	Expr1 Expr
	Expr2 Expr
	Stmts []Stmt
}

func (x *NumForStmt) Simplify() {
	x.Expr1 = x.Expr1.Simplify()
	x.Expr2 = x.Expr2.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// CForStmt provide C-style "for (;;)" expression statement.
// type CForStmt struct {
// 	StmtImpl
// 	Expr1 Expr
// 	Expr2 Expr
// 	Expr3 Expr
// 	Stmts []Stmt
// }

// LoopStmt provide "for expr" expression statement.
type LoopStmt struct {
	StmtImpl
	Expr  Expr
	Stmts []Stmt
}

func (x *LoopStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// BreakStmt provide "break" expression statement.
type BreakStmt struct {
	StmtImpl
}

func (x *BreakStmt) Simplify() {}

// ContinueStmt provide "continue" expression statement.
type ContinueStmt struct {
	StmtImpl
}

func (x *ContinueStmt) Simplify() {}

// ForStmt provide "return" expression statement.
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

func (x *ReturnStmt) Simplify() {
	for i := range x.Exprs {
		x.Exprs[i] = x.Exprs[i].Simplify()
	}
}

// ThrowStmt provide "throw" expression statement.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

func (x *ThrowStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
}

// ModuleStmt provide "module" expression statement.
type ModuleStmt struct {
	StmtImpl
	Name  int //string
	Stmts []Stmt
}

func (x *ModuleStmt) Simplify() {
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// VarStmt provide statement to let variables in current scope.
type VarStmt struct {
	StmtImpl
	Names []int //string
	Exprs []Expr
}

func (x *VarStmt) Simplify() {
	for i := range x.Exprs {
		x.Exprs[i] = x.Exprs[i].Simplify()
	}
}

// SwitchStmt provide switch statement.
type SwitchStmt struct {
	StmtImpl
	Expr  Expr
	Cases []Stmt
}

func (x *SwitchStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Cases {
		st.Simplify()
	}
}

// SelectStmt provide switch statement.
type SelectStmt struct {
	StmtImpl
	Cases []Stmt
}

func (x *SelectStmt) Simplify() {
	for _, st := range x.Cases {
		st.Simplify()
	}
}

// CaseStmt provide switch/case statement.
type CaseStmt struct {
	StmtImpl
	Expr  Expr
	Stmts []Stmt
}

func (x *CaseStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// DefaultStmt provide switch/default statement.
type DefaultStmt struct {
	StmtImpl
	Stmts []Stmt
}

func (x *DefaultStmt) Simplify() {
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

// LetsStmt provide multiple statement of let.
type LetsStmt struct {
	StmtImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

func (x *LetsStmt) Simplify() {
	for i := range x.Lhss {
		x.Lhss[i] = x.Lhss[i].Simplify()
	}
	for i := range x.Rhss {
		x.Rhss[i] = x.Rhss[i].Simplify()
	}
}
