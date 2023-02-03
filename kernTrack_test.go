package goafm

import (
	"reflect"
	"testing"
)

func TestKernTrack_Fn(t *testing.T) {
	type fields struct {
		Degree  int
		MinPt   Number
		MinKern Number
		MaxPt   Number
		MaxKern Number
	}
	type args struct {
		p Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		{"min", fields{1, 0, 0, 1, 1}, args{0}, 0},
		{"max", fields{1, 0, 0, 1, 1}, args{1}, 1},
		{"mid", fields{1, 0, 0, 1, 100}, args{0.5}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kt := &KernTrack{
				Degree:  tt.fields.Degree,
				MinPt:   tt.fields.MinPt,
				MinKern: tt.fields.MinKern,
				MaxPt:   tt.fields.MaxPt,
				MaxKern: tt.fields.MaxKern,
			}
			if got := kt.Fn(tt.args.p); got != tt.want {
				t.Errorf("Fn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseKernTrack(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name string
		args args
		want *KernTrack
	}{
		{"ints", args{[]byte("TrackKern 1 0 0 1 1")}, &KernTrack{1, 0, 0, 1, 1}},
		{"floats", args{[]byte("TrackKern 1 0.0 0.0 1.0 1.0")}, &KernTrack{1, 0, 0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKernTrack(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKernTrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
