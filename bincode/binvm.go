// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	envir "github.com/covrom/gonec/env"
)

func Interrupt(env *envir.Env) {
	env.Interrupt()
}

func Run(stmts BinCode, env *envir.Env) (rv interface{}, err error) {
	for _, stmt := range stmts {
		if _, ok := stmt.(*BinBREAK); ok {
			return nil, BreakError
		}
		if _, ok := stmt.(*BinCONTINUE); ok {
			return nil, ContinueError
		}
		rv, err = RunSingleStmt(stmt, env)
		if err != nil {
			return rv, err
		}
		if _, ok := stmt.(*BinRET); ok {
			return rv, ReturnError
		}
	}
	return rv, nil
}

func RunSingleStmt(stmt BinStmt, env *envir.Env) (interface{}, error) {
	if env.CheckInterrupt() {
		return nil, InterruptError
	}
	// подготавливаем состояние машины: регистры значений, управляющие регистры
	regs := make([]interface{}, 20)
	lenregs := len(regs)

	switch s := stmt.(type) {
	case *BinLOAD:
		if s.Reg < lenregs {
			regs[s.Reg] = s.Val
		} else {
			for i := lenregs; i <= s.Reg; i++ {
				regs = append(regs, nil)
			}
			lenregs = len(regs)
			regs[s.Reg] = s.Val
		}

	}
	return nil, NewStringError(stmt, "Неизвестная инструкция")
}
