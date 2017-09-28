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
	if strings.ToLower(s) != "неопределено" {
		return errors.New("Значение определено")
	}
	return nil
}
func (x VMNilType) BinaryType() VMBinaryType {
	return VMNIL
}

// TODO: маршалинг Nil,NULL и сравнение их со значениями

var VMNil = VMNilType{}

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
