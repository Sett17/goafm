package goafm

import (
	"reflect"
	"testing"
)

func Test_parseForKernPair(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name string
		args args
		want *KernPair
	}{
		{"KPX", args{[]byte("KPX A B 1")}, &KernPair{"A", "B", Number(1), Number(0)}},
		{"KPY", args{[]byte("KPY A B 1")}, &KernPair{"A", "B", Number(0), Number(1)}},
		{"KP", args{[]byte("KP A B 1 2")}, &KernPair{"A", "B", Number(1), Number(2)}},
		{"KPH", args{[]byte("KPH <41> <42> 1 2")}, &KernPair{"A", "B", Number(1), Number(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKernPair(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseForKernPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
