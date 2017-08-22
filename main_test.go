package main

import (
	"log"
	"testing"

	gonec_core "github.com/covrom/gonec/builtins"
	"github.com/covrom/gonec/parser"
	"github.com/covrom/gonec/vm"
)

func TestRun(t *testing.T) {
	env := vm.NewEnv()
	gonec_core.LoadAllBuiltins(env)

	script := `
	а = Новый("__ФункциональнаяСтруктураТест__",{"ПолеЦелоеЧисло":5,"ПолеСтрока":"srtg"})
	а.ПолеСтрока = "edrgwerg"
	Сообщить(а.ВСтроку())
	`
	parser.EnableErrorVerbose()
	stmts, err := parser.ParseSrc(script)
	if err != nil {
		log.Fatal()
	}

	_, err = vm.Run(stmts, env)
	if err != nil {
		log.Fatal()
	}
}
