package bincode

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/covrom/gonec/ast"
)

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

var OperMapR = map[int]string{
	ADD:  "+",  // +
	SUB:  "-",  // -
	MUL:  "*",  // *
	QUO:  "/",  // /
	REM:  "%",  // %
	EQL:  "==", // ==
	NEQ:  "!=", // !=
	GTR:  ">",  // >
	GEQ:  ">=", // >=
	LSS:  "<",  // <
	LEQ:  "<=", // <=
	OR:   "|",  // |
	LOR:  "||", // ||
	AND:  "&",  // &
	LAND: "&&", // &&
	POW:  "**", //**
	SHL:  "<<", // <<
	SHR:  ">>", // >>
}

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	Message string
	Pos     ast.Position
}

var (
	BreakError     = errors.New("Неверное применение оператора Прервать")
	ContinueError  = errors.New("Неверное применение оператора Продолжить")
	ReturnError    = errors.New("Неверное применение оператора Возврат")
	InterruptError = errors.New("Выполнение прервано")
	// JumpError      = errors.New("Переход на метку")
)

// NewStringError makes error interface with message.
func NewStringError(pos ast.Pos, err string) error {
	if pos == nil {
		return &Error{Message: err, Pos: ast.Position{1, 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// NewErrorf makes error interface with message.
func NewErrorf(pos ast.Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

// NewError makes error interface with message.
// This doesn't overwrite last error.
func NewError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError || err == ReturnError {
		return err
	}
	// if pe, ok := err.(*parser.Error); ok {
	// 	return pe
	// }
	if ee, ok := err.(*Error); ok {
		return ee
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

// Error returns the error message.
func (e *Error) Error() string {
	// учитываем вставку модуля _ по умолчанию - вычитаем 1 из номера строки
	return fmt.Sprintf("[%d:%d] %s", e.Pos.Line-1, e.Pos.Column, e.Message)
}

type CatchFunc func() string

// Функции такого типа создаются на языке Гонец,
// их можно использовать в стандартной библиотеке, проверив на этот тип
type Func func(args ...interface{}) (interface{}, error)

func (f Func) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

// коллекции вирт. машины

type VMSlice []interface{}

type VMStringMap map[string]interface{}

// Регистры виртуальной машины

type VMRegs struct {
	Reg          []interface{} // регистры значений
	Labels       map[int]int   // [label]=index в BinCode
	TryLabel     []int         // последний элемент - это метка на текущий обработчик CATCH
	TryRegErr    []int         // последний элемент - это регистр с ошибкой текущего обработчика
	ForBreaks    []int         // последний элемент - это метка для break
	ForContinues []int         // последний элемент - это метка для continue
}

const initlenregs = 20

func NewVMRegs(stmts BinCode) *VMRegs {
	//собираем мапу переходов
	lbls := make(map[int]int)
	for i, stmt := range stmts {
		if s, ok := stmt.(*BinLABEL); ok {
			lbls[s.Label] = i
		}
	}
	return &VMRegs{
		Reg:          make([]interface{}, initlenregs),
		Labels:       lbls,
		TryLabel:     make([]int, 0, 5),
		TryRegErr:    make([]int, 0, 5),
		ForBreaks:    make([]int, 0, 5),
		ForContinues: make([]int, 0, 5),
	}
}

func (v *VMRegs) Set(reg int, val interface{}) {
	if reg < len(v.Reg) {
		v.Reg[reg] = val
	} else {
		for reg >= len(v.Reg) {
			v.Reg = append(v.Reg, make([]interface{}, initlenregs)...)
		}
		v.Reg[reg] = val
	}
}

func (v *VMRegs) FreeFromReg(reg int) {
	// освобождаем память, начиная с reg, для сборщика мусора
	// v.Reg = v.Reg[:reg]
	for i := reg; i < len(v.Reg); i++ {
		v.Reg[i] = nil
	}
}

func (v *VMRegs) PushTry(reg, label int) {
	v.TryRegErr = append(v.TryRegErr, reg)
	v.TryLabel = append(v.TryLabel, label)
}

func (v *VMRegs) TopTryLabel() int {
	l := len(v.TryLabel)
	if l == 0 {
		return -1
	}
	return v.TryLabel[l-1]
}

func (v *VMRegs) PopTry() (reg int, label int) {
	l := len(v.TryLabel)
	if l == 0 {
		return -1, -1
	}
	reg = v.TryRegErr[l-1]
	v.TryRegErr = v.TryRegErr[0 : l-1]
	label = v.TryLabel[l-1]
	v.TryLabel = v.TryLabel[0 : l-1]
	return
}

func (v *VMRegs) PushBreak(label int) {
	v.ForBreaks = append(v.ForBreaks, label)
}

func (v *VMRegs) TopBreak() int {
	l := len(v.ForBreaks)
	if l == 0 {
		return -1
	}
	return v.ForBreaks[l-1]
}

func (v *VMRegs) PopBreak() (label int) {
	l := len(v.ForBreaks)
	if l == 0 {
		return -1
	}
	label = v.ForBreaks[l-1]
	v.ForBreaks = v.ForBreaks[0 : l-1]
	return
}

func (v *VMRegs) PushContinue(label int) {
	v.ForContinues = append(v.ForContinues, label)
}

func (v *VMRegs) TopContinue() int {
	l := len(v.ForContinues)
	if l == 0 {
		return -1
	}
	return v.ForContinues[l-1]
}

func (v *VMRegs) PopContinue() (label int) {
	l := len(v.ForContinues)
	if l == 0 {
		return -1
	}
	label = v.ForContinues[l-1]
	v.ForBreaks = v.ForContinues[0 : l-1]
	return
}
