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

type BinCALL struct {
	BinStmtImpl

	Name        int  // либо вызов по имени из ast.UniqueNames, если Name != 0
	// либо вызов обработчика (Name==0), напр. для анонимной функции (выражение типа func, или ссылка или интерфейс с ним, находится в reg)
	NumArgs     int  // число аргументов, которое надо взять на входе из регистров (<=7) или массива (Reg)
	RegArgs     int  // первый регистр из числа регистров с параметрами (параметров<=7) или регистр с массивом аругментов (>7)

	// в последнем регистре (из серии, если <=7, или в RegArgs, если >7) передан
	// массив аргументов переменной длины, и это приемлемо для вызываемой функции (оператор "...")
	// здесь надо быть аккуратным при числе аргументов >7
	// таким массивом будет только последний аргумент
	VarArg bool

	Go bool // признак необходимости запуска в новой горутине
}
