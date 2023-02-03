package goafm

import (
	"encoding/hex"
	"strings"
)

// KernPair defines the KernPair data structure
type KernPair struct {
	// First is the first character of the kern pair
	First string
	// Second is the second character of the kern pair
	Second string
	// DX is the horizontal adjustment for the kern pair
	DX Number
	// DY is the vertical adjustment for the kern pair
	DY Number
}

// parseKernPair parse a line into a KernPair
func parseKernPair(line []byte) *KernPair {
	kp := new(KernPair)

	fields := Fields(string(line))
	if len(fields) < 4 {
		return nil
	}
	switch fields[0] {
	case "KPX":
		kp.First = fields[1]
		kp.Second = fields[2]
		kp.DX = NewNumber(fields[3])
		kp.DY = Number(0)
	case "KPY":
		kp.First = fields[1]
		kp.Second = fields[2]
		kp.DX = Number(0)
		kp.DY = NewNumber(fields[3])
	case "KP":
		kp.First = fields[1]
		kp.Second = fields[2]
		kp.DX = NewNumber(fields[3])
		kp.DY = NewNumber(fields[4])
	case "KPH":
		firstHex := strings.Trim(fields[1], "<>")
		secondHex := strings.Trim(fields[2], "<>")
		first, _ := hex.DecodeString(firstHex)
		second, _ := hex.DecodeString(secondHex)
		kp.First = string(first)
		kp.Second = string(second)
		kp.DX = NewNumber(fields[3])
		kp.DY = NewNumber(fields[4])
	}

	return kp
}
