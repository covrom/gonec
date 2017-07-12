package parser

import (
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
	Пакет Основной

		Функция а(б,в,г) экспОрт
			а=б
		КонецФункции

	`,
				mode: Trace,
			},
		},

		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseFile(tt.args.fset, tt.args.filename, tt.args.src, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
		})
	}
}
