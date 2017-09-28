package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	var i64 int64
	var err error
	if strings.HasPrefix(string(x), "0x") {
		i64, err = strconv.ParseInt(string(x)[2:], 16, 64)
	} else {
		i64, err = strconv.ParseInt(string(x), 10, 64)
	}
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

func (x VMString) InvokeNumber() (v VMNumberer, err error) {
	if strings.ContainsAny(string(x), ".eE") {
		v, err = ParseVMDecimal(string(x))
	} else {
		v, err = ParseVMInt(string(x))
	}
	return
}

func (x VMString) BinaryType() VMBinaryType {
	return VMSTRING
}

func (x VMString) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMString) Hash() VMString {
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes([]byte(x)))
	return VMString(hex.EncodeToString(h))
}

func (x VMString) Time() VMTime {
	t, err := time.Parse(time.RFC3339, string(x))
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("2006-01-02T15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("2006-01-02 15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006 15:04:05", string(x), time.Local)
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
	t, err = time.ParseInLocation("2006-01-02", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC1123, string(x))
	if err == nil {
		return VMTime(t)
	}
	panic("Неверный формат даты и времени")
}

func (x VMString) Bool() bool {
	r, _ := ParseVMBool(string(x))
	return r.Bool()
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
		return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		switch yy := y.(type) {
		case VMString:
			return VMString(strings.Replace(string(x), string(yy), "", -1)), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMString(strings.Repeat(string(x), int(yy))), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMString:
			return VMBool(!bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case GTR:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == 1), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case GEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == 1 || cmp == 0), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case LSS:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == -1), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case LEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == -1 || cmp == 0), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case OR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LOR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case AND:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LAND:
		return VMNil, errors.New("Операция между значениями невозможна")
	case POW:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SHR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SHL:
		return VMNil, errors.New("Операция между значениями невозможна")
	}
	return VMNil, errors.New("Неизвестная операция")
}

func (x VMString) ConvertToType(nt reflect.Type) (VMValuer, error) {
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
	case ReflectVMSlice:
		return VMSliceFromJson(string(x))
	case ReflectVMStringMap:
		return VMStringMapFromJson(string(x))
	}

	// попробуем десериализировать структуру из json
	if nt.Kind() == reflect.Struct {
		//парсим json из строки и пытаемся получить указатель на структуру
		rm := reflect.New(nt).Interface()
		if err := json.Unmarshal([]byte(x), rm); err != nil {
			return VMNil, err
		}
		if rv, ok := rm.(VMValuer); ok {
			return rv, nil
		}
		return VMNil, errors.New("Объект несовместим с типами интерпретатора")
	}
	return VMNil, errors.New("Приведение к типу невозможно")
}

func (x VMString) MarshalBinary() ([]byte, error) {
	return []byte(string(x)), nil
}

func (x *VMString) UnmarshalBinary(data []byte) error {
	*x = VMString(string(data))
	return nil
}

func (x VMString) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMString) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

// TODO:

// func (x VMString) MarshalJSON() ([]byte, error) {
// }

// func (x *VMString) UnmarshalJSON(data []byte) error {
// }

// func (x VMString) MarshalText() ([]byte, error) {
// }

// func (x *VMString) UnmarshalText(data []byte) error {
// }
