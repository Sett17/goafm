package goafm

import "testing"

func TestArray_Append(t *testing.T) {
	type fields struct {
		A []interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{"simple", fields{[]interface{}{"a", "b", "c"}}, args{"d"}, fields{[]interface{}{"a", "b", "c", "d"}}},
		{"empty", fields{[]interface{}{}}, args{"d"}, fields{[]interface{}{"d"}}},
		{"mixed", fields{[]interface{}{"a", "b", "c", 1, 2, 3}}, args{"d"}, fields{[]interface{}{"a", "b", "c", 1, 2, 3, "d"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				A: tt.fields.A,
			}
			a.Append(tt.args.v)
		})
	}
}

func TestArray_String(t *testing.T) {
	type fields struct {
		A []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"simple", fields{[]interface{}{"a", "b", "c"}}, "[ a b c ]"},
		{"empty", fields{[]interface{}{}}, "[ ]"},
		{"mixed", fields{[]interface{}{"a", "b", "c", 1, 2, 3}}, "[ a b c 1 2 3 ]"},
		{"subarray", fields{[]interface{}{"a", "b", "c", &Array{[]interface{}{"d", "e", "f"}}}}, "[ a b c [ d e f ] ]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				A: tt.fields.A,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
