package core

import (
	"math"
	"reflect"
	"time"

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

func (x VMDecNum) String() string {
	return x.num.String()
}

func (x VMDecNum) Int() int64 {
	i, err := x.num.ToInt64(decnum.RoundDown) //целая часть, без округления
	if err != nil {
		return 0
	}
	return i
}

func (x VMDecNum) RoundHalfUp() int64 {
	i, err := x.num.ToInt64(decnum.RoundHalfUp) //целая часть, округление вверх, если модуль>0.5
	if err != nil {
		return 0
	}
	return i
}

func (x VMDecNum) Float() float64 {
	i, err := x.num.ToFloat64()
	if err != nil {
		return i
	}
	return i
}

func (x VMDecNum) DecNum() VMDecNum {
	return x
}

func (x VMDecNum) InvokeNumber() (VMNumberer, error) {
	return x, nil
}

func (x VMDecNum) Bool() bool {
	return x.num.IsPositive()
}

func (x VMDecNum) BinaryType() VMBinaryType {
	return VMDECNUM
}

func (x VMDecNum) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMDecNum) Time() VMTime {
	intpart := x.Int()
	rem := x.num.Sub(decnum.FromInt64(intpart))
	nano := rem.Mul(decnum.FromInt64(int64(VMSecond)))
	nanopart, _ := nano.ToInt64(decnum.RoundHalfUp)
	return VMTime(time.Unix(intpart, nanopart))
}

func (x VMDecNum) Duration() VMTimeDuration {
	i, _ := x.num.Mul(decnum.FromInt64(int64(VMSecond))).ToInt64(decnum.RoundHalfUp)
	return VMTimeDuration(time.Duration(i))
}

func ParseVMDecNum(s string) (VMDecNum, error) {
	d, err := decnum.FromString(s)
	return VMDecNum{num: d}, err
}

func (x VMDecNum) Add(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Add(d2.num)}
}

func (x VMDecNum) Sub(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Sub(d2.num)}
}

func (x VMDecNum) Mul(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Mul(d2.num)}
}

func (x VMDecNum) Div(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Div(d2.num)}
}

func (x VMDecNum) Mod(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Mod(d2.num)}
}

// Pow работает только в диапазоне и точности чисел float64
func (x VMDecNum) Pow(d2 VMDecNum) VMDecNum {
	v1, err := x.num.ToFloat64()
	if err != nil {
		panic(err)
	}
	v2, err := d2.num.ToFloat64()
	if err != nil {
		panic(err)
	}
	rv := math.Pow(v1, v2)
	return VMDecNum{num: decnum.FromFloat(rv)}
}

func (x VMDecNum) Equal(d2 VMDecNum) VMBool {
	return VMBool(x.num.Equal(d2.num))
}

func (x VMDecNum) NotEqual(d2 VMDecNum) VMBool {
	return VMBool(!x.num.Equal(d2.num))
}
