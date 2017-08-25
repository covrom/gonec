package bincode

import "github.com/covrom/gonec/ast"

type BinStmt interface {
	ast.Pos
	binstmt()
}

type BinStmtImpl struct {
	ast.PosImpl
}

func (x *BinStmtImpl) binstmt() {}

type BinCode []BinStmt

//////////////////////
// команды байткода
//////////////////////

type BinLOAD struct {
	BinStmtImpl

	Reg int
	Val interface{}
}

type BinCASTNUM struct {
	BinStmtImpl

	Reg int
}

type BinMAKESLICE struct {
	BinStmtImpl

	Reg int
	Len int
	Cap int
}

type BinSETIDX struct {
	BinStmtImpl

	Reg    int
	Index  int
	ValReg int
}

type BinMAKEMAP struct {
	BinStmtImpl

	Reg int
	Len int
}

type BinSETKEY struct {
	BinStmtImpl

	Reg    int
	Key    string
	ValReg int
}

type BinGET struct {
	BinStmtImpl

	Reg    int
	Id     int
	Dotted bool // содержит точку "."
}

type BinUNARY struct {
	BinStmtImpl

	Reg int
	Op  rune // - ! ^
}

type BinADDR struct {
	BinStmtImpl

	Reg int
}

type BinUNREF struct {
	BinStmtImpl

	Reg int
}

type BinLABEL struct {
	BinStmtImpl

	Label int
}

type BinJMP struct {
	BinStmtImpl

	JumpTo int
}

type BinJTRUE struct {
	BinStmtImpl

	Reg    int
	JumpTo int
}

type BinJFALSE struct {
	BinStmtImpl

	Reg    int
	JumpTo int
}

const (
	_    = iota
	ADD  // +
	SUB  // -
	MUL  // *
	QUO  // /
	REM  // %
	EQL  // ==
	NEQ  // !=
	GTR  // >
	GEQ  // >=
	LSS  // <
	LEQ  // <=
	OR   // |
	LOR  // ||
	AND  // &
	LAND // &&
	POW  //**
	SHL  // <<
	SHR  // >>
)

var OperMap = map[string]int{
	"+":  ADD,  // +
	"-":  SUB,  // -
	"*":  MUL,  // *
	"/":  QUO,  // /
	"%":  REM,  // %
	"==": EQL,  // ==
	"!=": NEQ,  // !=
	">":  GTR,  // >
	">=": GEQ,  // >=
	"<":  LSS,  // <
	"<=": LEQ,  // <=
	"|":  OR,   // |
	"||": LOR,  // ||
	"&":  AND,  // &
	"&&": LAND, // &&
	"**": POW,  //**
	"<<": SHL,  // <<
	">>": SHR,  // >>
}

type BinOPER struct {
	BinStmtImpl

	RegL int
	RegR int
	Op   int
}
