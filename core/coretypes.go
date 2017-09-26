package core

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"reflect"
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

var VMNil = VMNilType{}

// тип NULL

type VMNullType struct{}

func (x VMNullType) vmval()                 {}
func (x VMNullType) null()                  {}
func (x VMNullType) String() string         { return "NULL" }
func (x VMNullType) Interface() interface{} { return x }

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
