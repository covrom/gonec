package core

import (
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

// VMDecimal с плавающей токой, для финансовых расчетов высокой точности (decimal)
type VMDecimal decimal.Decimal

func (x VMDecimal) vmval() {}

func (x VMDecimal) Interface() interface{} {
	return decimal.Decimal(x)
}

func (x *VMDecimal) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int64:
		*x = VMDecimal(decimal.New(vv, 0))
	case uint:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint64:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uintptr:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case float32:
		*x = VMDecimal(decimal.NewFromFloat(float64(vv)))
	case float64:
		*x = VMDecimal(decimal.NewFromFloat(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Float32 || rv.Kind() == reflect.Float64 {
			*x = VMDecimal(decimal.NewFromFloat(rv.Float()))
		} else {
			*x = VMDecimal(decimal.New(rv.Int(), 0)) // выдаст панику, если это не число
		}
	}
}

func (x VMDecimal) String() string {
	return decimal.Decimal(x).String()
}

func (x VMDecimal) Int() int64 {
	return decimal.Decimal(x).IntPart()
}

func (x VMDecimal) Float() float64 {
	f64, ok := decimal.Decimal(x).Float64()
	if !ok {
		panic("Невозможно получить значение с плавающей запятой 64 бит")
	}
	return f64
}

func (x VMDecimal) Decimal() VMDecimal {
	return x
}

func (x VMDecimal) Bool() bool {
	return decimal.Decimal(x).GreaterThan(decimal.Zero)
}

func (x VMDecimal) MakeChan(size int) VMChaner {

	return make(VMChan, size)
}

func (x VMDecimal) Time() VMTime {
	intpart := decimal.Decimal(x).IntPart()
	nanopart := decimal.Decimal(x).Sub(decimal.New(intpart, 0)).Mul(decimal.New(1e9, 0)).IntPart()
	return VMTime(time.Unix(intpart, nanopart))
}

func (x *VMDecimal) Parse(s string) error {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	*x = VMDecimal(d)
	return nil
}

func (x VMDecimal) Add(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Add(decimal.Decimal(d2)))
}

func (x VMDecimal) Mul(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Mul(decimal.Decimal(d2)))
}