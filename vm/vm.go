package vm

import (
	"fmt"
	"io"

	"math/big"
	"time"

	"github.com/covrom/gonec/gonecparser/ast"
)

// универсальный вариантный тип данных

const (
	NONE = iota
	NULL
	UNDEF
	NUM
	DATE
	STR
)

type variant struct {
	typ  int //может быть одна из констант
	num  big.Float
	date time.Time
	str  string
}

func (v variant) String() string {
	switch v.typ {
	case NULL:
		return "NULL"
	case UNDEF:
		return "Неопределено"
	case NUM:
		return v.num.String()
	case DATE:
		return v.date.Format(time.RFC3339)
	case STR:
		return v.str
	default:
		return "NONE"
	}
}

func (v *variant) SetString(s string) {
	v.typ = STR
	v.str = s
	v.date = time.Time{}
	v.num = big.Float{}
}

func (v *variant) SetDate(d time.Time) {
	v.typ = DATE
	v.str = ""
	v.date = d
	v.num = big.Float{}
}

func (v *variant) SetNum(n big.Float) {
	v.typ = NUM
	v.str = ""
	v.date = time.Time{}
	v.num = n
}

func (v *variant) SetUNDEF() {
	v.typ = UNDEF
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
}

func (v *variant) SetNULL() {
	v.typ = NULL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
}

func (v variant) GetValue() (typ int, val interface{}) {
	typ = v.typ
	switch typ {
	case DATE:
		val = v.date
	case STR:
		val = v.str
	case NUM:
		val = v.num
	default:
		val = nil
	}
	return
}

func (v variant) IsNULL() bool {
	return v.typ == NULL
}

func (v variant) IsUNDEF() bool {
	return v.typ == UNDEF
}

func (v variant) IsDate() bool {
	return v.typ == DATE
}

func (v variant) IsNum() bool {
	return v.typ == NUM
}

func (v variant) IsString() bool {
	return v.typ == STR
}

// виртуальная машина

type VirtMachine struct {
	af *ast.File
	w  io.Writer
}

func NewVM(af *ast.File, w io.Writer) *VirtMachine {
	return &VirtMachine{af: af, w: w}
}

func (v *VirtMachine) Run() error {
	ast.Inspect(v.af, v.astInspect)
	return nil
}

func (v *VirtMachine) astInspect(n ast.Node) bool {
	var s string
	switch x := n.(type) {
	// TODO: исполнение __init__
	case *ast.GenDecl:

		// s = x.Value

	case *ast.Ident:
		s = x.Name
	}
	if s != "" {
		fmt.Println(s)
		return true
	}
	return false
}
