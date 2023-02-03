package goafm

import "fmt"

// Array is the implementation of an array in the AFM specification.
type Array struct {
	A []interface{}
}

// String returns a string representation of Array.
func (a *Array) String() string {
	r := "[ "
	for _, v := range a.A {
		s, ok := v.(fmt.Stringer)
		if ok {
			r += s.String() + " "
		} else {
			r += fmt.Sprintf("%v ", v)
		}
	}
	r += "]"
	return r
}

// Append adds a value to the Array.
func (a *Array) Append(v interface{}) {
	a.A = append(a.A, v)
}
