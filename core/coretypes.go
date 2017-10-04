package core

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/covrom/gonec/names"
	"github.com/dchest/siphash"
	"github.com/shopspring/decimal"
)

type VMOperation int

const (
	_    VMOperation = iota
	ADD              // +
	SUB              // -
	MUL              // *
	QUO              // /
	REM              // %
	EQL              // ==
	NEQ              // !=
	GTR              // >
	GEQ              // >=
	LSS              // <
	LEQ              // <=
	OR               // |
	LOR              // ||
	AND              // &
	LAND             // &&
	POW              //**
	SHL              // <<
	SHR              // >>
)

var OperMap = map[string]VMOperation{
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

var OperMapR = map[VMOperation]string{
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

type VMBinaryType byte

const (
	_ VMBinaryType = iota
	VMBOOL
	VMINT
	VMDECIMAL
	VMSTRING
	VMSLICE
	VMSTRINGMAP
	VMTIME
	VMDURATION
	VMNIL
	VMNULL
)

func (x VMBinaryType) ParseBinary(data []byte) (VMValuer, error) {
	switch x {
	case VMBOOL:
		var v VMBool
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMINT:
		var v VMInt
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMDECIMAL:
		var v VMDecimal
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMSTRING:
		var v VMString
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMSLICE:
		var v VMSlice
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMSTRINGMAP:
		var v VMStringMap
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMTIME:
		var v VMTime
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMDURATION:
		var v VMTimeDuration
		err := (&v).UnmarshalBinary(data)
		return v, err
	case VMNIL:
		return VMNil, nil
	case VMNULL:
		return VMNullVar, nil
	}
	return nil, errors.New("Неизвестный тип данных")
}

// nil значение для интерпретатора

type VMNilType struct{}

func (x VMNilType) vmval()                 {}
func (x VMNilType) String() string         { return "Неопределено" }
func (x VMNilType) Interface() interface{} { return nil }
func (x VMNilType) ParseGoType(v interface{}) {
	if v != nil {
		panic("Значение определено")
	}
}
func (x VMNilType) Parse(s string) error {
	if names.FastToLower(s) != "неопределено" {
		return errors.New("Значение определено")
	}
	return nil
}
func (x VMNilType) BinaryType() VMBinaryType {
	return VMNIL
}

var VMNil = VMNilType{}

// EvalBinOp сравнивает два значения или выполняет бинарную операцию
func (x VMNilType) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch y.(type) {
		case VMNullType, VMNilType:
			return VMBool(true), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch y.(type) {
		case VMNullType, VMNilType:
			return VMBool(false), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case GTR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case GEQ:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LSS:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LEQ:
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

func (x VMNilType) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (x *VMNilType) UnmarshalBinary(data []byte) error {
	*x = VMNil
	return nil
}

func (x VMNilType) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMNilType) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMNilType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *VMNilType) UnmarshalText(data []byte) error {
	*x = VMNil
	return nil
}

func (x VMNilType) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (x *VMNilType) UnmarshalJSON(data []byte) error {
	*x = VMNil
	return nil
}

// тип NULL

type VMNullType struct{}

func (x VMNullType) vmval()                 {}
func (x VMNullType) null()                  {}
func (x VMNullType) String() string         { return "NULL" }
func (x VMNullType) Interface() interface{} { return x }
func (x VMNullType) BinaryType() VMBinaryType {
	return VMNULL
}

var VMNullVar = VMNullType{}

// EvalBinOp сравнивает два значения или выполняет бинарную операцию
func (x VMNullType) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch y.(type) {
		case VMNullType, VMNilType:
			return VMBool(true), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch y.(type) {
		case VMNullType, VMNilType:
			return VMBool(false), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case GTR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case GEQ:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LSS:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LEQ:
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

func (x VMNullType) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (x *VMNullType) UnmarshalBinary(data []byte) error {
	*x = VMNullVar
	return nil
}

func (x VMNullType) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMNullType) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMNullType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *VMNullType) UnmarshalText(data []byte) error {
	*x = VMNullVar
	return nil
}

func (x VMNullType) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (x *VMNullType) UnmarshalJSON(data []byte) error {
	*x = VMNullVar
	return nil
}

// HashBytes хэширует байты по алгоритму SipHash-2-4

func HashBytes(buf []byte) uint64 {
	return siphash.Hash(0xdda7806a4847ec61, 0xb5940c2623a5aabd, buf)
}

func MustGenerateRandomString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// ReflectToVMValue преобразовывает значение Го в наиболее подходящий тип значения для вирт. машшины
func ReflectToVMValue(rv reflect.Value) VMInterfacer {
	if !rv.IsValid() {
		return VMNil
	}
	if x, ok := rv.Interface().(VMInterfacer); ok {
		return x
	}
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return VMInt(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return VMInt(rv.Uint())
	case reflect.String:
		return VMString(rv.String())
	case reflect.Bool:
		return VMBool(rv.Bool())
	case reflect.Float32, reflect.Float64:
		return VMDecimal(decimal.NewFromFloat(rv.Float()))
	case reflect.Chan:
		// проверяем, может это VMChaner
		if x, ok := rv.Interface().(VMChaner); ok {
			return x
		}
	case reflect.Array, reflect.Slice:
		// проверяем, может это VMSlicer
		if x, ok := rv.Interface().(VMSlicer); ok {
			return x
		}
	case reflect.Map:
		// проверяем, может это VMStringMaper
		if x, ok := rv.Interface().(VMStringMaper); ok {
			return x
		}
	case reflect.Func:
		// проверяем, может это VMFuncer
		if x, ok := rv.Interface().(VMFuncer); ok {
			return x
		}
	case reflect.Struct:
		switch v := rv.Interface().(type) {
		case decimal.Decimal:
			return VMDecimal(v)
		case time.Time:
			return VMTime(v)
		case VMNumberer:
			return v
		case VMDateTimer:
			return v
		case VMMetaObject:
			return v
		}
	}
	panic("Невозможно привести к типу интерпретатора")
}

func VMValuerFromJSON(s string) (VMValuer, error) {
	var i64 int64
	var err error
	if strings.HasPrefix(s, "0x") {
		i64, err = strconv.ParseInt(s[2:], 16, 64)
	} else {
		i64, err = strconv.ParseInt(s, 10, 64)
	}
	if err == nil {
		return VMInt(i64), nil
	}
	d, err := decimal.NewFromString(s)
	if err == nil {
		return VMDecimal(d), nil
	}
	var rwi interface{}
	if err = json.Unmarshal([]byte(s), &rwi); err != nil {
		return nil, err
	}
	// bool, for JSON booleans
	// float64, for JSON numbers
	// string, for JSON strings
	// []interface{}, for JSON arrays
	// map[string]interface{}, for JSON objects
	// nil for JSON null
	switch w := rwi.(type) {
	case string:
		return VMString(w), nil
	case bool:
		return VMBool(w), nil
	case float64:
		return VMDecimal(decimal.NewFromFloat(w)), nil
	case []interface{}:
		return VMSliceFromJson(s)
	case map[string]interface{}:
		return VMStringMapFromJson(s)
	case nil:
		return VMNil, nil
	default:
		return VMNil, errors.New("Невозможно определить значение")
	}
}

func VMSliceFromJson(x string) (VMSlice, error) {
	//парсим json из строки и пытаемся получить массив
	var rvms VMSlice
	var rm []json.RawMessage
	var err error
	if err = json.Unmarshal([]byte(x), &rm); err != nil {
		return rvms, err
	}
	rvms = make(VMSlice, len(rm))
	for i, raw := range rm {
		rvms[i], err = VMValuerFromJSON(string(raw))
		if err != nil {
			return rvms, err
		}
	}
	return rvms, nil
}

func VMStringMapFromJson(x string) (VMStringMap, error) {
	//парсим json из строки и пытаемся получить массив
	var rvms VMStringMap
	var rm map[string]json.RawMessage
	var err error
	if err = json.Unmarshal([]byte(x), &rm); err != nil {
		return rvms, err
	}
	rvms = make(VMStringMap, len(rm))
	for i, raw := range rm {
		rvms[i], err = VMValuerFromJSON(string(raw))
		if err != nil {
			return rvms, err
		}
	}
	return rvms, nil
}

func EqualVMValues(v1, v2 VMValuer) bool {
	return BoolOperVMValues(v1, v2, EQL)
}

func BoolOperVMValues(v1, v2 VMValuer, op VMOperation) bool {
	if xop, ok := v1.(VMOperationer); ok {
		if yop, ok := v2.(VMOperationer); ok {
			cmp, err := xop.EvalBinOp(op, yop)
			if err == nil {
				if rcmp, ok := cmp.(VMBool); ok {
					return bool(rcmp)
				}
			}
		}
	}
	return false
}

func SortLessVMValues(v1, v2 VMValuer) bool {
	// числа
	if vi, ok := v1.(VMInt); ok {
		if vj, ok := v2.(VMInt); ok {
			return vi.Int() < vj.Int()
		}
		if vj, ok := v2.(VMDecimal); ok {
			vii := decimal.New(int64(vi), 0)
			return vii.LessThan(decimal.Decimal(vj))
		}
	}

	if vi, ok := v1.(VMDecimal); ok {
		if vj, ok := v2.(VMInt); ok {
			vjj := decimal.New(int64(vj), 0)
			return decimal.Decimal(vi).LessThan(vjj)
		}
		if vj, ok := v2.(VMDecimal); ok {
			return decimal.Decimal(vi).LessThan(decimal.Decimal(vj))
		}
	}

	// строки
	if vi, ok := v1.(VMString); ok {
		if vj, ok := v2.(VMString); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
		if vj, ok := v2.(VMInt); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
		if vj, ok := v2.(VMDecimal); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
	}

	// булево

	if vi, ok := v1.(VMBool); ok {
		if vj, ok := v2.(VMBool); ok {
			return !vi.Bool() && vj.Bool()
		}
	}

	// дата

	if vi, ok := v1.(VMTime); ok {
		if vj, ok := v2.(VMTime); ok {
			return vi.Before(vj)
		}
	}

	// длительность
	if vi, ok := v1.(VMTimeDuration); ok {
		if vj, ok := v2.(VMTimeDuration); ok {
			return int64(vi) < int64(vj)
		}
	}

	// прочее
	return BoolOperVMValues(v1, v2, LSS)
}
