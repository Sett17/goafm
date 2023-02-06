package goafm

import (
	"bytes"
	"fmt"
	"strconv"
)

// CharMetric defines the structure of a character metric in an AFM file.
// It holds information about a character's code, width, bounding box, name, and ligature sequence.
type CharMetric struct {
	C   int       // C is the character code. REQUIRED.
	WX  Number    // WX is the character width. REQUIRED.
	W0X Number    // W0X is the character width in x for direction 0.
	W1X Number    // W1X is the character width in x for direction 1.
	WY  Number    // WY is the character width in y.
	W0Y Number    // W0Y is the character width in y for direction 0.
	W1Y Number    // W1Y is the character width in y for direction 1.
	W   [2]Number // W is the character width vector.
	W0  [2]Number // W0 is the character width vector for direction 0.
	W1  [2]Number // W1 is the character width vector for direction 1.
	VV  [2]Number // VV has the same meaning as VVector in the global font program information, but for a single character.
	N   string    // N is the character name.
	B   [4]Number // B is the bounding box.
	L   [2]string // L is the ligature sequence.
}

func (c CharMetric) String() string {
	return fmt.Sprintf("'%c', 0x%X, %s: {WX: %d, B: %v}", c.C, c.C, c.N, c.WX.GetInt(), c.B)
}

// parseCharMetric parses a line of AFM file data and returns a CharMetric struct.
func parseCharMetric(line []byte) *CharMetric {
	cm := new(CharMetric)
	split := bytes.Split(line, []byte(";"))
	for _, s := range split {
		t := bytes.Trim(s, " 	\n")
		if len(t) == 0 {
			continue
		}
		field := Fields(string(t))
		if len(field) < 2 {
			continue
		}
		switch field[0] {
		case "C":
			cm.C, _ = strconv.Atoi(field[1])
		case "CH":
			i, _ := strconv.ParseInt(field[1], 16, 32)
			cm.C = int(i)
		case "WX":
			cm.WX = NewNumber(field[1])
		case "W0X":
			cm.W0X = NewNumber(field[1])
		case "W1X":
			cm.W1X = NewNumber(field[1])
		case "WY":
			cm.WY = NewNumber(field[1])
		case "W0Y":
			cm.W0Y = NewNumber(field[1])
		case "W1Y":
			cm.W1Y = NewNumber(field[1])
		case "W":
			cm.W[0] = NewNumber(field[1])
			cm.W[1] = NewNumber(field[2])
		case "W0":
			cm.W0[0] = NewNumber(field[1])
			cm.W0[1] = NewNumber(field[2])
		case "W1":
			cm.W1[0] = NewNumber(field[1])
			cm.W1[1] = NewNumber(field[2])
		case "VV":
			cm.VV[0] = NewNumber(field[1])
			cm.VV[1] = NewNumber(field[2])
		case "N":
			cm.N = field[1]
		case "B":
			cm.B[0] = NewNumber(field[1])
			cm.B[1] = NewNumber(field[2])
			cm.B[2] = NewNumber(field[3])
			cm.B[3] = NewNumber(field[4])
		case "L":
			cm.L[0] = field[1]
			cm.L[1] = field[2]
		}
	}
	return cm
}
