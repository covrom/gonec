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
	x = 5;
	y = 10;
	foobar = 838383;
	a=b=r;
		
	f = Дата('01010001') 
		+
				Фун("привет \n привет
		|ку");
		
		ТекстЗапроса = "ВЫБРАТЬ *
		|ИЗ Таблица.АБВ
		//комментарий в запросе
		//english comment
		//|ГДЕ аа=б|

		|ИТОГИ ПО а //коммент внутри не удаляем
		|";

		Функция а()
			а=б
		КонецФункции

		Функция б()
			б=а=акгр=дыврам
		КонецФункции
	bad =

	`)
	w := &bytes.Buffer{}
	err := ti1.ParseAndRun(in1, w)
	fmt.Println(w.String())
	if err != nil {
		fmt.Println(err.Error())
	}
}
