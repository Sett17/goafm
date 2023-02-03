package goafm

import (
	"strconv"
)

// KernTrack holds information about kerning for a font with a specific degree.
type KernTrack struct {
	// Degree holds the degree of the kerning.
	Degree int
	// MinPt holds the minimum point size for the kerning.
	MinPt Number
	// MinKern holds the minimum kerning for the minimum point size.
	MinKern Number
	// MaxPt holds the maximum point size for the kerning.
	MaxPt Number
	// MaxKern holds the maximum kerning for the maximum point size.
	MaxKern Number
}

// parseKernTrack parses a string into a KernTrack.
func parseKernTrack(line []byte) *KernTrack {
	kt := new(KernTrack)

	fields := Fields(string(line))
	kt.Degree, _ = strconv.Atoi(fields[1])
	kt.MinPt = NewNumber(fields[2])
	kt.MinKern = NewNumber(fields[3])
	kt.MaxPt = NewNumber(fields[4])
	kt.MaxKern = NewNumber(fields[5])

	return kt
}

// Fn returns the kerning for a given point size.
func (kt *KernTrack) Fn(pt Number) Number {
	if pt < kt.MinPt {
		return kt.MinKern
	}
	if pt > kt.MaxPt {
		return kt.MaxKern
	}
	return ((kt.MaxKern-kt.MinKern)/(kt.MaxPt-kt.MinPt))*(pt-kt.MinPt) + kt.MinKern
}
