package core

import "reflect"

type VMStringMap map[string]VMValuer

var ReflectVMStringMap = reflect.TypeOf(make(VMStringMap, 0))

func (x VMStringMap) vmval() {}

func (x VMStringMap) Interface() interface{} {
	return x
}

func (x VMStringMap) StringMap() VMStringMap {
	return x
}

func (x VMStringMap) Len() VMInt {
	return VMInt(len(x))
}

func (x VMStringMap) Index(i VMValuer) VMValuer {
	if s, ok := i.(VMString); ok {
		return x[string(s)]
	}
	panic("Индекс должен быть строкой")
}

// TODO: equal, convert

// TODO: маршаллинг по аналогии с VMSlice!!!