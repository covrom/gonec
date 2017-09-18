// Package core implements core interface for gonec script.
package core

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

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

/////////////////
// TttStructTest - тестовая структура для отладки работы с системными функциональными структурами
type TttStructTest struct {
	VMMetaObj

	ПолеЦелоеЧисло VMInt
	ПолеСтрока     VMString
}

// обратите внимание - русскоязычное название метода для структуры
func (tst *TttStructTest) ВСтроку() VMString {
	return VMString(fmt.Sprintf("ПолеЦелоеЧисло=%v, ПолеСтрока=%v", tst.ПолеЦелоеЧисло, tst.ПолеСтрока))
}

/////////////////

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
			rets.Append(Now().Вычесть(rv.Time()))
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

	env.DefineS("нрег", func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return strings.ToLower(string(b))
		}
		return strings.ToLower(fmt.Sprint(v))
	})

	env.DefineS("врег", func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return strings.ToUpper(string(b))
		}
		return strings.ToUpper(fmt.Sprint(v))
	})

	env.DefineS("формат", env.Sprintf)

	env.DefineS("кодсимвола", func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	})

	env.DefineS("вбайты", func(s string) []byte {
		return []byte(s)
	})

	env.DefineS("вруны", func(s string) []rune {
		return []rune(s)
	})

	env.DefineS("типзнч", func(v interface{}) string {
		if v == nil {
			return "Неопределено"
		}
		return envir.UniqueNames.Get(env.TypeName(reflect.TypeOf(v)))
	})

	env.DefineS("присвоенозначение", func(s string) bool {
		_, err := env.Get(envir.UniqueNames.Set(s))
		return err == nil
	})

	env.DefineS("паника", func(e interface{}) {
		// os.Setenv("GONEC_DEBUG", "1")
		panic(e)
	})

	env.DefineS("вывести", env.Print)
	env.DefineS("сообщить", env.Println)
	env.DefineS("сообщитьф", env.Printf)
	env.DefineS("stdout", env.StdOut())
	env.DefineS("закрыть", func(e interface{}) {
		reflect.ValueOf(e).Close()
	})
	env.DefineS("обработатьгорутины", runtime.Gosched)

	env.DefineTypeS("целоечисло", int64(0))
	env.DefineTypeS("число", float64(0.0))
	env.DefineTypeS("булево", true)
	env.DefineTypeS("строка", "")
	env.DefineTypeS("массив", VMSlice{})
	env.DefineTypeS("структура", VMStringMap{})
	env.DefineTypeS("дата", VMTime{})

	//////////////////
	env.DefineTypeS("__функциональнаяструктуратест__", TttStructTest{})
	env.DefineS("__дамп__", env.Dump)
	/////////////////////

	return env
}
