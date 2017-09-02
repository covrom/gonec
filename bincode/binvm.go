// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/covrom/gonec/ast"
	envir "github.com/covrom/gonec/env"
)

func Interrupt(env *envir.Env) {
	env.Interrupt()
}

func Run(stmts BinCode, env *envir.Env) (retval interface{}, reterr error) {
	// подготавливаем состояние машины: регистры значений, управляющие регистры
	regs := NewVMRegs(stmts)

	goschedidx := 0

	var catcherr error
	var idx int
	for idx < len(stmts) {
		goschedidx++
		if goschedidx == 1000 {
			// даем воздуха другим горутинам каждые 1000 инструкций
			runtime.Gosched()
			goschedidx = 0
		}

		if env.CheckInterrupt() {
			// проверяем, был ли прерван интерпретатор
			return nil, InterruptError
		}

		stmt := stmts[idx]
		switch s := stmt.(type) {

		case *BinJMP:
			idx = regs.Labels[s.JumpTo]
			continue

		case *BinJFALSE:
			if ok := ToBool(regs.Reg[s.Reg]); !ok {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *BinJTRUE:
			if ok := ToBool(regs.Reg[s.Reg]); ok {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *BinLABEL:
			// пропускаем

		case *BinLOAD:
			regs.Set(s.Reg, s.Val)

		case *BinMV:
			regs.Set(s.RegTo, regs.Reg[s.RegFrom])

		case *BinEQUAL:

			regs.Set(s.Reg, Equal(regs.Reg[s.Reg1], regs.Reg[s.Reg2]))

		case *BinCASTNUM:
			// ошибки обрабатываем в попытке
			var str string
			var ok bool
			if str, ok = regs.Reg[s.Reg].(string); !ok {
				regs.Set(s.Reg, nil)
				catcherr = NewStringError(stmt, "Литерал должен быть числом")
				break
			}
			v, err := InvokeNumber(str)
			if err != nil {
				regs.Set(s.Reg, nil)
				catcherr = NewError(stmt, err)
				break
			}
			regs.Set(s.Reg, v)

		case *BinMAKESLICE:
			regs.Set(s.Reg, make(VMSlice, s.Len, s.Cap))

		case *BinSETIDX:
			if v, ok := regs.Reg[s.Reg].(VMSlice); ok {
				v[s.Index] = regs.Reg[s.RegVal]
			} else {
				catcherr = NewStringError(stmt, "Невозможно изменить значение по индексу")
				break
			}
		case *BinMAKEMAP:
			regs.Set(s.Reg, make(VMStringMap, s.Len))

		case *BinSETKEY:
			if v, ok := regs.Reg[s.Reg].(VMStringMap); ok {
				v[s.Key] = regs.Reg[s.RegVal]
			} else {
				catcherr = NewStringError(stmt, "Невозможно изменить значение по ключу")
				break
			}

		case *BinGET:
			v, err := env.Get(s.Id)
			if err != nil {
				catcherr = NewStringError(stmt, "Невозможно получить значение")
				break
			}
			regs.Set(s.Reg, v.Interface())

		case *BinSET:
			env.Define(s.Id, regs.Reg[s.Reg])

		case *BinSETMEMBER:

			// TODO: обработка паники - передача ошибки в catch блок
			refregs := reflect.ValueOf(regs.Reg)
			v := reflect.Indirect(refregs.Index(s.Reg).Elem())
			rv := refregs.Index(s.RegVal).Elem()

			switch v.Kind() {
			case reflect.Struct:
				v, err := ast.FieldByNameCI(v, s.Id)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				if !v.CanSet() {
					catcherr = NewStringError(stmt, "Невозможно установить значение")
					break
				}
				v.Set(rv)
			case reflect.Map:
				v.SetMapIndex(reflect.ValueOf(ast.UniqueNames.Get(s.Id)), rv)
			default:
				if !v.CanSet() {
					catcherr = NewStringError(stmt, "Невозможно установить значение")
					break
				}
				v.Set(rv)
			}

		case *BinSETNAME:
			v, ok := regs.Reg[s.Reg].(string)
			if !ok {
				catcherr = NewStringError(stmt, "Имя типа должно быть строкой")
				break
			}
			eType := ast.UniqueNames.Set(v)
			regs.Set(s.Reg, eType)

		case *BinSETITEM:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			i := refregs.Index(s.RegIndex).Elem()
			rv := refregs.Index(s.RegVal).Elem()
			regs.Set(s.RegNeedLet, false)

			switch v.Kind() {

			case reflect.Array, reflect.Slice:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				ii := int(i.Int())
				if ii < 0 {
					ii += v.Len()
				}
				if ii < 0 || ii >= v.Len() {
					catcherr = NewStringError(stmt, "Индекс за пределами границ")
					break
				}

				// для элементов массивов и слайсов это работает, а для строк - нет
				vv := v.Index(ii)
				if !vv.CanSet() {
					catcherr = NewStringError(stmt, "Невозможно установить значение")
					break
				}
				vv.Set(rv)

			case reflect.String:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				rvs := []rune(rv.String())
				if len(rvs) != 1 {
					catcherr = NewStringError(stmt, "Длина присваиваемой строки должна быть ровно один символ")
					break
				}
				r := []rune(v.String())
				vlen := len(r)
				ii := int(i.Int())
				if ii < 0 {
					ii += vlen
				}
				if ii < 0 || ii >= vlen {
					catcherr = NewStringError(stmt, "Индекс за пределами границ")
					break
				}
				// заменяем руну
				r[ii] = rvs[0]

				// для строк здесь неадресуемое значение, поэтому, переприсваиваем
				regs.Set(s.Reg, string(r))
				regs.Set(s.RegNeedLet, true)

			case reflect.Map:
				if i.Kind() != reflect.String {
					catcherr = NewStringError(stmt, "Ключ должен быть строкой")
					break
				}
				v.SetMapIndex(i, rv)

			default:
				catcherr = NewStringError(stmt, "Неверная операция")
				break
			}

		case *BinSETSLICE:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			rb := refregs.Index(s.RegBegin).Elem()
			re := refregs.Index(s.RegEnd).Elem()
			rv := refregs.Index(s.RegVal).Elem()
			regs.Set(s.RegNeedLet, false)

			switch v.Kind() {
			case reflect.Array, reflect.Slice:
				vlen := v.Len()
				ii, ij, err := ast.LeftRightBounds(rb, re, vlen)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				if ij < ii {
					catcherr = NewStringError(stmt, "Окончание диапазона не может быть раньше его начала")
					break
				}
				vv := v.Slice(ii, ij)
				if vv.Len() != rv.Len() {
					catcherr = NewStringError(stmt, "Размер массива должен быть равен ширине диапазона")
					break
				}
				reflect.Copy(vv, rv)
			case reflect.String:
				r, ii, ij, err := ast.StringToRuneSliceAt(v, rb, re)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}

				rvs := []rune(rv.String())
				if len(rvs) != len(r[ii:ij]) {
					catcherr = NewStringError(stmt, "Длина строки должна быть равна длине диапазона")
					break
				}

				// заменяем руны
				copy(r[ii:ij], rvs)

				regs.Set(s.Reg, string(r))
				regs.Set(s.RegNeedLet, true)

			default:
				catcherr = NewStringError(stmt, "Неверная операция")
				break
			}

		case *BinUNARY:
			switch s.Op {
			case '-':
				if x, ok := regs.Reg[s.Reg].(float64); ok {
					regs.Set(s.Reg, -x)
				} else if x, ok := regs.Reg[s.Reg].(int64); ok {
					regs.Set(s.Reg, -x)
				} else {
					catcherr = NewStringError(stmt, "Операция применима только к числам")
					break
				}
			case '^':
				if x, ok := regs.Reg[s.Reg].(int64); ok {
					regs.Set(s.Reg, ^x)
				} else {
					catcherr = NewStringError(stmt, "Операция применима только к целым числам")
					break
				}
			case '!':
				regs.Set(s.Reg, !ToBool(regs.Reg[s.Reg]))
			default:
				catcherr = NewStringError(stmt, "Неизвестный оператор")
				break
			}

		case *BinADDRID:
			v, err := env.Get(s.Name)
			if err != nil {
				catcherr = NewStringError(stmt, "Невозможно получить значение")
				break
			}
			if !v.CanAddr() {
				catcherr = NewStringError(stmt, "Невозможно получить адрес значения")
				break
			}
			regs.Set(s.Reg, v.Addr().Interface())

		case *BinADDRMBR:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = NewStringError(stmt, "Значение не найдено")
					break
				}
				if !m.CanAddr() {
					catcherr = NewStringError(stmt, "Невозможно получить адрес значения")
					break
				}
				regs.Set(s.Reg, m.Addr().Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			if !m.CanAddr() {
				catcherr = NewStringError(stmt, "Невозможно получить адрес значения")
				break
			}
			regs.Set(s.Reg, m.Addr().Interface())

		case *BinUNREFID:
			v, err := env.Get(s.Name)
			if err != nil {
				catcherr = NewStringError(stmt, "Невозможно получить значение")
				break
			}
			if v.Kind() != reflect.Ptr {
				catcherr = NewStringError(stmt, "Отсутствует ссылка на значение")
				break
			}
			regs.Set(s.Reg, v.Elem().Interface())

		case *BinUNREFMBR:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = NewStringError(stmt, "Значение не найдено")
					break
				}
				if m.Kind() != reflect.Ptr {
					catcherr = NewStringError(stmt, "Отсутствует ссылка на значение")
					break
				}
				regs.Set(s.Reg, m.Elem().Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			if m.Kind() != reflect.Ptr {
				catcherr = NewStringError(stmt, "Отсутствует ссылка на значение")
				break
			}
			regs.Set(s.Reg, m.Elem().Interface())

		case *BinGETMEMBER:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			if vme, ok := v.Interface().(*envir.Env); ok {
				m, err := vme.Get(s.Name)
				if !m.IsValid() || err != nil {
					catcherr = NewStringError(stmt, "Значение не найдено")
					break
				}
				regs.Set(s.Reg, m.Interface())
				break
			}
			m, err := GetMember(v, s.Name, s)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			regs.Set(s.Reg, m.Interface())

		case *BinGETIDX:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			i := refregs.Index(s.RegIndex).Elem()

			switch v.Kind() {

			case reflect.Array, reflect.Slice:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				ii := int(i.Int())
				if ii < 0 {
					ii += v.Len()
				}
				if ii < 0 || ii >= v.Len() {
					catcherr = NewStringError(stmt, "Индекс за пределами границ")
					break
				}

				regs.Set(s.Reg, v.Index(ii).Interface())

			case reflect.String:
				if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
					catcherr = NewStringError(stmt, "Индекс должен быть целым числом")
					break
				}
				r := []rune(v.String())
				vlen := len(r)
				ii := int(i.Int())
				if ii < 0 {
					ii += vlen
				}
				if ii < 0 || ii >= vlen {
					catcherr = NewStringError(stmt, "Индекс за пределами границ")
					break
				}
				regs.Set(s.Reg, string(r[ii]))

			case reflect.Map:
				if i.Kind() != reflect.String {
					catcherr = NewStringError(stmt, "Ключ должен быть строкой")
					break
				}
				regs.Set(s.Reg, v.MapIndex(i))

			default:
				catcherr = NewStringError(stmt, "Неверная операция")
				break
			}

		case *BinGETSUBSLICE:
			refregs := reflect.ValueOf(regs.Reg)
			v := refregs.Index(s.Reg).Elem()
			rb := refregs.Index(s.RegBegin).Elem()
			re := refregs.Index(s.RegEnd).Elem()

			switch v.Kind() {
			case reflect.Array, reflect.Slice:
				rv, err := ast.SliceAt(v, rb, re, envir.NilValue)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				regs.Set(s.Reg, rv.Interface())
			case reflect.String:
				rv, err := ast.StringAt(v, rb, re, envir.NilValue)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				regs.Set(s.Reg, rv.Interface())
			default:
				catcherr = NewStringError(stmt, "Неверная операция")
				break
			}

		case *BinOPER:
			refregs := reflect.ValueOf(regs.Reg)
			lhsV := refregs.Index(s.RegL).Elem()
			rhsV := refregs.Index(s.RegR).Elem()

			// log.Println("lhsV", lhsV)
			// log.Println("rhsV", rhsV)

			r, err := EvalBinOp(s.Op, lhsV, rhsV)

			// log.Println("r", r)

			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}
			regs.Set(s.RegL, r)

		case *BinCALL:
			// получаем функцию
			var f, vargs reflect.Value
			var err error
			if s.Name == 0 {
				// в регистре - функция
				f = reflect.ValueOf(regs.Reg).Index(s.RegArgs).Elem()
				// в следующем - массив аргументов
				vargs = reflect.ValueOf(regs.Reg).Index(s.RegArgs + 1).Elem()
			} else {
				// функция берется из переменной или по имени в окружении
				f, err = env.Get(s.Name)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				// в регистре - массив аргументов
				vargs = reflect.ValueOf(regs.Reg).Index(s.RegArgs).Elem()
			}
			// это не функция - тогда ошибка
			if f.Kind() != reflect.Func {
				catcherr = NewStringError(stmt, "Не является функцией")
				break
			}
			ftype := f.Type()
			// isReflect = это функция на языке Гонец, а не встоенная в стандартную библиотеку
			_, isReflect := f.Interface().(Func)

			// готовим аргументы для вызываемой функции
			args := make([]reflect.Value, 0, s.NumArgs)
			// log.Printf("args-7 %v %T\n", args, args)

			for i := 0; i < s.NumArgs; i++ {
				// очередной аргумент
				arg := vargs.Index(i)
				// конвертируем параметр в целевой тип
				if i < ftype.NumIn() {
					// это функция с постоянным числом аргументов
					if !ftype.IsVariadic() {
						// целевой тип аргумента
						it := ftype.In(i)
						// if arg.Kind().String() == "unsafe.Pointer" {
						// 	arg = reflect.New(it).Elem()
						// }
						if arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it) {
							// типы не равны - пытаемся конвертировать
							arg = arg.Convert(it)
							// log.Printf("arg-1 %#v\n", arg)

						} else if arg.Kind() == reflect.Func {
							if _, isFunc := arg.Interface().(Func); isFunc {
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
								// log.Printf("arg-2 %#v\n", arg)
							}
						} else if !arg.IsValid() {
							arg = reflect.Zero(it)
							// log.Printf("arg-3 %#v\n", arg)
						}
					}
				}
				if !arg.IsValid() {
					arg = envir.NilValue
					// log.Printf("arg-4 %#v\n", arg)
				}
				// log.Printf("arg-5 %#v\n", arg)
				// if !isReflect {
				// 	// для функций на языке Го
				if s.VarArg && i == s.NumArgs-1 {
					// log.Println("arg-6")
					for j := 0; j < arg.Len(); j++ {
						args = append(args, arg.Index(j))
					}
				} else {
					// log.Printf("arg-7 %#v %T\n", arg, arg)
					args = append(args, arg)
					// log.Printf("args-7 %v %T\n", args, args)
				}
				// } else {
				// 	// для функций на языке Гонец
				// 	// if arg.Kind() == reflect.Interface {
				// 	// 	arg = arg.Elem()
				// 	// }
				// 	if s.VarArg && i == s.NumArgs-1 {
				// 		for j := 0; j < arg.Len(); j++ {
				// 			args = append(args, arg.Index(j))
				// 		}
				// 	} else {
				// 		args = append(args, reflect.ValueOf(arg))
				// 	}
				// }

			}
			// log.Printf("%v\n", args)

			// вызываем функцию

			fnc := func() (ret interface{}, err error) {
				defer func() {
					// если не было прерывания Interrupt()
					if os.Getenv("GONEC_DEBUG") == "" {
						// обрабатываем панику, которая могла возникнуть в вызванной функции
						if ex := recover(); ex != nil {
							if e, ok := ex.(error); ok {
								err = e
							} else {
								err = errors.New(fmt.Sprint(ex))
							}
						}
					}
				}()
				// if f.Kind() == reflect.Interface {
				// 	f = f.Elem()
				// }
				rets := f.Call(args)

				if isReflect {
					// возврат из функций на языке Гонец содержит массив возвращенных значений в [0]
					// и возникшую ошибку в [1]
					ev := rets[1].Interface()
					if ev != nil {
						err = ev.(error)
					}
					return rets[0].Interface(), err // массив возвращаемых значений
				} else {

					// возврат из функций на языке Го

					// for i, expr := range e.SubExprs {
					// 	if ae, ok := expr.(*ast.AddrExpr); ok {
					// 		if id, ok := ae.Expr.(*ast.IdentExpr); ok {
					// 			invokeLetExpr(id, args[i].Elem().Elem(), env)
					// 		}
					// 	}
					// }

					if f.Type().NumOut() == 1 {
						return rets[0].Interface(), nil // одно значение
					} else {
						var result []interface{}
						for _, r := range rets {
							result = append(result, r.Interface())
						}
						return result, nil // массив возвращаемых значений
					}
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
				catcherr = NewError(stmt, err)
				break
			}
			regs.Set(s.RegRets, ret)

		case *BinFUNC:
			f := func(expr *BinFUNC, env *envir.Env) Func {
				return func(args ...interface{}) (interface{}, error) {
					if !expr.VarArg {
						if len(args) != len(expr.Args) {
							return nil, NewStringError(expr, "Неверное количество аргументов")
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
					if err == ReturnError {
						err = nil
					}
					// TODO: проверить при единичном и множественном возврате, при "..." аргументах
					return rr, err
				}
			}(s, env)
			env.Define(s.Name, f)
			regs.Set(s.Reg, f)

		case *BinCASTTYPE:
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
			rv := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()

			v, err := ast.TypeCastConvert(rv, nt, false, envir.NilValue)
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}

			regs.Set(s.Reg, v.Interface())

		case *BinMAKE:
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

		case *BinMAKECHAN:
			size, ok := regs.Reg[s.Reg].(int)
			if !ok {
				catcherr = NewStringError(stmt, "Размер должен быть целым числом")
				break
			}
			v := make(chan interface{}, size)
			regs.Set(s.Reg, v)

		case *BinMAKEARR:
			alen := int(ToInt64(regs.Reg[s.Reg]))
			acap := int(ToInt64(regs.Reg[s.RegCap]))
			v := make(VMSlice, alen, acap)
			regs.Set(s.Reg, v)

		case *BinCHANRECV:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			v, _ := ch.Recv()
			regs.Set(s.RegVal, v.Interface())

		case *BinCHANSEND:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			v := regs.Reg[s.RegVal]
			ch.Send(reflect.ValueOf(v).Elem())

		case *BinISKIND:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			regs.Set(s.Reg, v.Kind() == s.Kind)

		case *BinINC:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			var x interface{}
			if v.Kind() == reflect.Float64 {
				x = ToFloat64(v.Interface()) + 1.0
			} else {
				x = ToInt64(v.Interface()) + 1
			}
			regs.Set(s.Reg, x)

		case *BinDEC:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			var x interface{}
			if v.Kind() == reflect.Float64 {
				x = ToFloat64(v.Interface()) - 1.0
			} else {
				x = ToInt64(v.Interface()) - 1
			}
			regs.Set(s.Reg, x)

		case *BinTRY:
			regs.PushTry(s.Reg, s.JumpTo)
			regs.Set(s.Reg, nil) // изначально ошибки нет

		case *BinCATCH:
			// получаем ошибку, и если ее нет, переходим на метку, иначе, выполняем дальше
			nerr := regs.Reg[s.Reg]
			if nerr == nil {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *BinPOPTRY:
			// если catch блок отработал, то стек уже очищен, иначе снимаем со стека (ошибок не было)
			if regs.TopTryLabel() == s.CatchLabel {
				regs.PopTry()
			}

		case *BinFOREACH:
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

		case *BinNEXT:
			val := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()

			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				iter := regs.Reg[s.RegIter].(int)
				iter++
				if iter < val.Len() {
					regs.Set(s.RegIter, iter)
					iv := val.Index(iter)
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

		case *BinPOPFOR:
			if regs.TopContinue() == s.ContinueLabel {
				regs.PopContinue()
				regs.PopBreak()
			}

		case *BinFORNUM:

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

		case *BinNEXTNUM:
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

		case *BinWHILE:
			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *BinRET:
			retval = regs.Reg[s.Reg]
			return retval, ReturnError

		case *BinTHROW:
			catcherr = NewStringError(stmt, fmt.Sprint(regs.Reg[s.Reg]))
			break

		case *BinMODULE:
			// модуль регистрируется в глобальном контексте
			newenv := env.NewModule(ast.UniqueNames.Get(s.Name))
			_, err := Run(s.Code, newenv) // инициируем модуль
			if err != nil {
				catcherr = NewError(stmt, err)
				break
			}

		case *BinERROR:
			// необрабатываемая в попытке ошибка
			return retval, NewStringError(s, s.Error)

		case *BinBREAK:
			label := regs.PopBreak()
			if label != -1 {
				regs.PopContinue()
				idx = regs.Labels[label]
				continue
			}
			return nil, BreakError

		case *BinCONTINUE:
			label := regs.PopContinue()
			if label != -1 {
				regs.PopBreak()
				idx = regs.Labels[label]
				continue
			}
			return nil, ContinueError

		case *BinTRYRECV:

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

		case *BinTRYSEND:
			ch := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			if ch.Kind() != reflect.Chan {
				catcherr = NewStringError(stmt, "Не является каналом")
				break
			}
			ok := ch.TrySend(reflect.ValueOf(regs.Reg).Index(s.RegVal).Elem())
			regs.Set(s.RegOk, ok)

		case *BinGOSHED:
			runtime.Gosched()

		default:
			return nil, NewStringError(stmt, "Неизвестная инструкция")
		}

		if catcherr != nil {
			nerr := NewError(stmt, catcherr)
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
