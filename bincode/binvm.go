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
	"sync"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/core"
	"github.com/covrom/gonec/names"
	"github.com/covrom/gonec/parser"
	"github.com/shopspring/decimal"
)

func Interrupt(env *core.Env) {
	env.Interrupt()
}

// ParseSrc provides way to parse the code from source.
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

var binRegsPool = sync.Pool{}

func getRegs(ln int) core.VMSlice {
	sl := binRegsPool.Get()
	if sl != nil {
		vsl := sl.(core.VMSlice)
		if len(vsl) >= ln && len(vsl) < ln*2 {
			return vsl
		}
	}
	return make(core.VMSlice, ln)
}

func putRegs(sl core.VMSlice) {
	binRegsPool.Put(sl)
}

// Run запускает код на исполнение, например, после загрузки из файла
func Run(stmts binstmt.BinCode, env *core.Env) (retval core.VMValuer, reterr error) {
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

	// стандартная библиотека - загружаем, если она еще не была загружена в это или в родительское окружение

	if !env.IsBuiltsLoaded() {
		// эту функцию определяем тут, чтобы исключить циклические зависимости пакетов
		env.DefineS("загрузитьивыполнить", core.VMFunc(func(args core.VMSlice, rets *core.VMSlice) error {
			if len(args) != 1 {
				return errors.New("Должен быть один параметр")
			}
			if s, ok := args[0].(core.VMString); ok {
				body, err := ioutil.ReadFile(string(s))
				if err != nil {
					panic(err)
				}
				isGNX := strings.HasSuffix(strings.ToLower(string(s)), ".gnx")
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
					rets.Append(rv)
					return nil
				} else {
					_, bins, err := ParseSrc(string(body))
					if err != nil {
						if pe, ok := err.(*parser.Error); ok {
							pe.Filename = string(s)
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
					rets.Append(rv)
					return nil
				}
				return nil
			}
			return errors.New("Должен быть параметр-строка")
		}))

		core.LoadAllBuiltins(env)
	}

	retval, reterr = RunWorker(stmts.Code, stmts.Labels, stmts.MaxReg+1, env, 0)

	return
}

// RunWorker исполняет кусок кода, начиная с инструкции idx
func RunWorker(stmts binstmt.BinStmts, labels []int, numofregs int, env *core.Env, idx int) (retval core.VMValuer, reterr error) {
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

	// подготавливаем состояние машины: регистры значений, управляющие регистры

	registers := make(core.VMSlice, numofregs) //getRegs(numofregs)

	regs := &VMRegs{
		Env: env,
		// Reg:          registers,
		Labels:       labels,
		TryLabel:     make([]int, 0, 8),
		TryRegErr:    make([]int, 0, 8),
		ForBreaks:    make([]int, 0, 8),
		ForContinues: make([]int, 0, 8),
	}

	var (
		catcherr error
	)

	cntInterrupt := 0

	for idx < len(stmts) {

		// проверка прерывания каждые 10 команд
		cntInterrupt++
		if cntInterrupt == 10 {
			cntInterrupt = 0
			if regs.Env.CheckInterrupt() {
				// проверяем, был ли прерван интерпретатор
				return nil, binstmt.InterruptError
			}
		}

		stmt := stmts[idx]
		switch s := stmt.(type) {

		case *binstmt.BinJMP:
			idx = regs.Labels[s.JumpTo]
			continue

		case *binstmt.BinJFALSE:
			if b, ok := registers[s.Reg].(core.VMBooler); ok {
				if !b.Bool() {
					idx = regs.Labels[s.JumpTo]
					continue
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно определить значение булево")
				break
			}

		case *binstmt.BinJTRUE:
			if b, ok := registers[s.Reg].(core.VMBooler); ok {
				if b.Bool() {
					idx = regs.Labels[s.JumpTo]
					continue
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно определить значение булево")
				break
			}

		case *binstmt.BinLABEL:
			// пропускаем

		case *binstmt.BinLOAD:
			registers[s.Reg] = s.Val

		case *binstmt.BinOPER:
			v1 := registers[s.RegL]
			v2 := registers[s.RegR]
			if vv1, ok := v1.(core.VMOperationer); ok {
				if vv2, ok := v2.(core.VMOperationer); ok {
					if rv, err := vv1.EvalBinOp(s.Op, vv2); err == nil {
						registers[s.RegL] = rv
					} else {
						catcherr = binstmt.NewError(stmt, err)
						goto catching
					}
				} else {
					catcherr = binstmt.NewStringError(stmt, "Значение нельзя использовать в выражении")
					goto catching
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Значение нельзя использовать в выражении")
				goto catching
			}

		case *binstmt.BinMV:
			registers[s.RegTo] = registers[s.RegFrom]

		case *binstmt.BinEQUAL:
			v1 := registers[s.Reg1]
			v2 := registers[s.Reg2]
			if vv1, ok := v1.(core.VMOperationer); ok {
				if vv2, ok := v2.(core.VMOperationer); ok {
					if rv, err := vv1.EvalBinOp(core.EQL, vv2); err == nil {
						registers[s.Reg] = rv
					} else {
						catcherr = binstmt.NewError(stmt, err)
						break
					}
				} else {
					catcherr = binstmt.NewStringError(stmt, "Значение нельзя сравнивать")
					break
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Значение нельзя сравнивать")
				break
			}

		case *binstmt.BinCASTNUM:
			// ошибки обрабатываем в попытке
			var num core.VMNumberer
			var ok bool
			if num, ok = registers[s.Reg].(core.VMNumberer); !ok {
				registers[s.Reg] = nil
				catcherr = binstmt.NewStringError(stmt, "Литерал должен быть числом")
				break
			}
			v, err := num.InvokeNumber()
			if err != nil {
				registers[s.Reg] = nil
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			registers[s.Reg] = v

		case *binstmt.BinMAKESLICE:
			registers[s.Reg] = make(core.VMSlice, s.Len, s.Cap)

		case *binstmt.BinSETIDX:
			if v, ok := registers[s.Reg].(core.VMSlice); ok {
				v[s.Index] = registers[s.RegVal]
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно изменить значение по индексу")
				break
			}
		case *binstmt.BinMAKEMAP:
			registers[s.Reg] = make(core.VMStringMap, s.Len)

		case *binstmt.BinSETKEY:
			if v, ok := registers[s.Reg].(core.VMStringMap); ok {
				v[s.Key] = registers[s.RegVal]
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
			registers[s.Reg] = v

		case *binstmt.BinSET:
			// всегда сохраняются локальные переменные, глобальные и из внешнего окружения можно только читать
			env.Define(s.Id, registers[s.Reg])

		case *binstmt.BinSETMEMBER:
			m := registers[s.Reg]
			mv := registers[s.RegVal]
			switch mm := m.(type) {
			case core.VMMetaObject:
				mm.VMSetField(s.Id, mv.(core.VMInterfacer))
			case core.VMStringMap:
				mm[names.UniqueNames.Get(s.Id)] = mv
			default:
				catcherr = binstmt.NewStringError(stmt, "Невозможно установить поле у значения")
				goto catching
			}

		case *binstmt.BinSETNAME:
			v, ok := registers[s.Reg].(core.VMString)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Имя типа должно быть строкой")
				break
			}
			eType := names.UniqueNames.Set(string(v))
			registers[s.Reg] = core.VMInt(eType)

		case *binstmt.BinSETITEM:
			v := registers[s.Reg]
			i := registers[s.RegIndex]
			rv := registers[s.RegVal]
			registers[s.RegNeedLet] = core.VMBool(false)

			switch vv := v.(type) {
			case core.VMSlice:
				var ii int
				if iiv, ok := i.(core.VMInt); ok {
					ii = int(iiv)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть числом")
					goto catching
				}
				if ii < 0 {
					ii += len(vv)
				}
				if ii < 0 || ii >= len(vv) {
					catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
					goto catching
				}
				vv[ii] = rv
			case core.VMStringMap:
				if s, ok := i.(core.VMString); ok {
					vv[string(s)] = rv
				}
			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				goto catching
			}

		case *binstmt.BinSETSLICE:
			if vv, ok := registers[s.Reg].(core.VMSlice); ok {
				if rv, ok := registers[s.RegVal].(core.VMSlice); ok {

					vlen := len(vv)

					var rb int
					if registers[s.RegBegin] == nil {
						rb = 0
					} else if rbv, ok := registers[s.RegBegin].(core.VMInt); ok {
						rb = int(rbv)
					} else {
						catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
						goto catching
					}

					var re int
					if registers[s.RegEnd] == nil {
						re = vlen
					} else if rev, ok := registers[s.RegEnd].(core.VMInt); ok {
						re = int(rev)
					} else {
						catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
						goto catching
					}

					registers[s.RegNeedLet] = core.VMBool(false)

					ii, ij := LeftRightBounds(rb, re, vlen)
					if ij < ii {
						catcherr = binstmt.NewStringError(stmt, "Окончание диапазона не может быть раньше его начала")
						goto catching
					}

					if len(vv[ii:ij]) != len(rv) {
						catcherr = binstmt.NewStringError(stmt, "Размер массива должен быть равен ширине диапазона")
						goto catching
					}
					copy(vv[ii:ij], rv)

				} else {
					catcherr = binstmt.NewStringError(stmt, "Правая часть выражения должна быть массивом")
					goto catching
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Операция возможна только над массивом")
				goto catching
			}

		case *binstmt.BinUNARY:
			if vv, ok := registers[s.Reg].(core.VMUnarer); ok {
				rv, err := vv.EvalUnOp(s.Op)
				if err == nil {
					registers[s.Reg] = rv
				} else {
					catcherr = err
					break
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможна унарная операция над данным значением")
				break
			}

		// варианты ниже - не используются
		// case *binstmt.BinADDRID:
		// 	v, err := env.Get(s.Name)
		// 	if err != nil {
		// 		catcherr = binstmt.NewStringError(stmt, "Невозможно получить значение")
		// 		break
		// 	}
		// 	if !v.CanAddr() {
		// 		catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
		// 		break
		// 	}
		// 	regs.Set(s.Reg, v.Addr().Interface())

		// case *binstmt.BinADDRMBR:
		// 	refregs := reflect.ValueOf(regs.Reg)
		// 	v := refregs.Index(s.Reg).Elem()
		// 	if vme, ok := v.Interface().(*core.Env); ok {
		// 		m, err := vme.Get(s.Name)
		// 		if !m.IsValid() || err != nil {
		// 			catcherr = binstmt.NewStringError(stmt, "Значение не найдено")
		// 			break
		// 		}
		// 		if !m.CanAddr() {
		// 			catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
		// 			break
		// 		}
		// 		regs.Set(s.Reg, m.Addr().Interface())
		// 		break
		// 	}
		// 	m, err := GetMember(v, s.Name, s)
		// 	if err != nil {
		// 		catcherr = binstmt.NewError(stmt, err)
		// 		break
		// 	}
		// 	if !m.CanAddr() {
		// 		catcherr = binstmt.NewStringError(stmt, "Невозможно получить адрес значения")
		// 		break
		// 	}
		// 	regs.Set(s.Reg, m.Addr().Interface())

		// case *binstmt.BinUNREFID:
		// 	v, err := env.Get(s.Name)
		// 	if err != nil {
		// 		catcherr = binstmt.NewStringError(stmt, "Невозможно получить значение")
		// 		break
		// 	}
		// 	if v.Kind() != reflect.Ptr {
		// 		catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
		// 		break
		// 	}
		// 	regs.Set(s.Reg, v.Elem().Interface())

		// case *binstmt.BinUNREFMBR:
		// 	refregs := reflect.ValueOf(regs.Reg)
		// 	v := refregs.Index(s.Reg).Elem()
		// 	if vme, ok := v.Interface().(*core.Env); ok {
		// 		m, err := vme.Get(s.Name)
		// 		if !m.IsValid() || err != nil {
		// 			catcherr = binstmt.NewStringError(stmt, "Значение не найдено")
		// 			break
		// 		}
		// 		if m.Kind() != reflect.Ptr {
		// 			catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
		// 			break
		// 		}
		// 		regs.Set(s.Reg, m.Elem().Interface())
		// 		break
		// 	}
		// 	m, err := GetMember(v, s.Name, s)
		// 	if err != nil {
		// 		catcherr = binstmt.NewError(stmt, err)
		// 		break
		// 	}
		// 	if m.Kind() != reflect.Ptr {
		// 		catcherr = binstmt.NewStringError(stmt, "Отсутствует ссылка на значение")
		// 		break
		// 	}
		// 	regs.Set(s.Reg, m.Elem().Interface())

		case *binstmt.BinGETMEMBER:
			v := registers[s.Reg]
			switch vv := v.(type) {
			case *core.Env:
				// это идентификатор из модуля или окружения
				m, err := vv.Get(s.Name)
				if m == nil || err != nil {
					catcherr = binstmt.NewStringError(stmt, "Имя не найдено")
					goto catching
				}
				registers[s.Reg] = m
				goto catching
			case core.VMStringMap:
				// Сначала ищем поле, в нем может быть переопределен метод как функция
				if rv, ok := vv[names.UniqueNames.Get(s.Name)]; ok {
					registers[s.Reg] = rv
				} else {
					if ff, ok := vv.MethodMember(s.Name); ok {
						registers[s.Reg] = ff
					} else {
						registers[s.Reg] = core.VMNil
					}
				}
			case core.VMMetaObject:
				if vv.VMIsField(s.Name) {
					registers[s.Reg] = vv.VMGetField(s.Name)
				} else {
					if ff, ok := vv.VMGetMethod(s.Name); ok {
						registers[s.Reg] = ff
					} else {
						catcherr = binstmt.NewStringError(stmt, "Нет поля или метода с таким именем")
						goto catching
					}
				}
			case core.VMMethodImplementer:
				if ff, ok := vv.MethodMember(s.Name); ok {
					registers[s.Reg] = ff
				} else {
					catcherr = binstmt.NewStringError(stmt, "Нет метода с таким именем")
					goto catching
				}
			default:
				catcherr = binstmt.NewStringError(stmt, "У значения не бывает полей или методов")
				goto catching
			}

		case *binstmt.BinGETIDX:
			v := registers[s.Reg]
			i := registers[s.RegIndex]
			switch vv := v.(type) {
			case core.VMSlice:
				if iv, ok := i.(core.VMInt); ok {
					ii := int(iv)
					if ii < 0 {
						ii += len(vv)
					}
					if ii < 0 || ii >= len(vv) {
						catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
						goto catching
					}
					registers[s.Reg] = vv[ii]
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					goto catching
				}
			case core.VMString:
				if iv, ok := i.(core.VMInt); ok {
					ii := int(iv)
					r := []rune(string(vv))
					if ii < 0 {
						ii += len(r)
					}
					if ii < 0 || ii >= len(r) {
						catcherr = binstmt.NewStringError(stmt, "Индекс за пределами границ")
						goto catching
					}
					registers[s.Reg] = core.VMString(string(r[ii]))
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					goto catching
				}
			case core.VMStringMap:
				if k, ok := i.(core.VMString); ok {
					registers[s.Reg] = vv[string(k)]
				} else {
					catcherr = binstmt.NewStringError(stmt, "Ключ должен быть строкой")
					goto catching
				}
			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				goto catching
			}

		case *binstmt.BinGETSUBSLICE:

			var rb int
			if registers[s.RegBegin] == nil {
				rb = 0
			} else if rbv, ok := registers[s.RegBegin].(core.VMInt); ok {
				rb = int(rbv)
			} else {
				catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
				goto catching
			}

			switch vv := registers[s.Reg].(type) {
			case core.VMSlice:
				vlen := len(vv)

				var re int
				if registers[s.RegEnd] == nil {
					re = vlen
				} else if rev, ok := registers[s.RegEnd].(core.VMInt); ok {
					re = int(rev)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					goto catching
				}

				ii, ij := LeftRightBounds(rb, re, vlen)

				if ij < ii {
					catcherr = binstmt.NewStringError(stmt, "Окончание диапазона не может быть раньше его начала")
					goto catching
				}

				registers[s.Reg] = vv[ii:ij]

			case core.VMString:
				r := []rune(string(vv))

				vlen := len(r)

				var re int
				if registers[s.RegEnd] == nil {
					re = vlen
				} else if rev, ok := registers[s.RegEnd].(core.VMInt); ok {
					re = int(rev)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					goto catching
				}

				ii, ij := LeftRightBounds(rb, re, vlen)

				if ij < ii {
					catcherr = binstmt.NewStringError(stmt, "Окончание диапазона не может быть раньше его начала")
					goto catching
				}

				registers[s.Reg] = core.VMString(string(r[ii:ij]))

			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinCALL:

			var err error

			//функцию на языке Гонец можно вызывать прямо с аргументами из слайса в регистре
			var fgnc core.VMValuer
			var argsl core.VMSlice
			if s.Name == 0 {
				fgnc = registers[s.RegArgs]
				argsl = registers[s.RegArgs+1 : s.RegArgs+1+s.NumArgs]
			} else {
				fgnc, err = env.Get(s.Name)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					goto catching
				}
				argsl = registers[s.RegArgs : s.RegArgs+s.NumArgs]
			}
			rets := core.GetGlobalVMSlice()
			if fnc, ok := fgnc.(core.VMFunc); ok {
				// если ее надо вызвать в горутине - вызываем
				if s.Go {
					env.SetGoRunned(true)
					rets = core.GetGlobalVMSlice() // для каждой горутины отдельный массив возвратов, который потом не используется
					go func(args, rets core.VMSlice) {
						fnc(argsl, &rets)
						core.PutGlobalVMSlice(rets) // всегда возвращаем в пул
					}(argsl, rets)
					registers[s.RegRets] = core.VMSlice{} // для такого вызова - всегда пустой массив возвратов
					break
				}

				// не в горутине
				err = fnc(argsl, &rets)

				// TODO: проверить, если был передан слайс, и он изменен внутри функции, то что происходит в исходном слайсе?

				if err != nil {
					// ошибку передаем в блок обработки исключений
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				switch len(rets) {
				case 0:
					registers[s.RegRets] = core.VMNil
					core.PutGlobalVMSlice(rets)
				case 1:
					registers[s.RegRets] = rets[0]
					core.PutGlobalVMSlice(rets)
				default:
					registers[s.RegRets] = rets //не возвращаем в пул
				}
				break
			} else {

				// fmt.Printf("%T\n", fgnc)

				catcherr = binstmt.NewStringError(stmt, "Неверный тип функции")
				goto catching
			}

		case *binstmt.BinFUNC:

			f := func(expr *binstmt.BinFUNC, fstmts binstmt.BinStmts, flabels []int, fenv *core.Env) core.VMFunc {
				return func(args core.VMSlice, rets *core.VMSlice) error {
					if !expr.VarArg {
						if len(args) != len(expr.Args) {
							return binstmt.NewStringError(expr, "Неверное количество аргументов")
						}
					}
					var newenv *core.Env
					if expr.Name == 0 {
						// наследуем от окружения текущей функции
						newenv = fenv.NewSubEnv()
					} else {
						// наследуем от модуля или глобального окружения
						newenv = fenv.NewEnv()
					}

					// переменное число аргументов передается как один параметр-слайс
					if expr.VarArg {
						newenv.Define(expr.Args[0], args)
					} else {
						for i, arg := range expr.Args {
							newenv.Define(arg, args[i])
						}
					}
					// вызов функции возвращает одиночное значение (в т.ч. VMNil) или VMSlice

					rr, err := RunWorker(fstmts, flabels, expr.MaxReg+1, newenv, flabels[expr.LabelStart])

					if err == binstmt.ReturnError {
						err = nil
					}
					// возврат массива возвращается сразу, иначе добавляется
					if vsl, ok := rr.(core.VMSlice); ok {
						*rets = vsl
					} else {
						rets.Append(rr)
					}
					newenv.Destroy()
					return err
				}
			}(s, stmts, labels, env)

			env.Define(s.Name, f)
			registers[s.Reg] = f
			idx = regs.Labels[s.LabelEnd]

		case *binstmt.BinRET:
			retval = registers[s.Reg]
			return retval, binstmt.ReturnError

		case *binstmt.BinCASTTYPE:
			// приведение типов, включая приведение типов в массиве как новый типизированный массив
			eType, ok := registers[s.TypeReg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Неизвестный тип")
				break
			}
			nt, err := env.Type(int(eType))
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			rv := registers[s.Reg]
			if cv, ok := rv.(core.VMConverter); ok {
				v, err := cv.ConvertToType(nt)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				registers[s.Reg] = v
			} else {
				catcherr = binstmt.NewStringError(stmt, "Значение не может быть преобразовано")
				break
			}

		case *binstmt.BinMAKE:
			eType, ok := registers[s.Reg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Неизвестный тип")
				break
			}
			rt, err := env.Type(int(eType))
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			var v reflect.Value
			if rt.Kind() == reflect.Map {
				v = reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt)
			} else if rt.Kind() == reflect.Struct {
				// структуру создаем всегда ссылочной
				v = reflect.New(rt)
			} else {
				v = reflect.Zero(rt)
			}
			if vv, ok := v.Interface().(core.VMValuer); ok {
				if vobj, ok := vv.(core.VMMetaObject); ok {
					vobj.VMInit(vobj)
					vobj.VMRegister()
					registers[s.Reg] = vobj
				} else {
					registers[s.Reg] = vv
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Неизвестный тип")
				break
			}

		case *binstmt.BinMAKECHAN:
			size, ok := registers[s.Reg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Размер должен быть целым числом")
				break
			}
			v := make(core.VMChan, int(size))
			registers[s.Reg] = v

		case *binstmt.BinMAKEARR:
			alen, ok := registers[s.Reg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Длина должна быть целым числом")
				break
			}
			acap, ok := registers[s.RegCap].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Размер должен быть целым числом")
				break
			}

			v := make(core.VMSlice, int(alen), int(acap))
			registers[s.Reg] = v

		case *binstmt.BinCHANRECV:
			ch, ok := registers[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			v, ok := ch.Recv()
			if !ok {
				// если закрыт, то пишем nil
				registers[s.RegVal] = core.VMNil
			} else {
				registers[s.RegVal] = v
			}

		case *binstmt.BinCHANSEND:
			ch, ok := registers[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			v := registers[s.RegVal]
			ch.Send(v)

		case *binstmt.BinISKIND:
			v := reflect.ValueOf(registers).Index(s.Reg).Elem()
			registers[s.Reg] = core.VMBool(v.Kind() == s.Kind)

		case *binstmt.BinISSLICE:
			_, ok := registers[s.Reg].(core.VMSlice)
			registers[s.RegBool] = core.VMBool(ok)

		case *binstmt.BinINC:
			v := registers[s.Reg]
			var x core.VMValuer
			if vv, ok := v.(core.VMInt); ok {
				x = core.VMInt(int64(vv) + 1)
			} else if vv, ok := v.(core.VMDecimal); ok {
				x = vv.Add(core.VMDecimal(decimal.New(1, 0)))
			}
			registers[s.Reg] = x

		case *binstmt.BinDEC:
			v := registers[s.Reg]
			var x core.VMValuer
			if vv, ok := v.(core.VMInt); ok {
				x = core.VMInt(int64(vv) - 1)
			} else if vv, ok := v.(core.VMDecimal); ok {
				x = vv.Add(core.VMDecimal(decimal.New(-1, 0)))
			}
			registers[s.Reg] = x

		case *binstmt.BinTRY:
			regs.PushTry(s.Reg, s.JumpTo)
			registers[s.Reg] = nil // изначально ошибки нет

		case *binstmt.BinCATCH:
			// получаем ошибку, и если ее нет, переходим на метку, иначе, выполняем дальше
			nerr := registers[s.Reg]
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
			val := registers[s.Reg]

			switch val.(type) {
			case core.VMSlice:
				registers[s.RegIter] = core.VMInt(-1)
			case core.VMChan:
				registers[s.RegIter] = nil
			default:
				catcherr = binstmt.NewStringError(stmt, "Не является коллекцией или каналом")
				goto catching
			}

			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *binstmt.BinNEXT:
			val := registers[s.Reg]

			switch vv := val.(type) {
			case core.VMSlice:
				iter := int(registers[s.RegIter].(core.VMInt))
				iter++
				if iter < len(vv) {
					registers[s.RegIter] = core.VMInt(iter)
					registers[s.RegVal] = vv[iter]
				} else {
					idx = regs.Labels[s.JumpTo]
					continue
				}
			case core.VMChan:
				iv, ok := vv.Recv()
				if !ok {
					registers[s.RegVal] = core.VMNil
				} else {
					registers[s.RegVal] = iv
				}

			default:
				catcherr = binstmt.NewStringError(stmt, "Не является коллекцией или каналом")
				goto catching
			}

		case *binstmt.BinPOPFOR:
			if regs.TopContinue() == s.ContinueLabel {
				regs.PopContinue()
				regs.PopBreak()
			}

		case *binstmt.BinFORNUM:
			if _, ok := registers[s.RegFrom].(core.VMInt); ok {
				if _, ok := registers[s.RegTo].(core.VMInt); ok {
					registers[s.Reg] = nil
					regs.PushBreak(s.BreakLabel)
					regs.PushContinue(s.ContinueLabel)
				} else {
					catcherr = binstmt.NewStringError(stmt, "Конечное значение должно быть целым числом")
					break
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Начальное значение должно быть целым числом")
				break
			}

		case *binstmt.BinNEXTNUM:
			afrom := int64(registers[s.RegFrom].(core.VMInt))
			ato := int64(registers[s.RegTo].(core.VMInt))
			fviadd := int64(1)
			if afrom > ato {
				fviadd = int64(-1) // если конечное значение меньше первого, идем в обратном порядке
			}
			vv := registers[s.Reg]
			var iter int64
			if vv == nil {
				iter = afrom
			} else {
				iter = int64(vv.(core.VMInt))
				iter += fviadd
			}
			inrange := iter <= ato
			if afrom > ato {
				inrange = iter >= ato
			}
			if inrange {
				registers[s.Reg] = core.VMInt(iter)
			} else {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *binstmt.BinWHILE:
			regs.PushBreak(s.BreakLabel)
			regs.PushContinue(s.ContinueLabel)

		case *binstmt.BinTHROW:
			catcherr = binstmt.NewStringError(stmt, fmt.Sprint(registers[s.Reg]))
			break

		case *binstmt.BinMODULE:
			// модуль регистрируется в глобальном контексте
			newenv := env.NewModule(names.UniqueNames.Get(s.Name))
			_, err := Run(s.Code, newenv) // инициируем модуль
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}

		case *binstmt.BinERROR:
			// необрабатываемая в попытке ошибка
			return retval, binstmt.NewStringError(s, s.Error)

		case *binstmt.BinBREAK:
			label := regs.PopBreak()
			if label != -1 {
				regs.PopContinue()
				idx = regs.Labels[label]
				continue
			}
			return nil, binstmt.BreakError

		case *binstmt.BinCONTINUE:
			label := regs.PopContinue()
			if label != -1 {
				regs.PopBreak()
				idx = regs.Labels[label]
				continue
			}
			return nil, binstmt.ContinueError

		case *binstmt.BinTRYRECV:

			ch, ok := registers[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			v, ok, notready := ch.TryRecv()
			if !ok {
				registers[s.RegVal] = core.VMNil
				registers[s.RegOk] = core.VMBool(ok)
				registers[s.RegClosed] = core.VMBool(!notready)
			} else {
				registers[s.RegVal] = v
				registers[s.RegOk] = core.VMBool(ok)
				registers[s.RegClosed] = core.VMBool(false)
			}

		case *binstmt.BinTRYSEND:
			ch, ok := registers[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			ok = ch.TrySend(registers[s.RegVal])
			registers[s.RegOk] = core.VMBool(ok)

		case *binstmt.BinGOSHED:
			runtime.Gosched()

		// case *binstmt.BinFREE:
		// 	regs.FreeFromReg(s.Reg)

		default:
			return nil, binstmt.NewStringError(stmt, "Неизвестная инструкция")
		}

	catching:
		if catcherr != nil {
			nerr := binstmt.NewError(stmt, catcherr)
			catcherr = nil
			// учитываем стек обработки ошибок
			if regs.TopTryLabel() == -1 {
				return nil, nerr
			} else {
				env.DefineS("описаниеошибки", func(s string) core.VMFunc {
					return func(args core.VMSlice, rets *core.VMSlice) error {
						if len(args) != 0 {
							return errors.New("Данная функция не требует параметров")
						}
						rets.Append(core.VMString(s))
						return nil
					}
				}(nerr.Error()))

				r, idxl := regs.PopTry()
				registers[r] = core.VMString(nerr.Error())
				idx = regs.Labels[idxl] // переходим в catch блок, функция с описанием ошибки определена
				continue
			}
		}

		idx++
	}

	// putRegs(registers)

	return retval, nil
}
