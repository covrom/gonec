// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	envir "github.com/covrom/gonec/env"
)

func Interrupt(env *envir.Env) {
	env.Interrupt()
}

func Run(stmts BinCode, env *envir.Env) (rv interface{}, err error) {
	// подготавливаем состояние машины: регистры значений, управляющие регистры
	regs := NewVMRegs(stmts)

	var idx int
	for idx < len(stmts) {
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
		default:
			rv, err = RunSingleStmt(s, env, regs)
			if err != nil {
				return rv, err
			}
			if _, ok := stmt.(*BinRET); ok {
				return rv, ReturnError
			}
		}
		idx++
	}
	return rv, nil
}

func RunSingleStmt(stmt BinStmt, env *envir.Env, regs *VMRegs) (interface{}, error) {
	if env.CheckInterrupt() {
		return nil, InterruptError
	}

	switch s := stmt.(type) {
	case *BinLOAD:
		regs.Set(s.Reg, s.Val)
	case *BinMV:

	case *BinEQUAL:

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

	case *BinTHROW:

	case *BinMODULE:

	case *BinERROR:

	case *BinTRYRECV:

	case *BinTRYSEND:

	case *BinGOSHED:

	}
	return nil, NewStringError(stmt, "Неизвестная инструкция")
}
