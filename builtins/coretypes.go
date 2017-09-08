package core

import (
	"reflect"
	"time"
)

// коллекции и типы вирт. машины

type VMSlice []interface{}

type VMStringMap map[string]interface{}

type VMTime time.Time

var ReflectVMTime = reflect.TypeOf(VMTime{})

func (t VMTime) Вычесть(t2 VMTime) time.Duration {
	return time.Time(t).Sub(time.Time(t2))
}
