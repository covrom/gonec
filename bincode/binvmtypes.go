package bincode

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/covrom/gonec/ast"
)

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
	return e.Message
}

type CatchFunc func() string

// Func is function interface to reflect functions internaly.
type Func func(args ...reflect.Value) (reflect.Value, error)

func (f Func) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

// Регистры виртуальной машины

type VMRegs struct {
	Reg       []interface{} // регистры значений
	Labels    map[int]int   // [label]=index в BinCode
	TryLabel  []int         // последний элемент - это метка на текущий обработчик CATCH
	TryRegErr []int         // последний элемент - это регистр с ошибкой текущего обработчика
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
		Reg:       make([]interface{}, initlenregs),
		Labels:    lbls,
		TryLabel:  make([]int, 0, 5),
		TryRegErr: make([]int, 0, 5),
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

func InvokeNumber(lit string) (interface{}, error) {
	if strings.Contains(lit, ".") || strings.Contains(lit, "e") || strings.Contains(lit, "E") {
		v, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return v, err
		}
		return v, nil
	}
	var i int64
	var err error
	if strings.HasPrefix(lit, "0x") {
		i, err = strconv.ParseInt(lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(lit, 10, 64)
	}
	if err != nil {
		return i, err
	}
	return i, nil
}
