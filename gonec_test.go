package gonec

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_interpreter_ParseAndRun(t *testing.T) {
	type args struct {
		r io.Reader
	}

	ti1 := Interpreter()
	in1 := strings.NewReader(`
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

			;

		КонецЕсли
		`)

	tests := []struct {
		name    string
		i       *interpreter
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "Тест 1",
			i:    ti1,
			args: args{r: in1},
			wantW: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.i.ParseAndRun(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("interpreter.ParseAndRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("interpreter.ParseAndRun() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
