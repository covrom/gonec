package bincode

import (
	"fmt"
	"testing"
)

func TestParseSrc(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Свертка констант",
			args: args{
				src: `
				а = Неопределено
				сообщить(а=123)
				//а="узцкещпоцз"[1:]
					`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, bins, err := ParseSrc(tt.args.src)
			fmt.Println(bins)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSrc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
