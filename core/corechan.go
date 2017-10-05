package core

import (
	"github.com/covrom/gonec/names"
)

// VMChan - канал для передачи любого типа вирт. машины
type VMChan chan VMValuer

func (x VMChan) vmval() {}

func (x VMChan) Interface() interface{} {
	return x
}

func (x VMChan) Send(v VMValuer) {
	x <- v
}

func (x VMChan) Recv() (VMValuer, bool) {
	rv, ok := <-x
	return rv, ok
}

func (x VMChan) TrySend(v VMValuer) (ok bool) {
	select {
	case x <- v:
		ok = true
	default:
		ok = false
	}
	return
}

func (x VMChan) TryRecv() (v VMValuer, ok bool, notready bool) {
	select {
	case v, ok = <-x:
		notready = false
	default:
		ok = false
		notready = true
	}
	return
}

func (x VMChan) Close() { close(x) }

func (x VMChan) Size() int { return cap(x) }

func (x VMChan) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!
	switch names.UniqueNames.GetLowerCase(name) {
	case "закрыть":
		return VMFuncMustParams(0, x.Закрыть), true
	case "размер":
		return VMFuncMustParams(0, x.Размер), true
		// TODO: подключить соединение
	}
	return nil, false
}

func (x VMChan) Закрыть(args VMSlice, rets *VMSlice) error {
	x.Close()
	return nil
}

func (x VMChan) Размер(args VMSlice, rets *VMSlice) error {
	rets.Append(VMInt(x.Size()))
	return nil
}
