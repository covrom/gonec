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

	registers := make([]core.VMValuer, stmts.MaxReg+1)
	return RunWorker(stmts.Code, stmts.Labels, registers, env, 0)
}

// RunWorker исполняет кусок кода, начиная с инструкции idx
func RunWorker(stmts binstmt.BinStmts, labels []int, registers []core.VMValuer, env *core.Env, idx int) (retval core.VMValuer, reterr error) {
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

	regs := &VMRegs{
		Env:          env,
		Reg:          registers,
		Labels:       labels,
		TryLabel:     make([]int, 0, 8),
		TryRegErr:    make([]int, 0, 8),
		ForBreaks:    make([]int, 0, 8),
		ForContinues: make([]int, 0, 8),
	}

	retsSlice := make(core.VMSlice, 0, 20) // кэширующий слайс возвращаемых значений из функций VMFunc

	var (
		catcherr error
	)

	for idx < len(stmts) {

		if regs.Env.CheckInterrupt() {
			// проверяем, был ли прерван интерпретатор
			return nil, binstmt.InterruptError
		}

		stmt := stmts[idx]
		switch s := stmt.(type) {

		case *binstmt.BinJMP:
			idx = regs.Labels[s.JumpTo]
			continue

		case *binstmt.BinJFALSE:
			if b, ok := regs.Reg[s.Reg].(core.VMBooler); ok {
				if !b.Bool() {
					idx = regs.Labels[s.JumpTo]
					continue
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно определить значение булево")
				break
			}

		case *binstmt.BinJTRUE:
			if b, ok := regs.Reg[s.Reg].(core.VMBooler); ok {
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
			regs.Reg[s.Reg] = s.Val

		case *binstmt.BinMV:
			regs.Reg[s.RegTo] = regs.Reg[s.RegFrom]

		case *binstmt.BinEQUAL:
			v1 := regs.Reg[s.Reg1]
			v2 := regs.Reg[s.Reg2]
			if vv1, ok := v1.(core.VMOperationer); ok {
				if vv2, ok := v2.(core.VMOperationer); ok {
					if rv, err := vv1.EvalBinOp(core.EQL, vv2); err == nil {
						regs.Reg[s.Reg] = rv
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
			if num, ok = regs.Reg[s.Reg].(core.VMNumberer); !ok {
				regs.Reg[s.Reg] = nil
				catcherr = binstmt.NewStringError(stmt, "Литерал должен быть числом")
				break
			}
			v, err := num.InvokeNumber()
			if err != nil {
				regs.Reg[s.Reg] = nil
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			regs.Reg[s.Reg] = v

		case *binstmt.BinMAKESLICE:
			regs.Reg[s.Reg] = make(core.VMSlice, s.Len, s.Cap)

		case *binstmt.BinSETIDX:
			if v, ok := regs.Reg[s.Reg].(core.VMSlice); ok {
				v[s.Index] = regs.Reg[s.RegVal]
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможно изменить значение по индексу")
				break
			}
		case *binstmt.BinMAKEMAP:
			regs.Reg[s.Reg] = make(core.VMStringMap, s.Len)

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
			regs.Reg[s.Reg] = v

		case *binstmt.BinSET:
			env.Define(s.Id, regs.Reg[s.Reg])

		case *binstmt.BinSETMEMBER:
			m := regs.Reg[s.Reg]
			mv := regs.Reg[s.RegVal]
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
			v, ok := regs.Reg[s.Reg].(core.VMString)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Имя типа должно быть строкой")
				break
			}
			eType := names.UniqueNames.Set(string(v))
			regs.Reg[s.Reg] = core.VMInt(eType)

		case *binstmt.BinSETITEM:
			v := regs.Reg[s.Reg]
			i := regs.Reg[s.RegIndex]
			rv := regs.Reg[s.RegVal]
			regs.Reg[s.RegNeedLet] = core.VMBool(false)

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
			if vv, ok := regs.Reg[s.Reg].(core.VMSlice); ok {
				if rv, ok := regs.Reg[s.RegVal].(core.VMSlice); ok {

					vlen := len(vv)

					var rb int
					if regs.Reg[s.RegBegin] == nil {
						rb = 0
					} else if rbv, ok := regs.Reg[s.RegBegin].(core.VMInt); ok {
						rb = int(rbv)
					} else {
						catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
						goto catching
					}

					var re int
					if regs.Reg[s.RegEnd] == nil {
						re = vlen
					} else if rev, ok := regs.Reg[s.RegEnd].(core.VMInt); ok {
						re = int(rev)
					} else {
						catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
						goto catching
					}

					regs.Reg[s.RegNeedLet] = core.VMBool(false)

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
			if vv, ok := regs.Reg[s.Reg].(core.VMUnarer); ok {
				rv, err := vv.EvalUnOp(s.Op)
				if err == nil {
					regs.Reg[s.Reg] = rv
				} else {
					catcherr = err
					break
				}
			} else {
				catcherr = binstmt.NewStringError(stmt, "Невозможна унарная операция над данным значением")
				break
			}

			// switch s.Op {
			// case '-':
			// 	if x, ok := regs.Reg[s.Reg].(core.VMInt); ok {
			// 		regs.Reg[s.Reg]= core.VMInt(-int64(x))
			// 	} else if x, ok := regs.Reg[s.Reg].(int64); ok {
			// 		regs.Set(s.Reg, -x)
			// 	} else {
			// 		catcherr = binstmt.NewStringError(stmt, "Операция применима только к числам")
			// 		break
			// 	}
			// case '^':
			// 	if x, ok := regs.Reg[s.Reg].(int64); ok {
			// 		regs.Set(s.Reg, ^x)
			// 	} else {
			// 		catcherr = binstmt.NewStringError(stmt, "Операция применима только к целым числам")
			// 		break
			// 	}
			// case '!':
			// 	regs.Set(s.Reg, !ToBool(regs.Reg[s.Reg]))

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
			v := regs.Reg[s.Reg]
			switch vv := v.(type) {
			case *core.Env:
				// это идентификатор из модуля или окружения
				m, err := vv.Get(s.Name)
				if m == nil || err != nil {
					catcherr = binstmt.NewStringError(stmt, "Имя не найдено")
					goto catching
				}
				regs.Reg[s.Reg] = m
				goto catching
			case core.VMStringMap:
				if rv, ok := vv[names.UniqueNames.Get(s.Name)]; ok {
					regs.Reg[s.Reg] = rv
				} else {
					regs.Reg[s.Reg] = core.VMNil
				}
			case core.VMMetaObject:
				if vv.VMIsField(s.Name) {
					regs.Reg[s.Reg] = vv.VMGetField(s.Name)
				} else {
					if ff, ok := vv.VMGetMethod(s.Name); ok {
						regs.Reg[s.Reg] = ff
					} else {
						catcherr = binstmt.NewStringError(stmt, "Нет поля или метода с таким именем")
						goto catching
					}
				}
			default:
				catcherr = binstmt.NewStringError(stmt, "У значения не бывает полей или методов")
				goto catching
			}

		case *binstmt.BinGETIDX:
			v := regs.Reg[s.Reg]
			i := regs.Reg[s.RegIndex]
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
					regs.Reg[s.Reg] = vv[ii]
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
					regs.Reg[s.Reg] = core.VMString(string(r[ii]))
				} else {
					catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
					goto catching
				}
			case core.VMStringMap:
				if k, ok := i.(core.VMString); ok {
					regs.Reg[s.Reg] = vv[string(k)]
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
			if regs.Reg[s.RegBegin] == nil {
				rb = 0
			} else if rbv, ok := regs.Reg[s.RegBegin].(core.VMInt); ok {
				rb = int(rbv)
			} else {
				catcherr = binstmt.NewStringError(stmt, "Индекс должен быть целым числом")
				goto catching
			}

			switch vv := regs.Reg[s.Reg].(type) {
			case core.VMSlice:
				vlen := len(vv)

				var re int
				if regs.Reg[s.RegEnd] == nil {
					re = vlen
				} else if rev, ok := regs.Reg[s.RegEnd].(core.VMInt); ok {
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

				regs.Reg[s.Reg] = vv[ii:ij]

			case core.VMString:
				r := []rune(string(vv))

				vlen := len(r)

				var re int
				if regs.Reg[s.RegEnd] == nil {
					re = vlen
				} else if rev, ok := regs.Reg[s.RegEnd].(core.VMInt); ok {
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

				regs.Reg[s.Reg] = core.VMString(string(r[ii:ij]))

			default:
				catcherr = binstmt.NewStringError(stmt, "Неверная операция")
				break
			}

		case *binstmt.BinOPER:
			v1 := regs.Reg[s.RegL]
			v2 := regs.Reg[s.RegR]
			if vv1, ok := v1.(core.VMOperationer); ok {
				if vv2, ok := v2.(core.VMOperationer); ok {
					if rv, err := vv1.EvalBinOp(s.Op, vv2); err == nil {
						regs.Reg[s.RegL] = rv
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

		case *binstmt.BinCALL:

			var err error

			//функцию на языке Гонец можно вызывать прямо с аргументами из слайса в регистре
			var fgnc core.VMValuer
			var argsl core.VMSlice
			if s.Name == 0 {
				fgnc = regs.Reg[s.RegArgs]
				argsl = regs.Reg[s.RegArgs+1].(core.VMSlice)
			} else {
				fgnc, err = env.Get(s.Name)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					goto catching
				}
				argsl = regs.Reg[s.RegArgs].(core.VMSlice)
			}
			rets := retsSlice[:0]
			if fnc, ok := fgnc.(core.VMFunc); ok {
				// если ее надо вызвать в горутине - вызываем
				if s.Go {
					env.SetGoRunned(true)
					rets = make(core.VMSlice, 0) // для каждой горутины отдельный массив возвратов, который потом не используется
					go fnc(argsl, &rets)
					regs.Reg[s.RegRets] = rets
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
				regs.Reg[s.RegRets] = rets
				break
			} else {
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

					if expr.VarArg {
						newenv.Define(expr.Args[0], args)
					} else {
						for i, arg := range expr.Args {
							newenv.Define(arg, args[i])
						}
					}
					// вызов функции возвращает одиночное значение (в т.ч. VMNil) или VMSlice
					callregs := make([]core.VMValuer, expr.MaxReg+1)
					rr, err := RunWorker(fstmts, flabels, callregs, newenv, expr.LabelStart)
					if err == binstmt.ReturnError {
						err = nil
					}
					// возврат массива возвращается сразу, иначе добавляется
					if vsl, ok := rr.(core.VMSlice); ok {
						*rets = vsl
					} else {
						rets.Append(rr)
					}
					return err
				}
			}(s, stmts, labels, env)

			env.Define(s.Name, f)
			regs.Reg[s.Reg] = f

		case *binstmt.BinRET:
			retval = regs.Reg[s.Reg]
			return retval, binstmt.ReturnError

		case *binstmt.BinCASTTYPE:
			// приведение типов, включая приведение типов в массиве как новый типизированный массив
			eType, ok := regs.Reg[s.TypeReg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Неизвестный тип")
				break
			}
			nt, err := env.Type(int(eType))
			if err != nil {
				catcherr = binstmt.NewError(stmt, err)
				break
			}
			rv := regs.Reg[s.Reg]
			if cv, ok := rv.(core.VMConverter); ok {
				v, err := cv.ConvertToType(nt, false)
				if err != nil {
					catcherr = binstmt.NewError(stmt, err)
					break
				}
				regs.Reg[s.Reg] = v
			} else {
				catcherr = binstmt.NewStringError(stmt, "Значение не может быть преобразовано")
				break
			}

		case *binstmt.BinMAKE:
			eType, ok := regs.Reg[s.Reg].(core.VMInt)
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
				regs.Reg[s.Reg] = vv
			} else {
				catcherr = binstmt.NewStringError(stmt, "Неизвестный тип")
				break
			}

		case *binstmt.BinMAKECHAN:
			size, ok := regs.Reg[s.Reg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Размер должен быть целым числом")
				break
			}
			v := make(core.VMChan, int(size))
			regs.Reg[s.Reg] = v

		case *binstmt.BinMAKEARR:
			alen, ok := regs.Reg[s.Reg].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Длина должна быть целым числом")
				break
			}
			acap, ok := regs.Reg[s.RegCap].(core.VMInt)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Размер должен быть целым числом")
				break
			}

			v := make(core.VMSlice, int(alen), int(acap))
			regs.Reg[s.Reg] = v

		case *binstmt.BinCHANRECV:
			ch, ok := regs.Reg[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			v := ch.Recv()
			regs.Reg[s.RegVal] = v

		case *binstmt.BinCHANSEND:
			ch, ok := regs.Reg[s.Reg].(core.VMChan)
			if !ok {
				catcherr = binstmt.NewStringError(stmt, "Не является каналом")
				break
			}
			v := regs.Reg[s.RegVal]
			ch.Send(v)

		case *binstmt.BinISKIND:
			v := reflect.ValueOf(regs.Reg).Index(s.Reg).Elem()
			regs.Reg[s.Reg] = core.VMBool(v.Kind() == s.Kind)

		case *binstmt.BinISSLICE:
			_, ok := regs.Reg[s.Reg].(core.VMSlice)
			regs.Reg[s.RegBool] = core.VMBool(ok)

		case *binstmt.BinINC:
			v := regs.Reg[s.Reg]
			var x core.VMValuer
			if vv, ok := v.(core.VMInt); ok {
				x = core.VMInt(int64(vv) + 1)
			} else if vv, ok := v.(core.VMDecimal); ok {
				x = vv.Add(core.VMDecimal(decimal.New(1, 0)))
			}
			regs.Reg[s.Reg] = x

		case *binstmt.BinDEC:
			v := regs.Reg[s.Reg]
			var x core.VMValuer
			if vv, ok := v.(core.VMInt); ok {
				x = core.VMInt(int64(vv) - 1)
			} else if vv, ok := v.(core.VMDecimal); ok {
				x = vv.Add(core.VMDecimal(decimal.New(-1, 0)))
			}
			regs.Reg[s.Reg] = x

		case *binstmt.BinTRY:
			regs.PushTry(s.Reg, s.JumpTo)
			regs.Reg[s.Reg] = nil // изначально ошибки нет

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

		case *binstmt.BinTHROW:
			catcherr = NewStringError(stmt, fmt.Sprint(regs.Reg[s.Reg]))
			break

		case *binstmt.BinMODULE:
			// модуль регистрируется в глобальном контексте
			newenv := env.NewModule(names.UniqueNames.Get(s.Name))
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

	catching:
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
