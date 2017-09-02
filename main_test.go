package main

import (
	"log"
	"testing"

	"github.com/covrom/gonec/bincode"
	gonec_core "github.com/covrom/gonec/builtins"
	envir "github.com/covrom/gonec/env"
	"github.com/covrom/gonec/parser"
)

func TestRun(t *testing.T) {
	env := envir.NewEnv()
	gonec_core.LoadAllBuiltins(env)

	script := `
	Модуль ААА
	а = 1
	
	Модуль _

	Функция ТрехкратныйВозврат()
    	абв = 0
    	Возврат 10.5, абв, ААА.а
	КонецФункции

	п1, п2, п3 = ТрехкратныйВозврат()
	сообщить(п1,п2,п3)
	
	`
	parser.EnableErrorVerbose()
	_, stmts, err := parser.ParseSrc(script)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bincode.Run(stmts, env)
	if err != nil {
		log.Fatal(err)
	}
}
