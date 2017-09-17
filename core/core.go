// Package core implements core interface for gonec script.
package core

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// LoadAllBuiltins is a convenience function that loads all defineSd builtins.
func LoadAllBuiltins(env *Env) {
	Import(env)

	pkgs := map[string]func(env *Env) *Env{
	// "sort":          gonec_sort.Import,
	// "strings":       gonec_strings.Import,
	}

	env.DefineS("импорт", VMFunc(func(args VMSlicer) (VMValuer, error) {
		as := args.Slice()
		if len(as) != 1 {
			return nil, errors.New("Должно быть одно название пакета")
		}
		if s, ok := as[0].(VMString); ok {
			if loader, ok := pkgs[strings.ToLower(string(s))]; ok {
				m := loader(env)
				return m, nil // возвращает окружение, инициализированное пакетом
			}
			return nil, fmt.Errorf("Пакет '%s' не найден", s)
		} else {
			return nil, errors.New("Название пакета должно быть строкой")
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

	env.DefineS("длина", VMFunc(func(args VMSlicer) (VMValuer, error) {
		as := args.Slice()
		if len(as) != 1 {
			return nil, errors.New("Должен быть один параметр")
		}
		if rv, ok := as[0].(VMIndexer); ok {
			return rv.Length(), nil
		}
		return nil, errors.New("Аргумент должен иметь длину")
	}))

	env.DefineS("ключи", VMFunc(func(args VMSlicer) (VMValuer, error) {
		as := args.Slice()
		if len(as) != 1 {
			return nil, errors.New("Должен быть один параметр")
		}
		if vv, ok := as[0].(VMStringMaper); ok {
			rv := vv.StringMap()
			keys := make(VMSlice, len(rv))
			i := 0
			for k := range rv {
				keys[i] = VMString(k)
				i++
			}
			keys.SortDefault()
			return keys, nil
		}
		return nil, errors.New("Аргумент должен быть структурой")
	}))

	env.DefineS("диапазон", VMFunc(func(args VMSlicer) (VMValuer, error) {
		as := args.Slice()
		if len(as) < 1 {
			return nil, errors.New("Отсутствуют аргументы")
		}
		if len(as) > 2 {
			return nil, errors.New("Должна быть длина диапазона или начало и конец")
		}
		var min, max int64
		var arr VMSlice
		if len(as) == 1 {
			min = 0
			maxvm, ok := as[0].(VMInt)
			if !ok {
				return nil, errors.New("Длина диапазона должна быть целым числом")
			}
			max = maxvm.Int() - 1
		} else {
			minvm, ok := as[0].(VMInt)
			if !ok {
				return nil, errors.New("Начало диапазона должно быть целым числом")
			}
			min = minvm.Int()
			maxvm, ok := as[1].(VMInt)
			if !ok {
				return nil, errors.New("Конец диапазона должен быть целым числом")
			}
			max = maxvm.Int()
		}
		if min > max {
			return nil, errors.New("Длина диапазона должна быть больше нуля")
		}
		arr = make(VMSlice, max-min+1)

		for i := min; i <= max; i++ {
			arr[i-min] = VMInt(i)
		}
		return arr, nil
	}))

	env.DefineS("текущаядата", VMFunc(func(args VMSlicer) (VMValuer, error) {
		return VMTime(time.Now()), nil
	}))

	env.DefineS("прошловременис", VMFunc(func(args VMSlicer) (VMValuer, error) {
		as := args.Slice()
		if len(as) != 1 {
			return nil, errors.New("Должен быть один параметр")
		}
		if rv, ok := as[0].(VMTime); ok {
			return VMTimeDuration(time.Since(time.Time(rv.Time()))), nil
		}
		return nil, errors.New("Допустим только аргумент типа Дата")

	}))

	env.DefineS("пауза", time.Sleep)

	env.DefineS("выполнитьспустя", time.AfterFunc)

	env.DefineS("длительностьнаносекунды", time.Nanosecond)
	env.DefineS("длительностьмикросекунды", time.Microsecond)
	env.DefineS("длительностьмиллисекунды", time.Millisecond)
	env.DefineS("длительностьсекунды", time.Second)
	env.DefineS("длительностьминуты", time.Minute)
	env.DefineS("длительностьчаса", time.Hour)
	env.DefineS("длительностьдня", VMTimeDuration(time.Hour*24))
	env.DefineS("длительность", func(s string) VMTimeDuration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(err)
		}
		return d
	})

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
