// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"reflect"
	"runtime"

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
			if r, ok := regs.Reg[s.Reg].(bool); ok && !r {
				idx = regs.Labels[s.JumpTo]
				continue
			}

		case *BinJTRUE:
			if r, ok := regs.Reg[s.Reg].(bool); ok && r {
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
			// сначала простое сравнение
			if regs.Reg[s.Reg1] == regs.Reg[s.Reg2] {
				regs.Set(s.Reg, true)
			} else {
				// более глубокое сравнение через рефлексию
				regs.Set(s.Reg, reflect.DeepEqual(regs.Reg[s.Reg1], regs.Reg[s.Reg2]))
			}

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
			regs.Reg[s.Reg] = v.Interface()

		case *BinSET:
			env.Define(s.Id, regs.Reg[s.Reg])

		case *BinSETMEMBER:

		case *BinSETNAME:

		case *BinSETITEM:

		case *BinSETSLICE:

		case *BinUNARY:

		case *BinADDR:

		case *BinUNREF:

		case *BinOPER:

		case *BinCALL:

		case *BinGETMEMBER:

		case *BinGETIDX:

		case *BinGETSUBSLICE:

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
