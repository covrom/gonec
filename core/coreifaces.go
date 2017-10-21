package core

import (
	"encoding"
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

	// VMOperationer может выполнить операцию с другим значением, операцию сравнения или математическую
	VMOperationer interface {
		VMValuer
		EvalBinOp(VMOperation, VMOperationer) (VMValuer, error) // возвращает результат выражения с другим значением
	}

	// VMUnarer может выполнить унарную операцию над свои значением
	VMUnarer interface {
		VMValuer
		EvalUnOp(rune) (VMValuer, error) // возвращает результат выражения с другим значением
	}

	// VMConverter может конвертироваться в тип reflect.Type
	VMConverter interface {
		VMValuer
		ConvertToType(t reflect.Type) (VMValuer, error)
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

	// VMBinaryTyper может сериализовываться в бинарные данные внутри слайсов и структур
	VMBinaryTyper interface {
		VMValuer
		encoding.BinaryMarshaler
		BinaryType() VMBinaryType
	}

	// конкретные типы виртуальной машины

	// VMStringer строка
	VMStringer interface {
		VMInterfacer
		String() string
	}

	// VMNumberer число, внутреннее хранение в int64 или decimal формате
	VMNumberer interface {
		VMInterfacer
		Int() int64
		Float() float64
		DecNum() VMDecNum
		InvokeNumber() (VMNumberer, error) // извлекает VMInt или VMDecNum, в зависимости от наличия .eE
	}

	// VMBooler сообщает значение булево
	VMBooler interface {
		VMInterfacer
		Bool() bool
	}

	// VMSlicer может быть представлен в виде слайса Гонец
	VMSlicer interface {
		VMInterfacer
		Slice() VMSlice
	}

	// VMStringMaper может быть представлен в виде структуры Гонец
	VMStringMaper interface {
		VMInterfacer
		StringMap() VMStringMap
	}

	// VMFuncer это функция Гонец
	VMFuncer interface {
		VMInterfacer
		Func() VMFunc
	}

	// VMDateTimer это дата/время
	VMDateTimer interface {
		VMInterfacer
		Time() VMTime
	}

	// VMHasher возвращает хэш значения по алгоритму SipHash-2-4 в виде hex-строки
	VMHasher interface {
		VMInterfacer
		Hash() VMString
	}

	// VMDurationer это промежуток времени (time.Duration)
	VMDurationer interface {
		VMInterfacer
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
		VMInterfacer         // реализовано в VMMetaObj
		VMInit(VMMetaObject) // реализовано в VMMetaObj

		// !!!эта функция должна быть обязательно реализована в конечном объекте!!!
		VMRegister()

		VMRegisterMethod(string, VMMethod) // реализовано в VMMetaObj
		VMRegisterField(string, VMValuer)  // реализовано в VMMetaObj

		VMIsField(int) bool             // реализовано в VMMetaObj
		VMGetField(int) VMValuer        // реализовано в VMMetaObj
		VMSetField(int, VMValuer)       // реализовано в VMMetaObj
		VMGetMethod(int) (VMFunc, bool) // реализовано в VMMetaObj
	}

	// VMMethodImplementer реализует только методы, доступные в языке Гонец
	VMMethodImplementer interface{
		VMValuer
		MethodMember(int) (VMFunc, bool) // возвращает метод в нужном формате		
	}

	// VMServicer определяет микросервис, который может регистрироваться в главном менеджере сервисов
	VMServicer interface{
		VMValuer
		Header() VMServiceHeader
		Start() error // запускает горутину, и если не стартовал, возвращает ошибку
		HealthCheck() error // если не живой, то возвращает ошибку
		Stop() error // последняя ошибка при остановке
	}
		
	// VMNullable означает значение null
	VMNullable interface {
		VMStringer
		null()
	}
)
