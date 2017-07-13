package gonec

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_interpreter_ParseAndRun(t *testing.T) {
	ti1 := Interpreter()
	in1 := strings.NewReader(`
	// This is scanned code.
	Пакет Основной

		Функция а(б,в,г) экспОрт
			а=б
		КонецФункции

		Процедура Проверка(проверяемое)
				Сообщить(проверяемое=0)
				БезТочкиСЗапятой()
				СТочкойСЗапятой();
		КонецПроцедуры

		б = а(1,2,3)
		Сообщить(б)
	`)
	w := &bytes.Buffer{}
	err := ti1.ParseAndRun(in1, w)
	fmt.Println(w.String())
	if err != nil {
		fmt.Println(err.Error())
	}
}
