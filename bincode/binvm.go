// Package bincode - виртуальная машина исполнения байткода
package bincode

import (
	"fmt"

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

	switch s := stmt.(type) {
	case *BinLOAD:
		//
		fmt.Println(s)

	}
	return nil, NewStringError(stmt, "Неизвестная инструкция")

}
