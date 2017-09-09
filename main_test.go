package main

import (
	"log"
	"testing"

	"github.com/covrom/gonec/bincode"
	envir "github.com/covrom/gonec/env"
	"github.com/covrom/gonec/parser"
)

func TestRun(t *testing.T) {
	env := envir.NewEnv()

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

	функция фиб(н)
		если н = 0 тогда
			возврат 0
		иначеесли н = 1 тогда
			возврат 1
		конецесли
		возврат фиб(н-1) + фиб(н-2)
  	конецфункции
  
  	сообщить(фиб(10))
	
	функция фибт(н0, н1, к)
		если к = 0 тогда
			возврат н0
	  	иначеесли к = 1 тогда
			возврат н1
	  	конецесли
	  	возврат фибт(н1, н0+н1, к-1)
	конецфункции
		
	функция фиб2(н)
		возврат фибт(0, 1, н)
	конецфункции
  
  	сообщить(фиб2(10))
	`
	parser.EnableErrorVerbose()
	_, stmts, err := bincode.ParseSrc(script)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bincode.Run(stmts, env)
	if err != nil {
		log.Fatal(err)
	}
}
