package goafm

// Boolean represents a boolean value in the AFM specification.
type Boolean bool

// NewBoolean creates a new Boolean with the value represented by the string.
func NewBoolean(b string) Boolean {
	return Boolean(b == "true")
}
