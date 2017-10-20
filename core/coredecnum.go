package core

import (
	"reflect"

	"github.com/covrom/decnum"
)

// VMDecNum с плавающей токой, для финансовых расчетов высокой точности (decnum)
type VMDecNum struct {
	num decnum.Quad
}

var ReflectVMDecNum = reflect.TypeOf(VMDecNum{num: decnum.Zero()})

func (x VMDecNum) vmval() {}

func (x VMDecNum) Interface() interface{} {
	return x.num
}

func (x *VMDecNum) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int8:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int16:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int32:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int64:
		*x = VMDecNum{num: decnum.FromInt64(vv)}
	case uint:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint8:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint16:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint32:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint64:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uintptr:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case float32:
		*x = VMDecNum{num: decnum.FromFloat(float64(vv))}
	case float64:
		*x = VMDecNum{num: decnum.FromFloat(vv)}
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Float32 || rv.Kind() == reflect.Float64 {
			*x = VMDecNum{num: decnum.FromFloat(rv.Float())}
		} else {
			*x = VMDecNum{num: decnum.FromInt64(rv.Int())}
		}
	}
}
