package goafm

import (
	"strings"
)

// Composite represents a structure containing information about a composite
type Composite struct {
	// Name is the name of the composite
	Name string
	// Parts is a slice of CompositePart instances
	Parts []CompositePart
}

// CompositePart represents a single part in a composite
type CompositePart struct {
	// Name is the name of the part
	Name string
	// Dx is the x-dimension of the part
	Dx Number
	// Dy is the y-dimension of the part
	Dy Number
}

// parseComposite takes a line of bytes as input and returns a pointer to a new
// Composite instance.
func parseComposite(line []byte) *Composite {
	c := new(Composite)

	split := strings.Split(string(line), ";")

	for _, p := range split {
		if len(p) == 0 {
			continue
		}

		fields := Fields(p)

		switch fields[0] {
		case "CC":
			c.Name = fields[1]
		case "PCC":
			part := CompositePart{
				Name: fields[1],
				Dx:   NewNumber(fields[2]),
				Dy:   NewNumber(fields[3]),
			}
			c.Parts = append(c.Parts, part)
		}
	}

	return c
}
