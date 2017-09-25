package core

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
)

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMBool bool

var ReflectVMBool = reflect.TypeOf(VMBool(true))

func (x VMBool) vmval() {}

func (x VMBool) Interface() interface{} {
	return bool(x)
}

func (x *VMBool) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case bool:
		*x = VMBool(vv)
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		*x = VMBool(rv.Bool()) // выдаст панику, если это не булево
	}
}

func (x VMBool) String() string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func (x VMBool) Int() int64 {
	if x {
		return 1
	} else {
		return 0
	}
}

func (x VMBool) Float() float64 {
	if x {
		return 1.0
	} else {
		return 0.0
	}
}

func (x VMBool) Decimal() VMDecimal {
	if x {
		return VMDecimal(decimal.New(1, 0))
	} else {
		return VMDecimal(decimal.New(0, 0))
	}
}

func (x VMBool) Bool() bool {
	return bool(x)
}

func (x VMBool) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func ParseVMBool(s string) (VMBool, error) {
	switch strings.ToLower(s) {
	case "true", "истина":
		return true, nil
	case "false", "ложь":
		return false, nil
	}
	return false, fmt.Errorf("Неверное значение для преобразования в булево")
}
