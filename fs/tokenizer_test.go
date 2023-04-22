package fs

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{input: "hello world 1 2 3"},
			want:    []string{"hello", "world", "1", "2", "3"},
			wantErr: false,
		},
		{
			name:    "quoted",
			args:    args{input: "hello 'world 1 2 3' nice"},
			want:    []string{"hello", "world 1 2 3", "nice"},
			wantErr: false,
		},
		{
			name:    "nested quote",
			args:    args{input: "hello 'world \"nice\" now' but wait"},
			want:    []string{"hello", "world \"nice\" now", "but", "wait"},
			wantErr: false,
		},
		{
			name:    "quote not ended",
			args:    args{input: "hello 'world \"nice\" now"},
			want:    []string{"hello", "world \"nice\" now"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
