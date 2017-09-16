// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/builtins"
	envir "github.com/covrom/gonec/env"
	"github.com/covrom/gonec/parser"
)

func Interrupt(env *envir.Env) {
	env.Interrupt()
}

// ParserSrc provides way to parse the code from source.
func ParseSrc(src string) (prs ast.Stmts, bin binstmt.BinCode, err error) {
	defer func() {
		// если это не паника из кода языка
		// if os.Getenv("GONEC_DEBUG") == "" {
		// обрабатываем панику, которая могла возникнуть в вызванной функции
		if ex := recover(); ex != nil {
			if e, ok := ex.(error); ok {
				err = e
			} else {
				err = errors.New(fmt.Sprint(ex))
			}
		}
		// }
	}()

	// По умолчанию добавляем глобальный модуль "_" в начало, чтобы код без заголовка "модуль" мог успешно исполниться
	// Если будет объявлен модуль в коде, он скроет данное объявление
	src = "Модуль _\n" + src

	scanner := &parser.Scanner{}
	scanner.Init(src)

	prs, err = parser.Parse(scanner)
	if err != nil {
		panic(err)
	}
	// оптимизируем дерево AST
	// свертка констант и нативные значения
	prs = parser.ConstFolding(prs)
	// компиляция в бинарный код
	lid := 0
	bin = prs.BinaryCode(0, &lid)

	return prs, bin, err
}

func Run(stmts binstmt.BinCode, env *envir.Env) (retval core.VMValuer, reterr error) {
	defer func() {
		// если это не паника из кода языка
		// if os.Getenv("GONEC_DEBUG") == "" {
		// обрабатываем панику, которая могла возникнуть в вызванной функции
		if ex := recover(); ex != nil {
			if e, ok := ex.(error); ok {
				reterr = e
			} else {
				reterr = errors.New(fmt.Sprint(ex))
			}
		}
		// }
	}()

	const (
		LastNone uint64 = 0
		LastSet  uint64 = 1 << iota
	)

	flagset := LastNone

	// подготавливаем состояние машины: регистры значений, управляющие регистры
	regs := NewVMRegs(stmts,env)

	// стандартная библиотека - загружаем, если она еще не была загружена в это или в родительское окружение

	if !env.IsBuiltsLoaded() {
		// эту функцию определяем тут, чтобы исключить циклические зависимости пакетов
		env.DefineS("загрузитьивыполнить", func(s string) interface{} {
			body, err := ioutil.ReadFile(s)
			if err != nil {
				panic(err)
			}
			isGNX := strings.HasSuffix(strings.ToLower(s), ".gnx")
			if isGNX {
				bbuf := bytes.NewBuffer(body)
				bins, err := binstmt.ReadBinCode(bbuf)
				if err != nil {
					panic(err)
				}
				// env.Dump()
				rv, err := Run(bins, env)
				// env.Dump()
				if err != nil {
					panic(err)
				}
				return rv
			} else {
				_, bins, err := ParseSrc(string(body))
				if err != nil {
					if pe, ok := err.(*parser.Error); ok {
						pe.Filename = s
						panic(pe)
					}
					panic(err)
				}
				// env.Dump()
				rv, err := Run(bins, env)
				// env.Dump()
				if err != nil {
					panic(err)
				}
				return rv
			}
			return nil
		})

		core.LoadAllBuiltins(env)
	}

	goschedidx := 0

	var (
		catcherr error
		idx      int
	)

	for idx < len(stmts.Code) {
		goschedidx++
		if goschedidx == 1000 {
			// даем воздуха другим горутинам каждые 1000 инструкций
			runtime.Gosched()
			goschedidx = 0
		}

		if env.CheckInterrupt() {
			// проверяем, был ли прерван интерпретатор
			return nil, binstmt.InterruptError
		}

		stmt := stmts.Code[idx]
		switch s := stmt.(type) {

		case *binstmt.BinJMP:
			idx = regs.Labels[s.JumpTo]
			continue

		case *binstmt.BinJFALSE:
			if ok := ToBool(regs.Reg[s.Reg]); !ok {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *binstmt.BinJTRUE:
			if ok := ToBool(regs.Reg[s.Reg]); ok {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *binstmt.BinLABEL:
			// пропускаем

		case *binstmt.BinLOAD:
			regs.Set(s.Reg, s.Val)

		case *binstmt.BinMV:
			regs.Set(s.RegTo, regs.Reg[s.RegFrom])

		case *binstmt.BinEQUAL:

			regs.Set(s.Reg, Equal(regs.Reg[s.Reg1], regs.Reg[s.Reg2]))

		case *binstmt.BinCASTNUM:
			// ошибки обрабатываем в попытке
			var str core.VMString
			var ok bool
			if str, ok = regs.Reg[s.Reg].(core.VMString); !ok {
				regs.Set(s.Reg, nil)
				catcherr = binstmt.NewStringError(stmt, "Литерал должен быть числом")
				break
			}
			v, err := InvokeNumber(str)
			if err != nil {
				regs.Set(s.Reg, nil)
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			regs.Set(s.Reg, v)

		case *binstmt.BinMAKESLICE:
			regs.Set(s.Reg, make(core.VMSlice, s.Len, s.Cap))

		case *binstmt.BinSETIDX:
			if v, ok := regs.Reg[s.Reg].(core.VMSlice); ok {
				v[s.Index] = regs.Reg[s.RegVal]
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно изменить значение по индексу")
				break
			}
		case *binstmt.BinMAKEMAP:
			regs.Set(s.Reg, make(core.VMStringMap, s.Len))

		case *binstmt.BinSETKEY:
			if v, ok := regs.Reg[s.Reg].(core.VMStringMap); ok {
				v[s.Key] = regs.Reg[s.RegVal]
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно изменить значение по ключу")
				break
			}

		case *binstmt.BinGET:
			v, err := env.Get(s.Id)
			if err != nil {
				catcherr = binstmt.NewStringError(stmt, "Невозможно получить значение")
				break
			}
			if !v.IsValid() {
				regs.Set(s.Reg, nil)
			} else {
				regs.Set(s.Reg, v.Interface())
			}

		case *binstmt.BinSET:
			env.Define(s.Id, regs.Reg[s.Reg])

		case *binstmt.BinSETMEMBER:

			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			rv := refregs.Index(s.RegVal).Elem()

			switch v.Kind() {
			case reflect.Struct:

				////////////////////////////////////////////////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				// TODO:

				// v, err := ast.FieldByNameCI(v, s.Id)
				// if err != nil {
				// 	catcherr = NewError(stmt, err)
				// 	break
				// }
				// if !v.CanSet() {
				// 	catcherr = NewStringError(stmt, "Невозможно установить значение")
				// 	break
				// }
				// v.Set(rv)
			case reflect.Map:
				v.SetMapIndex(reflect.ValueOf(envir.UniqueNames.Get(s.Id)), rv)
			default:
				if !v.CanSet() {
					catcherr = binstmt.NewStringError(stmt, "Невозможно установить значение")
					break
				}
				v.Set(rv)
			}

		case *binstmt.BinSETNAME:
			v, ok := regs.Reg[s.Reg].(core.VMString)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Имя типа должно быть строкой")
				break
			}
			eType := envir.UniqueNames.Set(v)
			regs.Set(s.Reg, eType)

		case *binstmt.BinSETITEM:
			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			i := reflect.Indirect(refregs.Index(s.RegIndex).Elem())
			rv := refregs.Index(s.RegVal).Elem()
			regs.Set(s.RegNeedLet, false)

			switch v.Kind() {

			case reflect.Array, reflect.Slice:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				ii := int(i.Int())
				if ii < 0 {
					ii += v.Len()
				}
				if ii < 0 || ii >= v.Len() {
					catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
					break
				}

				// для элементов массивов и слайсов это работает, а для строк - нет
				vv := v.Index(ii)
				if !vv.CanSet() {
					catcherr = binstmt.NewStringError(stmt, "Невозможно установить значение")
					break
				}
				vv.Set(rv)

			case reflect.String:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				rvs := []rune(rv.String())
				if len(rvs) != 1 {
					catcherr = binstmt.NewStringError(stmt, "Длина присваиваемой строки должна быть ровно один символ")
					break
				}
				r := []rune(v.String())
				vlen := len(r)
				ii := int(i.Int())
				if ii < 0 {
					ii += vlen
				}
				if ii < 0 || ii >= vlen {
					catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
					break
				}
				// заменяем руну
				r[ii] = rvs[0]

				// для строк здесь неадресуемое значение, поэтому, переприсваиваем
				regs.Set(s.Reg, string(r))
				regs.Set(s.RegNeedLet, true)

			case reflect.Map:
				if i.Kind() != reflect.String {
					catcherr = binstmt.NewStringError(stmt, "Ключ должен быть строкой")
					break
				}
				v.SetMapIndex(i, rv)

			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinSETSLICE:
			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			rb := reflect.Indirect(refregs.Index(s.RegBegin).Elem())
			re := reflect.Indirect(refregs.Index(s.RegEnd).Elem())
			rv := refregs.Index(s.RegVal).Elem()
			regs.Set(s.RegNeedLet, false)

			switch v.Kind() {
			case reflect.Array, reflect.Slice:
				vlen := v.Len()
				ii, ij, err := LeftRightBounds(rb, re, vlen)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				if ij < ii {
					catcherr = binstmt.NewStringError(stmt, "Окончание диапазона не может быть раньше его начала")
					break
				}
				vv := v.Slice(ii, ij)
				if vv.Len() != rv.Len() {
					catcherr = binstmt.NewStringError(stmt, "Размер массива должен быть равен ширине диапазона")
					break
				}
				reflect.Copy(vv, rv)
			case reflect.String:
				r, ii, ij, err := StringToRuneSliceAt(v, rb, re)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}

				rvs := []rune(rv.String())
				if len(rvs) != len(r[ii:ij]) {
					catcherr = binstmt.NewStringError(stmt, "Длина строки должна быть равна длине диапазона")
					break
				}

				// заменяем руны
				copy(r[ii:ij], rvs)

				regs.Set(s.Reg, string(r))
				regs.Set(s.RegNeedLet, true)

			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinUNARY:
			switch s.Op {
			case '-':
				if x, ok := regs.Reg[s.Reg].(float64); ok {
					regs.Set(s.Reg, -x)
				} else if x, ok := regs.Reg[s.Reg].(int64); ok {
					regs.Set(s.Reg, -x)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Операция применима только к числам")
					break
				}
			case '^':
				if x, ok := regs.Reg[s.Reg].(int64); ok {
					regs.Set(s.Reg, ^x)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Операция применима только к целым числам")
					break
				}
			case '!':
				regs.Set(s.Reg, !ToBool(regs.Reg[s.Reg]))
			default:
				catcherr = binstmt.NewStringError(stmt, "Неизвестный оператор")
				break
			}

		case *binstmt.BinADDRID:
			v, err := env.Get(s.Name)
			if err != nil {
				catcherr = binstmt.NewStringError(stmt, "Невозможно получить значение")
				break
			}
			if !v.CanAddr() {
				catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
				break
			}
			regs.Set(s.Reg, v.Addr().Interface())

		case *binstmt.BinADDRMBR:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = binstmt.NewStringError(stmt, "Значение не найдено")
					break
				}
				if !m.CanAddr() {
					catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
					break
				}
				regs.Set(s.Reg, m.Addr().Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			if !m.CanAddr() {
				catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
				break
			}
			regs.Set(s.Reg, m.Addr().Interface())

		case *binstmt.BinUNREFID:
			v, err := env.Get(s.Name)
			if err != nil {
				catcherr = binstmt.NewStringError(stmt, "Невозможно получить значение")
				break
			}
			if v.Kind() != reflect.Ptr {
				catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
				break
			}
			regs.Set(s.Reg, v.Elem().Interface())

		case *binstmt.BinUNREFMBR:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = binstmt.NewStringError(stmt, "Значение не найдено")
					break
				}
				if m.Kind() != reflect.Ptr {
					catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
					break
				}
				regs.Set(s.Reg, m.Elem().Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			if m.Kind() != reflect.Ptr {
				catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
				break
			}
			regs.Set(s.Reg, m.Elem().Interface())

		case *binstmt.BinGETMEMBER:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = binstmt.NewStringError(stmt, "Значение не найдено")
					break
				}
				regs.Set(s.Reg, m.Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			regs.Set(s.Reg, m.Interface())

		case *binstmt.BinGETIDX:
			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			i := reflect.Indirect(refregs.Index(s.RegIndex).Elem())

			switch v.Kind() {

			case reflect.Array, reflect.Slice:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				ii := int(i.Int())
				if ii < 0 {
					ii += v.Len()
				}
				if ii < 0 || ii >= v.Len() {
					catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
					break
				}

				regs.Set(s.Reg, v.Index(ii).Interface())

			case reflect.String:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				r := []rune(v.String())
				vlen := len(r)
				ii := int(i.Int())
				if ii < 0 {
					ii += vlen
				}
				if ii < 0 || ii >= vlen {
					catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
					break
				}
				regs.Set(s.Reg, string(r[ii]))

			case reflect.Map:
				if i.Kind() != reflect.String {
					catcherr = binstmt.NewStringError(stmt, "Ключ должен быть строкой")
					break
				}
				regs.Set(s.Reg, v.MapIndex(i))

			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinGETSUBSLICE:
			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			rb := reflect.Indirect(refregs.Index(s.RegBegin).Elem())
			re := reflect.Indirect(refregs.Index(s.RegEnd).Elem())

			switch v.Kind() {
			case reflect.Array, reflect.Slice:
				rv, err := SliceAt(v, rb, re)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				regs.Set(s.Reg, rv)
			case reflect.String:
				rv, err := StringAt(v, rb, re)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				regs.Set(s.Reg, rv)
			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinOPER:
			refregs := reflect.ValueOf(regs.Reg)
			lhsV := refregs.Index(s.RegL).Elem()
			rhsV := refregs.Index(s.RegR).Elem()

			// log.Println("lhsV", lhsV)
			// log.Println("rhsV", rhsV)

			r, err := EvalBinOp(s.Op, lhsV, rhsV)

			// log.Println("r", r)

			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			regs.Set(s.RegL, r)

		case *binstmt.BinCALL:

			// TODO: сохранять все текущее состояние в стэке, включая набор меток перехода, т.к. в функциях модулей они могут повторяться

			var err error

			//функцию на языке Гонец можно вызывать прямо с аргументами из слайса в регистре
			var fgnc, argssl interface{}
			if s.Name == 0 {
				fgnc = regs.Reg[s.RegArgs]
				argssl = regs.Reg[s.RegArgs+1]
			} else {
				var fgncv reflect.Value
				fgncv, err = env.Get(s.Name)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				fgnc = fgncv.Interface()
				argssl = regs.Reg[s.RegArgs]
			}
			if fnc, ok := fgnc.(core.VMFunc); ok {
				// если ее надо вызвать в горутине - вызываем
				if s.Go {
					env.SetGoRunned(true)
					go fnc(argssl.(core.VMSlice)...)
					regs.Set(s.RegRets, nil)
					break
				}

				// не в горутине
				ret, err := fnc(argssl.(core.VMSlice)...)

				// TODO: проверить, если был передан слайс, и он изменен внутри функции, то что происходит в исходном слайсе?
				// и аналогично проверить значения в переданных указателях

				if err != nil {
					// ошибку передаем в блок обработки исключений
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				regs.Set(s.RegRets, ret)
				break
			}

			// получаем функцию на языке Го
			f := reflect.ValueOf(fgnc)
			vargs := reflect.ValueOf(argssl)
			// это не функция - тогда ошибка
			if f.Kind() != reflect.Func {
				catcherr = binstmt.NewStringError(stmt, "Не является функцией")
				break
			}
			ftype := f.Type()

			// готовим аргументы для вызываемой функции
			args := make([]reflect.Value, 0, s.NumArgs)

			for i := 0; i < s.NumArgs; i++ {
				// очередной аргумент
				arg := vargs.Index(i).Elem()
				// if arg.Kind() == reflect.Interface || arg.Kind() == reflect.Ptr {
				// 		arg = arg.Elem()
				// }
				// конвертируем параметр в целевой тип
				if i < ftype.NumIn() {
					// это функция с постоянным числом аргументов
					if !ftype.IsVariadic() {
						// целевой тип аргумента
						it := ftype.In(i)
						if arg.Kind().String() == "unsafe.Pointer" {
							arg = reflect.New(it).Elem()
						}
						if arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it) {
							// типы не равны - пытаемся конвертировать
							arg = arg.Convert(it)

						} else if arg.Kind() == reflect.Func {
							if _, isFunc := arg.Interface().(core.VMFunc); isFunc {
								// это функция на языке Гонец (т.е. обработчик) - делаем обертку в целевую функцию типа it
								rfunc := arg
								if s.Go {
									env.SetGoRunned(true)
								}
								arg = reflect.MakeFunc(it, func(args []reflect.Value) []reflect.Value {
									// for i := range args {
									// 	args[i] = reflect.ValueOf(args[i])
									// }
									if s.Go {
										go func() {
											rfunc.Call(args)
										}()
										return []reflect.Value{}
									}
									// var rets []reflect.Value
									// for _, v := range rfunc.Call(args)[:it.NumOut()] {
									// 	rets = append(rets, v.Interface().(reflect.Value))
									// }
									// return rets
									return rfunc.Call(args)[:it.NumOut()]
								})
							}
						} else if !arg.IsValid() {
							arg = reflect.Zero(it)
						}
					}
				}
				if !arg.IsValid() {
					arg = envir.NilValue
				}
				if s.VarArg && i == s.NumArgs-1 {
					for j := 0; j < arg.Len(); j++ {
						args = append(args, arg.Index(j))
					}
				} else {
					args = append(args, arg)
				}
			}

			// вызываем функцию

			fnc := func() (ret interface{}, err error) {
				defer func() {
					// если не было прерывания Interrupt()
					// if os.Getenv("GONEC_DEBUG") == "" {
					// обрабатываем панику, которая могла возникнуть в вызванной функции
					if ex := recover(); ex != nil {
						if e, ok := ex.(error); ok {
							err = e
						} else {
							err = errors.New(fmt.Sprint(ex))
						}
					}
					// }
				}()
				// if f.Kind() == reflect.Interface {
				// 	f = f.Elem()
				// }
				rets := f.Call(args)

				// возврат из функций на языке Го

				if f.Type().NumOut() == 1 {
					return rets[0].Interface(), nil // одно значение
				} else {
					var result []interface{}
					for _, r := range rets {
						result = append(result, r.Interface())
					}
					return core.VMSlice(result), nil // массив возвращаемых значений
				}
			}

			// если ее надо вызвать в горутине - вызываем
			if s.Go {
				env.SetGoRunned(true)
				go fnc()
				regs.Set(s.RegRets, nil)
				break
			}

			// не в горутине
			ret, err := fnc()

			// TODO: проверить, если был передан слайс, и он изменен внутри функции, то что происходит в исходном слайсе?
			// и аналогично проверить значения в переданных указателях

			if err != nil {
				// ошибку передаем в блок обработки исключений
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			regs.Set(s.RegRets, ret)

		case *binstmt.BinFUNC:
			f := func(expr *BinFUNC, env *envir.Env) core.VMFunc {
				return func(args ...interface{}) (interface{}, error) {
					if !expr.VarArg {
						if len(args) != len(expr.Args) {
							return nil, binstmt.NewStringError(expr, "Неверное количество аргументов")
						}
					}
					var newenv *envir.Env
					if expr.Name == 0 {
						// наследуем от окружения текущей функции
						newenv = env.NewSubEnv()
					} else {
						// наследуем от модуля или глобального окружения
						newenv = env.NewEnv()
					}

					if expr.VarArg {
						newenv.Define(expr.Args[0], args)
					} else {
						for i, arg := range expr.Args {
							newenv.Define(arg, args[i])
						}
					}
					rr, err := Run(expr.Code, newenv)
					if err == binstmt.ReturnError {
						err = nil
					}
					// TODO: проверить при единичном и множественном возврате, при "..." аргументах
					return rr, err
				}
			}(s, env)
			env.Define(s.Name, f)
			regs.Set(s.Reg, f)

		case *binstmt.BinCASTTYPE:
			// приведение типов, включая приведение типов в массиве как новый типизированный массив
			eType, ok := regs.Reg[s.TypeReg].(int)
			if !ok {
				catcherr = NewStringError(stmt, "Неизвестный тип")
				break
			}
			nt, err := env.Type(eType)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			rv := regs.Reg[s.Reg]
			v, err := TypeCastConvert(rv, nt, false)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}

			regs.Set(s.Reg, v)

		case *binstmt.BinMAKE:
			eType, ok := regs.Reg[s.Reg].(int)
			if !ok {
				catcherr = NewStringError(stmt, "Неизвестный тип")
				break
			}
			rt, err := env.Type(eType)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			var v reflect.Value
			if rt.Kind() == reflect.Map {
				v = reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt)
			} else if rt.Kind() == reflect.Struct {
				// структуру создаем всегда ссылочной
				// иначе не работает присвоение полей через рефлексию
				v = reflect.New(rt)
			} else {
				v = reflect.Zero(rt)
			}
			regs.Set(s.Reg, v.Interface())

		case *binstmt.BinMAKECHAN:
			size, ok := regs.Reg[s.Reg].(int64)
			if !ok {
				catcherr = NewStringError(stmt, "Размер должен быть целым числом")
				break
			}
			v := make(core.VMChannel, size)
			regs.Set(s.Reg, v)

		case *binstmt.BinMAKEARR:
			alen := int(ToInt64(regs.Reg[s.Reg]))
			acap := int(ToInt64(regs.Reg[s.RegCap]))
			v := make(core.VMSlice, alen, acap)
			regs.Set(s.Reg, v)

		case *binstmt.BinCHANRECV:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			v, _ := ch.Recv()
			regs.Set(s.RegVal, v.Interface())

		case *binstmt.BinCHANSEND:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			v := regs.Reg[s.RegVal]
			ch.Send(reflect.ValueOf(v))

		case *binstmt.BinISKIND:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			regs.Set(s.Reg, v.Kind() == s.Kind)

		case *binstmt.BinISSLICE:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			regs.Set(s.RegBool, v.Kind() == reflect.Array || v.Kind() == reflect.Slice)

		case *binstmt.BinINC:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			var x interface{}
			if v.Kind() == reflect.Float64 {
				x = ToFloat64(v.Interface()) + 1.0
			} else {
				x = ToInt64(v.Interface()) + 1
			}
			regs.Set(s.Reg, x)

		case *binstmt.BinDEC:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			var x interface{}
			if v.Kind() == reflect.Float64 {
				x = ToFloat64(v.Interface()) - 1.0
			} else {
				x = ToInt64(v.Interface()) - 1
			}
			regs.Set(s.Reg, x)

		case *binstmt.BinTRY:
			regs.PushTry(s.Reg, s.JumpTo)
			regs.Set(s.Reg, nil) // изначально ошибки нет

		case *binstmt.BinCATCH:
			// получаем ошибку, и если ее нет, переходим на метку, иначе, выполняем дальше
			nerr := regs.Reg[s.Reg]
			if nerr == nil {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *binstmt.BinPOPTRY:
			// если catch блок отработал, то стек уже очищен, иначе снимаем со стека (ошибок не было)
			if regs.TopTryLabel() == s.CatchLabel {
				regs.PopTry()
			}

		case *binstmt.BinFOREACH:
			val := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()

			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				regs.Set(s.RegIter, int(-1))
			case reflect.Chan:
				regs.Set(s.RegIter, nil)
			default:
				catcherr = NewStringError(stmt, "Не является коллекцией или каналом")
				break
			}
			if catcherr != nil {
				break
			}

			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *binstmt.BinNEXT:
			val := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()

			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				iter := ToInt64(regs.Reg[s.RegIter])
				iter++
				if iter < int64(val.Len()) {
					regs.Set(s.RegIter, iter)
					iv := val.Index(int(iter))
					regs.Set(s.RegVal, iv)
				} else {
					idx = regs.Labels[s.JumpTo]
					continue
				}
			case reflect.Chan:
				iv, ok := val.Recv()
				if !ok {
					catcherr = NewStringError(stmt, "Канал был закрыт")
					break
				}
				regs.Set(s.RegVal, iv.Interface())
			default:
				catcherr = NewStringError(stmt, "Не является коллекцией или каналом")
				break
			}

		case *binstmt.BinPOPFOR:
			if regs.TopContinue() == s.ContinueLabel {
				regs.PopContinue()
				regs.PopBreak()
			}

		case *binstmt.BinFORNUM:

			if !IsNum(regs.Reg[s.RegFrom]) {
				catcherr = NewStringError(stmt, "Начальное значение должно быть целым числом")
				break
			}
			if !IsNum(regs.Reg[s.RegTo]) {
				catcherr = NewStringError(stmt, "Конечное значение должно быть целым числом")
				break
			}

			regs.Set(s.Reg, nil)
			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *binstmt.BinNEXTNUM:
			afrom := ToInt64(regs.Reg[s.RegFrom])
			ato := ToInt64(regs.Reg[s.RegTo])
			fviadd := int64(1)
			if afrom > ato {
				fviadd = int64(-1) // если конечное значение меньше первого, идем в обратном порядке
			}
			vv := regs.Reg[s.Reg]
			var iter int64
			if vv == nil {
				iter = afrom
			} else {
				iter = ToInt64(vv)
				iter += fviadd
			}
			inrange := iter <= ato
			if afrom > ato {
				inrange = iter >= ato
			}
			if inrange {
				regs.Set(s.Reg, iter)
			} else {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *binstmt.BinWHILE:
			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *binstmt.BinRET:
			retval = regs.Reg[s.Reg]
			return retval, ReturnError

		case *binstmt.BinTHROW:
			catcherr = NewStringError(stmt, fmt.Sprint(regs.Reg[s.Reg]))
			break

		case *binstmt.BinMODULE:
			// модуль регистрируется в глобальном контексте
			newenv := env.NewModule(envir.UniqueNames.Get(s.Name))
			_, err := Run(s.Code, newenv) // инициируем модуль
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}

		case *binstmt.BinERROR:
			// необрабатываемая в попытке ошибка
			return retval, NewStringError(s, s.Error)

		case *binstmt.BinBREAK:
			label := regs.PopBreak()
			if label != -1 {
				regs.PopContinue()
				idx = regs.Labels[label]
				continue
			}
			return nil, BreakError

		case *binstmt.BinCONTINUE:
			label := regs.PopContinue()
			if label != -1 {
				regs.PopBreak()
				idx = regs.Labels[label]
				continue
			}
			return nil, ContinueError

		case *binstmt.BinTRYRECV:

			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			v, ok := ch.TryRecv()
			if !v.IsValid() {
				regs.Set(s.RegVal, nil)
				regs.Set(s.RegOk, ok)
				regs.Set(s.RegClosed, true)
			} else {
				regs.Set(s.RegVal, v.Interface())
				regs.Set(s.RegOk, ok)
				regs.Set(s.RegClosed, false)
			}

		case *binstmt.BinTRYSEND:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			ok := ch.TrySend(reflect.ValueOf(regs.Reg).Index(s.RegVal).Elem())
			regs.Set(s.RegOk, ok)

		case *binstmt.BinGOSHED:
			runtime.Gosched()

		case *binstmt.BinFREE:
			regs.FreeFromReg(s.Reg)

		default:
			return nil, binstmt.NewStringError(stmt, "Неизвестная инструкция")
		}

		if catcherr != nil {
			nerr := binstmt.NewError(stmt, catcherr)
			catcherr = nil
			// учитываем стек обработки ошибок
			if regs.TopTryLabel() == -1 {
				return nil, nerr
			} else {
				env.DefineS("описаниеошибки", func(s string) CatchFunc {
					return func() string { return s }
				}(nerr.Error()))

				r, idxl := regs.PopTry()
				regs.Set(r, nerr)
				idx = regs.Labels[idxl] // переходим в catch блок, функция с описанием ошибки определена
				continue
			}
		}

		idx++
	}
	return retval, nil
}
