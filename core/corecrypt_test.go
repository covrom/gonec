package core

import (
	"reflect"
	"testing"
)

func TestGZip(t *testing.T) {
	type args struct {
		src []byte
	}

	smpl := []byte(`{
		"id":соед.Идентификатор(),
		"query":"Запрос по tcp протоколу",
		"num":1,
		}`)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "gzip 1",
			args: args{
				src: smpl,
			},
			want:    smpl,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotz, err := GZip(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GZip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(len(gotz), gotz)
			gotuz, err := UnGZip(gotz)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnGZip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(len(gotuz), gotuz)
			if !reflect.DeepEqual(gotuz, tt.want) {
				t.Errorf("UnGZip(Gzip()) = %v, want %v", gotuz, tt.want)
			}
		})
	}
}

func TestEncryptAES128(t *testing.T) {
	type args struct {
		plaintext []byte
	}
	smpl := []byte(`{
		"id":соед.Идентификатор(),
		"query":"Запрос по tcp протоколу",
		"num":1,
		}`)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "gzip 1",
			args: args{
				plaintext: smpl,
			},
			want:    smpl,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptAES128(tt.args.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(len(got), got)
			got2, err := DecryptAES128(got)
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("DecryptAES128(EncryptAES128()) = %v, want %v", got2, tt.want)
			}
		})
	}
}
