package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// VMString строки
type VMString string

var ReflectVMString = reflect.TypeOf(VMString(""))

func (x VMString) vmval() {}

func (x VMString) Interface() interface{} {
	return string(x)
}

func (x VMString) String() string {
	return string(x)
}

func (x VMString) Int() int64 {
	i64, err := strconv.ParseInt(string(x), 10, 64)
	if err != nil {
		panic(err)
	}
	return i64
}

func (x VMString) Float() float64 {
	f64, err := strconv.ParseFloat(string(x), 64)
	if err != nil {
		panic(err)
	}
	return f64
}

func (x VMString) Decimal() VMDecimal {
	d, err := decimal.NewFromString(string(x))
	if err != nil {
		panic(err)
	}
	return VMDecimal(d)
}

func (x VMString) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMString) Time() VMTime {
	t, err := time.ParseInLocation("2006-01-02T15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC3339, string(x))
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102150405", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006 15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC1123, string(x))
	if err == nil {
		return VMTime(t)
	}
	panic("Неверный формат даты и времени")
}

func (x *VMString) Parse(s string) error {
	*x = VMString(s)
	return nil
}

func (x VMString) Bool() bool {
	switch strings.ToLower(string(x)) {
	case "истина", "true":
		return true
	}
	return false
}

func (x VMString) Slice() VMSlice {
	var rm VMSlice
	if err := json.Unmarshal([]byte(x), &rm); err != nil {
		panic(err)
	}
	return rm
}

func (x VMString) StringMap() VMStringMap {
	var rm VMStringMap
	if err := json.Unmarshal([]byte(x), rm); err != nil {
		panic(err)
	}
	return rm
}

func (x VMString) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMString:
			return VMString(string(x) + string(yy)), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SUB:
		switch yy := y.(type) {
		case VMString:
			return VMString(strings.Replace(string(x), string(yy), "", -1)), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMString(strings.Repeat(string(x), int(yy))), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case QUO:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case REM:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMString:
			return VMBool(!bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case GTR:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == 1), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case GEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == 1 || cmp == 0), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LSS:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == -1), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == -1 || cmp == 0), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case OR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LOR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case AND:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LAND:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case POW:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SHR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SHL:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	}
	return VMNil, fmt.Errorf("Неизвестная операция")
}

func (x VMString) ConvertToType(nt reflect.Type, skipCollections bool) (VMValuer, error) {
	// приведение к дате - исходное число в секундах
	switch nt {
	case ReflectVMString:
		return x, nil
	case ReflectVMInt:
		return VMInt(x.Int()), nil
	case ReflectVMTime:
		return x.Time(), nil
	case ReflectVMBool:
		return VMBool(x.Bool()), nil
	case ReflectVMDecimal:
		return x.Decimal(), nil
	// TODO: остальные преобразования - в слайс, структуру и т.п.
	}
	return VMNil, fmt.Errorf("Приведение к типу невозможно")
}
