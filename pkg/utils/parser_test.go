package utils

import "testing"

func TestParseText(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Parsea correctamente un string normal",
			args: args{
				[]byte("Probando un string normal"),
			},
			want:    "Probando un string normal",
			wantErr: false,
		},
		{
			name: "Parsea correctamente un string con un acento",
			args: args{
				[]byte("Probando un string con un papá"),
			},
			want:    "Probando un string con un papá",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseText(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
