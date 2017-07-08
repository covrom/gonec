package gonecscan

import (
	"fmt"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	const src = `
	// This is scanned code.
	Если a > 10 then
		
		someParsable = text;
		раз.два();
		
		если "а=б" тогда;
		(5+("4-3"));
		
		Дата('01010001');
		
		Фун("привет \n привет
		|ку");
		
		ТекстЗапроса = "ВЫБРАТЬ *
		|ИЗ Таблица.АБВ
		//комментарий в запросе
		//english comment
		//|ГДЕ аа=б|

		|ИТОГИ ПО а";

		Плохой запрос = "ВЫБРАТЬ *
		|ГДЕ

		//нет хвоста
		;

	КонецЕсли`
	var s Scanner
	s.Filename = "example"
	s.Init(strings.NewReader(src))
	var tok rune
	for tok != EOF {
		tok = s.Scan()
		fmt.Println(s.Pos(), ":", s.TokenText())
	}

}
