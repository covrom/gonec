package core

import (
	"encoding/json"
	"reflect"
	"strconv"
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
