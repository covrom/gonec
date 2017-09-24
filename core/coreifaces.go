package core

import (
	"reflect"
)

// иерархия базовых типов вирт. машины
type (
	// VMValuer корневой тип всех значений, доступных вирт. машине
	VMValuer interface {
		vmval()
	}

	// VMInterfacer корневой тип всех значений,
	// которые могут преобразовываться в значения для функций на языке Го в родные типы Го
	VMInterfacer interface {
		VMValuer
		Interface() interface{} // в типах Го, может возвращать в т.ч. nil
	}

	// VMFromGoParser может парсить из значений на языке Го
	VMFromGoParser interface {
		VMValuer
		ParseGoType(interface{}) // используется для указателей, т.к. парсит в их значения
	}

	// TODO: реализовать VMOperationer во всех типах с возможностью операции с другим операндом

	// VMOperationer может выполнить операцию с другим значением, операцию сравнения или математическую
	VMOperationer interface {
		VMValuer
		EvalBinOp(VMOperation, VMOperationer) (VMValuer, error) // возвращает результат выражения с другим значением
	}

	// TODO: реализовать VMUnarer во всех типах с унарными операциями

	// VMUnarer может выполнить унарную операцию над свои значением
	VMUnarer interface {
		VMValuer
		EvalUnOp(rune) (VMValuer, error) // возвращает результат выражения с другим значением
	}

	// TODO: реализовать VMConverter во всех типах, кроме функций

	// VMConverter может конвертироваться в тип reflect.Type
	VMConverter interface {
		VMValuer
		ConvertToType(t reflect.Type, skipCollections bool) (VMValuer, error)
		// если skipCollections=true и вызов идет для значения-коллекции, нужно вернуть значение без преобразования
		// это требуется для рекурсивного преобразования коллекций
	}

	// VMParser может парсить из строки
	VMParser interface {
		VMInterfacer
		Parse(string) error // используется для указателей, т.к. парсит в их значения
	}

	// VMChaner реализует поведение канала
	VMChaner interface {
		VMInterfacer
		Send(VMValuer)
		Recv() (VMValuer, bool)
		TrySend(VMValuer) bool
		TryRecv() (VMValuer, bool, bool)
	}

	// VMIndexer имеет длину и значение по индексу
	VMIndexer interface {
		VMInterfacer
		Length() VMInt
		IndexVal(VMValuer) VMValuer
	}

	// конкретные типы виртуальной машины

	// VMStringer строка
	VMStringer interface {
		VMInterfacer
		String() string
	}

	// VMNumberer число, внутреннее хранение в int64 или decimal формате
	VMNumberer interface {
		VMParser
		Int() int64
		Float() float64
		Decimal() VMDecimal
		InvokeNumber() (VMNumberer, error) // извлекает VMInt или VMDecimal, в зависимости от наличия .eE
	}

	// VMBooler сообщает значение булево
	VMBooler interface {	
		VMParser
		Bool() bool
	}

	// VMSlicer может быть представлен в виде слайса Гонец
	VMSlicer interface {
		VMParser
		Slice() VMSlice
	}

	// VMStringMaper может быть представлен в виде структуры Гонец
	VMStringMaper interface {
		VMParser
		StringMap() VMStringMap
	}

	// VMFuncer это функция Гонец
	VMFuncer interface {
		VMInterfacer
		Func() VMFunc
	}

	// VMDateTimer это дата/время
	VMDateTimer interface {
		VMParser
		Time() VMTime
	}

	// VMDurationer это промежуток времени (time.Duration)
	VMDurationer interface {
		VMParser
		Duration() VMTimeDuration
	}

	// VMChanMaker может создать новый канал
	VMChanMaker interface {
		VMInterfacer
		MakeChan(int) VMChaner //размер
	}

	// VMMetaObject реализует поведение системной функциональной структуры (объекта метаданных)
	// реализация должна быть в виде обертки над структурным типом на языке Го
	// обертка получается через встраивание базовой структуры VMMetaObj
	VMMetaObject interface {
		VMInterfacer
		VMCacheMembers(VMMetaObject) // создает внутренние хранилища полей и методов,
		// содержащие id строки с именем (независимое от регистра букв)
		// и индекс среди полей или методов, для получения через рефлексию
		VMIsField(int) bool
		VMGetField(int) VMInterfacer
		VMSetField(int, VMInterfacer)
		VMGetMethod(int) (VMFunc, bool) // получает обертку метода
	}

	// VMNullable означает значение null
	VMNullable interface {
		VMStringer
		null()
	}
)