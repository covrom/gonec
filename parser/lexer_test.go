package parser

import (
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
					а="узцкещпоцз"[1:]
					`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseSrc(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSrc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
