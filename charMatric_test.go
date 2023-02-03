package goafm

import (
	"reflect"
	"testing"
)

func Test_parseCharMetric(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name string
		args args
		want *CharMetric
	}{
		{name: "simple", args: args{[]byte("C 32 ; WX 250 ; N space ; B 0 0 0 0 ;")}, want: &CharMetric{C: 32, WX: 250, N: "space", B: [4]Number{0, 0, 0, 0}}},
		{"real", args{[]byte("C 37 ; WX 833 ; N percent ; B 34 -14 799 666 ;")}, &CharMetric{C: 37, WX: 833, N: "percent", B: [4]Number{34, -14, 799, 666}}},
		{"L", args{[]byte("C 37 ; WX 833 ; N percent ; B 34 -14 799 666 ; L i fi ;")}, &CharMetric{C: 37, WX: 833, N: "percent", B: [4]Number{34, -14, 799, 666}, L: [2]string{"i", "fi"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCharMetric(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCharMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
