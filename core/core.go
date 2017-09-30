// Package core implements core interface for gonec script.
package core

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/covrom/gonec/names"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// LoadAllBuiltins is a convenience function that loads all defineSd builtins.
func LoadAllBuiltins(env *Env) {
	Import(env)

	pkgs := map[string]func(env *Env) *Env{
	// "sort":          gonec_sort.Import,
	// "strings":       gonec_strings.Import,
	}

	env.DefineS("импорт", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должно быть одно название пакета")
		}
		if s, ok := args[0].(VMString); ok {
			if loader, ok := pkgs[strings.ToLower(string(s))]; ok {
				rets.Append(loader(env)) // возвращает окружение, инициализированное пакетом
				return nil
			}
			return fmt.Errorf("Пакет '%s' не найден", s)
		} else {
			return errors.New("Название пакета должно быть строкой")
		}
	}))

	// успешно загружен глобальный контекст
	env.SetBuiltsIsLoaded()
}

// Import общая стандартная бибилиотека
func Import(env *Env) *Env {

	env.DefineS("длина", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if rv, ok := args[0].(VMIndexer); ok {
			rets.Append(rv.Length())
			return nil
		}
		return errors.New("Аргумент должен иметь длину")
	}))

	env.DefineS("ключи", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if vv, ok := args[0].(VMStringMaper); ok {
			rv := vv.StringMap()
			keys := make(VMSlice, len(rv))
			i := 0
			for k := range rv {
				keys[i] = VMString(k)
				i++
			}
			keys.SortDefault()
			rets.Append(keys)
			return nil
		}
		return errors.New("Аргумент должен быть структурой")
	}))

	env.DefineS("диапазон", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) < 1 {
			return errors.New("Отсутствуют аргументы")
		}
		if len(args) > 2 {
			return errors.New("Должна быть длина диапазона или начало и конец")
		}
		var min, max int64
		var arr VMSlice
		if len(args) == 1 {
			min = 0
			maxvm, ok := args[0].(VMInt)
			if !ok {
				return errors.New("Длина диапазона должна быть целым числом")
			}
			max = maxvm.Int() - 1
		} else {
			minvm, ok := args[0].(VMInt)
			if !ok {
				return errors.New("Начало диапазона должно быть целым числом")
			}
			min = minvm.Int()
			maxvm, ok := args[1].(VMInt)
			if !ok {
				return errors.New("Конец диапазона должен быть целым числом")
			}
			max = maxvm.Int()
		}
		if min > max {
			return errors.New("Длина диапазона должна быть больше нуля")
		}
		arr = make(VMSlice, max-min+1)

		for i := min; i <= max; i++ {
			arr[i-min] = VMInt(i)
		}
		rets.Append(arr)
		return nil
	}))

	env.DefineS("текущаядата", VMFunc(func(args VMSlice, rets *VMSlice) error {
		rets.Append(Now())
		return nil
	}))

	env.DefineS("прошловременис", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if rv, ok := args[0].(VMDateTimer); ok {
			rets.Append(Now().Sub(rv.Time()))
			return nil
		}
		return errors.New("Допустим только аргумент типа Дата или совместимый с ним")
	}))

	env.DefineS("пауза", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if v, ok := args[0].(VMNumberer); ok {
			sec1 := decimal.New(int64(VMSecond), 0)
			time.Sleep(time.Duration(v.Decimal().Mul(VMDecimal(sec1)).Int()))
			return nil
		}
		return errors.New("Должно быть число секунд (допустимо с дробной частью)")
	}))

	// TODO: добавить операции по умножению длительностей VMTimeDuration на числа VMNumberer.Decimal

	env.DefineS("длительностьнаносекунды", VMNanosecond)
	env.DefineS("длительностьмикросекунды", VMMicrosecond)
	env.DefineS("длительностьмиллисекунды", VMMillisecond)
	env.DefineS("длительностьсекунды", VMSecond)
	env.DefineS("длительностьминуты", VMMinute)
	env.DefineS("длительностьчаса", VMHour)
	env.DefineS("длительностьдня", VMDay)

	env.DefineS("хэш", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if v, ok := args[0].(VMHasher); ok {
			rets.Append(v.Hash())
			return nil
		}
		return errors.New("Параметр не может быть хэширован")
	}))

	env.DefineS("уникальныйидентификатор", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 0 {
			return errors.New("Параметры не требуются")
		}
		rets.Append(VMString(uuid.NewV1().String()))
		return nil
	}))

	env.DefineS("случайнаястрока", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if v, ok := args[0].(VMInt); ok {
			rets.Append(VMString(MustGenerateRandomString(int(v))))
			return nil
		}
		return errors.New("Параметр должен быть целым числом")
	}))

	env.DefineS("нрег", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if v, ok := args[0].(VMStringer); ok {
			rets.Append(VMString(strings.ToLower(string(v.String()))))
			return nil
		}
		return errors.New("Должен быть параметр-строка")
	}))

	env.DefineS("врег", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if v, ok := args[0].(VMStringer); ok {
			rets.Append(VMString(strings.ToUpper(string(v.String()))))
			return nil
		}
		return errors.New("Должен быть параметр-строка")
	}))

	env.DefineS("формат", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) < 2 {
			return errors.New("Должны быть форматная строка и хотя бы один параметр")
		}
		if v, ok := args[0].(VMString); ok {
			as := VMSlice(args[1:]).Args()
			rets.Append(VMString(env.Sprintf(string(v), as...)))
			return nil
		}
		return errors.New("Форматная строка должна быть строкой")
	}))

	env.DefineS("кодсимвола", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}

		if v, ok := args[0].(VMStringer); ok {
			s := v.String()
			if len(s) == 0 {
				rets.Append(VMInt(0))
			} else {
				rets.Append(VMInt([]rune(s)[0]))
			}
			return nil
		}
		return errors.New("Должен быть параметр-строка")
	}))

	env.DefineS("типзнч", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) != 1 {
			return errors.New("Должен быть один параметр")
		}
		if args[0] == nil || args[0] == VMNil {
			rets.Append(VMString("Неопределено"))
			return nil
		}
		rets.Append(VMString(names.UniqueNames.Get(env.TypeName(reflect.TypeOf(args[0])))))
		return nil
	}))

	env.DefineS("сообщить", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) == 0 {
			env.Println()
			return nil
		}
		as := args.Args()
		env.Println(as...)
		return nil
	}))

	env.DefineS("сообщитьф", VMFunc(func(args VMSlice, rets *VMSlice) error {
		if len(args) < 2 {
			return errors.New("Должны быть форматная строка и хотя бы один параметр")
		}
		if v, ok := args[0].(VMString); ok {
			as := VMSlice(args[1:]).Args()
			env.Printf(string(v), as...)
			return nil
		}
		return errors.New("Форматная строка должна быть строкой")

	}))

	env.DefineS("обработатьгорутины", VMFunc(func(args VMSlice, rets *VMSlice) error {
		runtime.Gosched()
		return nil
	}))

	// при изменении состава типов не забывать изменять их и в lexer.go
	env.DefineTypeS("целоечисло", ReflectVMInt)
	env.DefineTypeS("число", ReflectVMDecimal)
	env.DefineTypeS("булево", ReflectVMBool)
	env.DefineTypeS("строка", ReflectVMString)
	env.DefineTypeS("массив", ReflectVMSlice)
	env.DefineTypeS("структура", ReflectVMStringMap)
	env.DefineTypeS("дата", ReflectVMTime)
	env.DefineTypeS("длительность", ReflectVMTimeDuration)

	//////////////////
	env.DefineTypeS("__функциональнаяструктуратест__", reflect.TypeOf(TttStructTest{}))
	env.DefineS("__дамп__", VMFunc(func(args VMSlice, rets *VMSlice) error {
		env.Dump()
		return nil
	}))
	/////////////////////

	return env
}

/////////////////
// TttStructTest - тестовая структура для отладки работы с системными функциональными структурами
type TttStructTest struct {
	VMMetaObj

	ПолеЦелоеЧисло VMInt
	ПолеСтрока     VMString
}

func (tst *TttStructTest) VMRegister() {
	tst.VMRegisterMethod("ВСтроку", tst.ВСтроку)
	tst.VMRegisterField("ПолеЦелоеЧисло", &tst.ПолеЦелоеЧисло)
	tst.VMRegisterField("ПолеСтрока", &tst.ПолеСтрока)
}

// обратите внимание - русскоязычное название метода для структуры и формат для быстрого вызова
func (tst *TttStructTest) ВСтроку(args VMSlice, rets *VMSlice) error {
	rets.Append(VMString(fmt.Sprintf("ПолеЦелоеЧисло=%v, ПолеСтрока=%v", tst.ПолеЦелоеЧисло, tst.ПолеСтрока)))
	return nil
}

/////////////////
