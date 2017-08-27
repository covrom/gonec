// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"reflect"
	"runtime"

	"github.com/covrom/gonec/ast"
	envir "github.com/covrom/gonec/env"
)

func Interrupt(env *envir.Env) {
	env.Interrupt()
}

func Run(stmts BinCode, env *envir.Env) (rv interface{}, err error) {
	// подготавливаем состояние машины: регистры значений, управляющие регистры
	regs := NewVMRegs(stmts)

	goschedidx := 0

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
				regs.Set(s.Reg, ast.Equal(reflect.ValueOf(regs.Reg[s.Reg1]), reflect.ValueOf(regs.Reg[s.Reg2])))
			}
		case *BinCASTNUM:

		case *BinMAKESLICE:

		case *BinSETIDX:

		case *BinMAKEMAP:

		case *BinSETKEY:

		case *BinGET:

		case *BinSET:

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

		case *BinCATCH:

		case *BinPOPTRY:

		case *BinFOREACH:

		case *BinNEXT:

		case *BinPOPFOR:

		case *BinFORNUM:

		case *BinNEXTNUM:

		case *BinWHILE:

		case *BinBREAK:

		case *BinCONTINUE:

		case *BinRET:
			rv = regs.Reg[0] // из основного регистра
			return rv, ReturnError
		case *BinTHROW:

		case *BinMODULE:

		case *BinERROR:

		case *BinTRYRECV:

		case *BinTRYSEND:

		case *BinGOSHED:

		default:
			return nil, NewStringError(stmt, "Неизвестная инструкция")
		}
		idx++
	}
	return rv, nil
}
