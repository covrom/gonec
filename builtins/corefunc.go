package core

import "fmt"
import "reflect"

// VMFunc функции такого типа создаются на языке Гонец,
// их можно использовать в стандартной библиотеке, проверив на этот тип
type VMFunc func(args ...interface{}) (interface{}, error)

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
type VMMeth func(args VMSlicer) (VMInterfacer, error)

func (f VMMeth) vmval() {}

func (f VMMeth) Interface() interface{} {
	return f
}

func (f VMMeth) Func() VMFunc {
	// возвращает обертку, которая потребует первым параметром ссылку на объект метаданных (интерфейс VMMetaStructer)
	return VMFunc(func(args ...interface{}) (interface{}, error) {
		as := make(VMSlice, len(args))
		for i := range args {
			as[i] = ReflectToVMValue(reflect.ValueOf(args[i]))
		}

		v, err := f(as)

		if err != nil {
			return nil, err
		}
		return v.Interface(), nil
	})
}
