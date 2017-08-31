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
			v := refregs.Index(s.Reg)
			rv := refregs.Index(s.RegVal)
			// if v.Kind() == reflect.Interface {
			// 	v = v.Elem()
			// }
			// if v.Kind() == reflect.Slice {
			// 	v = v.Index(0)
			// }

			// if !v.IsValid() {
			// 	return NilValue, NewStringError(expr, "Поле недоступно")
			// }

			// if v.Kind() == reflect.Ptr {
			// 	v = v.Elem()
			// }
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
			v := refregs.Index(s.Reg)
			i := refregs.Index(s.RegIndex)
			rv := refregs.Index(s.RegVal)
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
			v := refregs.Index(s.Reg)
			rb := refregs.Index(s.RegBegin)
			re := refregs.Index(s.RegEnd)
			rv := refregs.Index(s.RegVal)
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
			v := refregs.Index(s.Reg)
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
			v := refregs.Index(s.Reg)
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
			v := refregs.Index(s.Reg)
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
			v := refregs.Index(s.Reg)
			i := refregs.Index(s.RegIndex)

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
			v := refregs.Index(s.Reg)
			rb := refregs.Index(s.RegBegin)
			re := refregs.Index(s.RegEnd)

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
			lhsV := refregs.Index(s.RegL)
			rhsV := refregs.Index(s.RegR)

			r, err := EvalBinOp(s.Op, lhsV, rhsV)
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
				f = reflect.ValueOf(regs.Reg).Index(s.RegArgs)
				// в следующем - массив аргументов
				vargs = reflect.ValueOf(regs.Reg).Index(s.RegArgs + 1)
			} else {
				// функция берется из переменной или по имени в окружении
				f, err = env.Get(s.Name)
				if err != nil {
					catcherr = NewError(stmt, err)
					break
				}
				// в регистре - массив аргументов
				vargs = reflect.ValueOf(regs.Reg).Index(s.RegArgs)
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
			args := make([]reflect.Value, s.NumArgs)

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
						} else if arg.Kind() == reflect.Func {
							if _, isFunc := arg.Interface().(Func); isFunc {
								// это функция на языке Гонец (т.е. обработчик) - делаем обертку в целевую функцию типа it
								rfunc := arg
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
				// if !isReflect {
				// 	// для функций на языке Го
				if s.VarArg && i == s.NumArgs-1 {
					for j := 0; j < arg.Len(); j++ {
						args = append(args, arg.Index(j))
					}
				} else {
					args = append(args, arg)
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
					// TODO: проверить при единичном и множественном возврате
					return rr, err
				}
			}(s, env)
			env.Define(s.Name, f)
			regs.Set(s.Reg, f)

		case *BinCASTTYPE:

		case *BinMAKE:

		case *BinMAKECHAN:

		case *BinMAKEARR:

		case *BinCHANRECV:

		case *BinCHANSEND:

		case *BinISKIND:

		case *BinINC:

		case *BinDEC:

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

		case *BinNEXT:

		case *BinPOPFOR:

		case *BinFORNUM:

		case *BinNEXTNUM:

		case *BinWHILE:

		case *BinBREAK:

		case *BinCONTINUE:

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
		case *BinTRYRECV:

		case *BinTRYSEND:

		case *BinGOSHED:
			runtime.Gosched()
		default:
			return nil, NewStringError(stmt, "Неизвестная инструкция")
		}
		if catcherr != nil {
			catcherr = nil
			nerr := NewError(stmt, catcherr)
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
