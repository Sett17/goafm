package goafm

import "testing"

func TestNewNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want Number
	}{
		{"1", args{"1"}, 1},
		{"1.0", args{"1.0"}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNumber(tt.args.num); got != tt.want {
				t.Errorf("NewNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_GetFloat(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want float64
	}{
		{"1", 1, 1},
		{"1.0", 1.0, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.GetFloat(); got != tt.want {
				t.Errorf("GetFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_GetInt(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want int
	}{
		{"1", 1, 1},
		{"1.0", 1.0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.GetInt(); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_IsInt(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want bool
	}{
		{"1", 1, true},
		{"1.5", 1.5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsInt(); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetFloat(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		n    Number
		args args
	}{
		{"1", 1, args{1.0}},
		{"1.0", 1.0, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.SetFloat(tt.args.f)
		})
	}
}

func TestNumber_SetInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    Number
		args args
	}{
		{"1", 1, args{1}},
		{"1.0", 1.0, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.SetInt(tt.args.i)
		})
	}
}
