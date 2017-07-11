package parser

import (
	"reflect"
	"testing"

	"github.com/covrom/gonec/gonecparser/ast"
	"github.com/covrom/gonec/gonecparser/token"
)

func TestParseFile(t *testing.T) {
	type args struct {
		fset     *token.FileSet
		filename string
		src      interface{}
		mode     Mode
	}

	fset := token.NewFileSet()

	tests := []struct {
		name    string
		args    args
		wantF   *ast.File
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				fset:     fset,
				filename: "",
				src: `
	// This is scanned code.
	Модуль Основной
	
	Если a > 10 then
		
		someParsable = text;
		раз.два();
		
		если "а=б" тогда;
		(5.26+("4-3.25"));
		б<>а
		в<=г
		д>=е

		a=b=r
		
		Дата('01010001');
		
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
			б=а
		КонецФункции

		Плохой запрос = "ВЫБРАТЬ *
		|ГДЕ

		//нет хвоста
		;

	КонецЕсли`,
				mode: Trace,
			},
		},
		
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, err := ParseFile(tt.args.fset, tt.args.filename, tt.args.src, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("ParseFile() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}
