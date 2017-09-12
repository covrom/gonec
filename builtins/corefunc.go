package core

import "fmt"

// VMFunc функции такого типа создаются на языке Гонец,
// их можно использовать в стандартной библиотеке, проверив на этот тип
type VMFunc func(args ...interface{}) (interface{}, error)

// TODO: переделать на VMInterfacer

func (f VMFunc) vmval() {}

func (f VMFunc) Interface() interface{} {
	return f
}

func (f VMFunc) Func() VMFunc {
	return f
}

func (f VMFunc) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

// VMMeth вызывается как обертка метода объекта метаданных
// возвращаемое из обертки значение должно быть приведено к типу вирт. машины
type VMMeth func(receiver VMMetaStructer, args VMSlicer) (VMInterfacer, error)

func (f VMMeth) vmval() {}

func (f VMMeth) Interface() interface{} {
	return f
}

func (f VMMeth) Func() VMFunc {
	// возвращает обертку, которая потребует первым параметром ссылку на объект метаданных (интерфейс VMMetaStructer)
	return VMFunc(func(args ...interface{}) (interface{}, error) {
		if len(args) == 0 {
			panic("Отсутствует объект метаданных")
		}
		if ms, ok := args[0].(VMMetaStructer); ok {
			return f(ms, VMSlice(args[1:])) // непосредственно вызов метода и возврат его результата
			// результат возвращается в типах вирт. машины (интерфейс VMInterfacer)
		}
		panic("Отсутствует объект метаданных")
	})
}
