package goafm

import (
	"fmt"
)

// AFMError represents an error in the AFM file.
type AFMError struct {
	Line int
	Msg  string
}

// NewAFMError creates a new AFMError with the given line number and error message.
func NewAFMError(line int, format string, a ...any) *AFMError {
	return &AFMError{Line: line, Msg: fmt.Sprintf(format, a...)}
}

// Error implements the error interface for AFMError.
func (A AFMError) Error() string {
	return fmt.Sprintf("invalid AFM file @ %d: %s\nline %d: %s", A.Line, A.Msg, A.Line, string(lines[A.Line]))
}
