// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"fmt"
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

		case *BinFUNC:

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
			retval = regs.Reg[0] // из основного регистра
			return retval, ReturnError
		case *BinTHROW:
			catcherr = NewStringError(stmt, fmt.Sprint(regs.Reg[s.Reg]))
			break

		case *BinMODULE:

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
