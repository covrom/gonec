package core

import "fmt"

// VMFunc вызывается как обертка метода объекта метаданных или обертка функции библиотеки
// возвращаемое из обертки значение должно быть приведено к типу вирт. машины
// функции такого типа создаются на языке Гонец,
// их можно использовать в стандартной библиотеке, проверив на этот тип
type VMFunc func(args VMSlicer) (VMValuer, error)

func (f VMFunc) vmval() {}

func (f VMFunc) Interface() interface{} {
	return f
}

func (f VMFunc) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

func (f VMFunc) Func() VMFunc {
	return f
}
