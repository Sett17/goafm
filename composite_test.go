package goafm

import (
	"reflect"
	"testing"
)

func Test_parseComposite(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name string
		args args
		want *Composite
	}{
		{"simple", args{[]byte("CC Aacute 2; PCC A 0 0; PCC acute 194 214;")}, &Composite{"Aacute", []CompositePart{{"A", 0, 0}, {"acute", 194, 214}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseComposite(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseComposite() = %v, want %v", got, tt.want)
			}
		})
	}
}
