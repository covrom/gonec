package core

import (
	"reflect"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMInt int64

func (x VMInt) vmval() {}

func (x VMInt) Interface() interface{} {
	return int64(x)
}

func (x *VMInt) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMInt(vv)
	case int8:
		*x = VMInt(vv)
	case int16:
		*x = VMInt(vv)
	case int32:
		*x = VMInt(vv)
	case int64:
		*x = VMInt(vv)
	case uint:
		*x = VMInt(vv)
	case uint8:
		*x = VMInt(vv)
	case uint16:
		*x = VMInt(vv)
	case uint32:
		*x = VMInt(vv)
	case uint64:
		*x = VMInt(vv)
	case uintptr:
		*x = VMInt(vv)
	case float32:
		*x = VMInt(int64(vv))
	case float64:
		*x = VMInt(int64(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		*x = VMInt(rv.Int()) // выдаст панику, если это не число
	}
}

func (x VMInt) String() string {
	return strconv.FormatInt(int64(x), 10)
}

func (x VMInt) Int() int64 {
	return int64(x)
}

func (x VMInt) Float() float64 {
	return float64(x)
}

func (x VMInt) Decimal() VMDecimal {
	return VMDecimal(decimal.New(int64(x), 0))
}

func (x VMInt) Bool() bool {
	return x > 0
}

func (x VMInt) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMInt) Time() VMTime {
	return VMTime(time.Unix(int64(x), 0))
}

func (x *VMInt) Parse(s string) error {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*x = VMInt(i64)
	return nil
}

// TODO: реализовать VMDurationer