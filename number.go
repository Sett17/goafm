// Package goafm provides a custom type Number which wraps float64 value.
package goafm

import "strconv"

// Number type wraps float64 value.
type Number float64

// NewNumber creates a new Number type from a string representation of a float64.
func NewNumber(num string) Number {
	f, _ := strconv.ParseFloat(num, 64)
	return Number(f)
}

// SetFloat sets the float64 value of a Number type.
func (n *Number) SetFloat(f float64) {
	*n = Number(f)
}

// SetInt sets the int value of a Number type.
func (n *Number) SetInt(i int) {
	*n = Number(i)
}

// GetFloat returns the float64 value of a Number type.
func (n *Number) GetFloat() float64 {
	return float64(*n)
}

// GetInt returns the int value of a Number type.
func (n *Number) GetInt() int {
	return int(*n)
}

// IsInt returns a boolean indicating whether the Number type holds an int value.
func (n *Number) IsInt() bool {
	return float64(*n) == float64(int(*n))
}
