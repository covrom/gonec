package binstmt

import (
	"errors"
	"fmt"

	posit "github.com/covrom/gonec/pos"
)

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	Message string
	Pos     posit.Position
}

var (
	BreakError     = errors.New("Неверное применение оператора Прервать")
	ContinueError  = errors.New("Неверное применение оператора Продолжить")
	ReturnError    = errors.New("Неверное применение оператора Возврат")
	InterruptError = errors.New("Выполнение прервано")
)

// NewStringError makes error interface with message.
func NewStringError(pos posit.Pos, err string) error {
	if pos == nil {
		return &Error{Message: err, Pos: posit.Position{1, 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// NewErrorf makes error interface with message.
func NewErrorf(pos posit.Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

// NewError makes error interface with message.
// This doesn't overwrite last error.
func NewError(pos posit.Pos, err error) error {
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

func (e *Error) String() string {
	// учитываем вставку модуля _ по умолчанию - вычитаем 1 из номера строки
	return e.Message
}

