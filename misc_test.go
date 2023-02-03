package goafm

import (
	"reflect"
	"strings"
	"testing"
)

func BenchmarkStringFields(b *testing.B) {
	str := "KernTrack 2 1 1 10 10"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Fields(str)
	}
	b.ReportAllocs()
}

func BenchmarkCustomFields(b *testing.B) {
	str := "KernTrack 2 1 1 10 10"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fields(str)
	}
	b.ReportAllocs()
}

func TestFields(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"simple", args{"one two three four"}, []string{"one", "two", "three", "four"}},
		{"with tab", args{"one\ttwo\tthree\tfour"}, []string{"one", "two", "three", "four"}},
		{"mixed", args{"one two\tthree four"}, []string{"one", "two", "three", "four"}},
		{"empty", args{""}, []string{}},
		{"empty2", args{"   "}, []string{}},
		{"single", args{"one"}, []string{"one"}},
		{"suffix spaces", args{"one two three four   "}, []string{"one", "two", "three", "four"}},
		{"prefix spaces", args{"   one two three four"}, []string{"one", "two", "three", "four"}},
		{"spaces", args{"   one two three four   "}, []string{"one", "two", "three", "four"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fields(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}
