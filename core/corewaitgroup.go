package core

import (
	"reflect"
	"sync"

	"github.com/covrom/gonec/names"
)

// VMWaitGroup - группа ожидания исполнения горутин
type VMWaitGroup struct {
	wg sync.WaitGroup
}

var ReflectVMWaitGroup = reflect.TypeOf(VMWaitGroup{})

func (x *VMWaitGroup) vmval() {}

func (x *VMWaitGroup) Interface() interface{} {
	return x
}

func (x *VMWaitGroup) Add(delta int) {
	x.wg.Add(delta)
}

func (x *VMWaitGroup) Done() {
	x.wg.Done()
}

func (x *VMWaitGroup) Wait() {
	x.wg.Wait()
}

func (x *VMWaitGroup) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!
	switch names.UniqueNames.GetLowerCase(name) {
	case "добавить":
		return VMFuncMustParams(1, x.Добавить), true
	case "завершить":
		return VMFuncMustParams(0, x.Завершить), true
	case "ожидать":
		return VMFuncMustParams(0, x.Ожидать), true
	}
	return nil, false
}

func (x *VMWaitGroup) Добавить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMInt)
	if !ok {
		return VMErrorNeedInt
	}
	x.Add(int(v))
	return nil
}

func (x *VMWaitGroup) Завершить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	x.Done()
	return nil
}

func (x *VMWaitGroup) Ожидать(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	x.Wait()
	return nil
}
