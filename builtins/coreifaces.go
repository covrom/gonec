package core

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

	// VMParser может парсить из строки
	VMParser interface {
		VMValuer
		Parse(string) error // используется для указателей, т.к. парсит в их значения
	}

	// VMChaner реализует поведение канала
	VMChaner interface {
		VMValuer
		Send(VMValuer)
		Recv() VMValuer
		TrySend(VMValuer) bool
		TryRecv() (VMValuer, bool)
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
		Decimal() VMDecimal
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

	// VMChanMaker может создать новый канал
	VMChanMaker interface {
		VMInterfacer
		MakeChan(int) VMChaner //размер
	}

	// VMMetaStructer реализует поведение системной функциональной структуры (объекта метаданных)
	// реализация должна быть в виде обертки над структурным типом на языке Го
	// обертка получается через встраивание базовой структуры VMMetaObj
	VMMetaStructer interface {
		VMInterfacer
		CacheMembers() // создает внутренние хранилища полей и методов,
		// содержащие id строки с именем (независимое от регистра букв)
		// и индекс среди полей или методов, для получения через рефлексию
		IsField(int) bool
		GetField(int) VMInterfacer
		SetField(int, VMInterfacer)
		GetMethod(int) VMMeth // получает обертку метода
	}
)

// VMMetaObj корневая структура для системных функциональных структур Го, доступных из языка Гонец
type VMMetaObj struct {
}

// TODO: доделать реализацию

func (v *VMMetaObj) CacheMembers() {
	panic("TODO")
}

func (v *VMMetaObj) IsField(name int) bool {
	panic("TODO")
}

func (v *VMMetaObj) GetField(name int) VMInterfacer {
	panic("TODO")
}

func (v *VMMetaObj) SetField(name int, val VMInterfacer) {
	panic("TODO")
}

func (v *VMMetaObj) GetMethod(name int) VMMeth {
	panic("TODO")
}
