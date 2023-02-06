package goafm

import (
	"bytes"
	"os"
	"strconv"
	"strings"
)

// FontMetric from an AFM file provides both global metrics for a font program
// and the metrics of each individual character.
type FontMetric struct {
	FontName             string    // (Required) Name of the font program as presented to the PostScript language `findfont` operator.
	FullName             string    // (Optional) The full text name of the font.
	FamilyName           string    // (Optional) The name of the typeface family to which the font belongs.
	Weight               string    // (Optional) Weight of the font.
	FontBBox             [4]Number // Four numbers giving the lower left corner and the upper right corner of the font bounding box, in the sequence llx lly urx ury.
	Version              string    // (Optional) Font program version identifier. Matches the string found in the FontInfo dictionary of the font program itself.
	Notice               string    // (Optional) Font name trademark or copyright notice.
	EncodingScheme       string    // (Optional) String indicating the default encoding vector for this font program.
	MappingScheme        int       // (Not present with base font programs.) Integer code describing the mapping scheme.
	EscapeChar           int       // (Required if MappingScheme 3, not present otherwise). The byte value of the escape character used for this escape-mapped font program.
	CharacterSet         string    // (Optional) String describing the character set (glyph complement) of this font program.
	Characters           int       // (Optional) The number of characters defined in this font program.
	IsBaseFont           Boolean   // (Optional) boolean is true if this font program is a base font and false otherwise.
	VVector              [2]Number // (Required when MetricsSets 2) Components of a vector from origin 0 (the origin for writing direction 0) to origin 1 (the origin for writing direction 1).
	IsFixedV             Boolean   // (Optional) If boolean is true, this indicates that VVector is the same for every character in this font.
	IsCIDFont            Boolean   // (Required if AFM is for a CID-keyed font) If the boolean is true, the font is a CID-keyed font, and the metrics are in CID number order
	CapHeight            Number    // (Optional) Usually the y-value of the top of the capital H. If this font program contains no capital H, this keyword might be missing or numbermight be 0.
	XHeight              Number    // (Optional) Typically the y-value of the top of the lowercase x. If this font program contains no lowercase x, this keyword might be missing or numbermight be 0.
	Ascender             Number    // (Optional) For roman font programs: usually the y-value of the top of the lowercase d. If this font program contains no lowercase d, this keyword mightbe missing or number might be 0.
	Descender            Number    // (Optional) For roman font programs: typically the y-value of the bottom of the lowercase p. If this font program contains no lowercase p, this keywordmight be missing or number might be 0.
	StdHW                Number    // (Optional) This number specifies the dominant width of horizontal stems
	StdVW                Number    // (Optional) This number specifies the dominant width of vertical stems
	BlendAxisTypes       Array     // (Required) The value of BlendAxisTypes is an array of strings that specify the name of each axis in the order in which the master designs are organizedin the multiple master font program.
	BlendDesignPositions Array     // (Required) The value of BlendDesignPositions is an array of k arrays that give the locations of the k master designs in the design space. E
	BlendDesignMap       Array     // (Required) The value of BlendDesignMap is an array of n arrays where n is the number of design axes contained in the multiple master font program.
	WeightVector         Array     // (Required) The WeightVector array specifies the factors for deriving a weighted average of the master designs in a multiple master font program.
	UnderlinePosition    Number    // (Optional) Distance from the baseline for centering underlining strokes.
	UnderlineThickness   Number    // (Optional) This is the stroke width for underlining, and is generally proportional to the stroke widths of characters in the font program.
	ItalicAngle          Number    // (Optional) Angle (in degrees counter-clockwise from the vertical) of the dominant vertical strokes of the font.
	CharWidth            [2]Number // (Optional) The x and y components of the width vector of this font program’s characters.
	IsFixedPitch         Boolean   // (Optional) If boolean is true, this indicates that the font program is a fixed pitch (monospaced) font.

	CharMetricsByCode map[int]*CharMetric
	CharMetricsByName map[string]*CharMetric

	KernPair         map[string]map[string]*KernPair
	KernPairVertical map[string]map[string]*KernPair

	KernTracks map[int]*KernTrack

	Composites map[string]*Composite
}

// MetricByCode returns the CharMetric of the font metric by code.
func (fm *FontMetric) MetricByCode(code int) *CharMetric {
	return fm.CharMetricsByCode[code]
}

// MetricByName returns the CharMetric of the font metric by name.
func (fm *FontMetric) MetricByName(name string) *CharMetric {
	return fm.CharMetricsByName[name]
}

// MetricByRune returns the CharMetric of the font metric by the rune representation.
func (fm *FontMetric) MetricByRune(r rune) *CharMetric {
	ret := fm.CharMetricsByCode[int(r)]
	if ret == nil {
		ret = fm.CharMetricsByName[AGLFN[(uint16(r))].short]
	}
	return ret
}

var lines [][]byte

const (
	FontMetrics = iota
	CharMetrics
	KernData
	KernPairsHorizontal
	KernPairsVertical
	KernTracking
	Composites
)

type genericLine struct {
	keyword string
	args    []string
}

func (g genericLine) parse(line []byte) genericLine {
	fields := bytes.Fields(line)
	g.keyword = string(fields[0])
	g.args = make([]string, len(fields)-1)
	for i, arg := range fields[1:] {
		g.args[i] = string(arg)
	}
	return g
}

// ParseFile takes in a filename and returns a FontMetric pointer and an error.
func ParseFile(filename string) (*FontMetric, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Parse(f)
}

// Parse takes in a byte slice and returns a FontMetric pointer and an error.
func Parse(b []byte) (*FontMetric, error) {
	fm := &FontMetric{}
	fm.CharMetricsByCode = make(map[int]*CharMetric)
	fm.CharMetricsByName = make(map[string]*CharMetric)
	fm.KernPair = make(map[string]map[string]*KernPair)
	fm.KernPairVertical = make(map[string]map[string]*KernPair)
	fm.KernTracks = make(map[int]*KernTrack)
	fm.Composites = make(map[string]*Composite)

	lines = bytes.Split(b, []byte("\n"))

	for i, line := range lines {
		lines[i] = bytes.Trim(line, " 	")
	}

	if !strings.HasPrefix(string(lines[0]), "StartFontMetrics") {
		return nil, NewAFMError(0, "must be 'StartFontMetrics version'")
	}

	var whereAreWe = FontMetrics

	for _, line := range lines[1:] {
		if bytes.HasPrefix(line, []byte("Start")) {
			l := genericLine{}.parse(line)
			switch l.keyword {
			case "StartCharMetrics":
				whereAreWe = CharMetrics
			case "StartKernData":
				whereAreWe = KernData
			case "StartKernPairs":
				whereAreWe = KernPairsHorizontal
			case "StartKernPairs0":
				whereAreWe = KernPairsHorizontal
			case "StartKernPairs1":
				whereAreWe = KernPairsVertical
			case "StartKernTrack":
				whereAreWe = KernTracking
			case "StartComposites":
				whereAreWe = Composites
			default:
				whereAreWe = FontMetrics
			}
			continue
		}
		if bytes.HasPrefix(line, []byte("End")) {
			continue
		}
		if len(line) == 0 {
			continue
		}
		switch whereAreWe {
		case CharMetrics:
			cm := parseCharMetric(line)
			if cm.C >= 0 {
				fm.CharMetricsByCode[cm.C] = cm
			}
			fm.CharMetricsByName[cm.N] = cm
		case KernPairsHorizontal:
			kp := parseKernPair(line)
			if _, ok := fm.KernPair[kp.First]; !ok {
				fm.KernPair[kp.First] = make(map[string]*KernPair)
			}
			fm.KernPair[kp.First][kp.Second] = kp
		case KernPairsVertical:
			kp := parseKernPair(line)
			if _, ok := fm.KernPairVertical[kp.First]; !ok {
				fm.KernPairVertical[kp.First] = make(map[string]*KernPair)
			}
			fm.KernPairVertical[kp.First][kp.Second] = kp
		case KernTracking:
			kt := parseKernTrack(line)
			fm.KernTracks[kt.Degree] = kt
		case Composites:
			c := parseComposite(line)
			fm.Composites[c.Name] = c
		default:
			parseForGlobal(fm, line)
		}
	}

	return fm, nil
}

func parseForGlobal(fm *FontMetric, line []byte) {
	l := genericLine{}.parse(line)
	switch l.keyword {
	case "FontName":
		fm.FontName = strings.Join(l.args, " ")
	case "FullName":
		fm.FullName = strings.Join(l.args, " ")
	case "FamilyName":
		fm.FamilyName = strings.Join(l.args, " ")
	case "Weight":
		fm.Weight = l.args[0]
	case "FontBBox":
		fm.FontBBox = [4]Number{}
		for i, arg := range l.args {
			fm.FontBBox[i] = NewNumber(arg)
		}
	case "Version":
		fm.Version = l.args[0]
	case "Notice":
		fm.Notice = l.args[0]
	case "EncodingScheme":
		fm.EncodingScheme = l.args[0]
	case "MappingScheme":
		fm.MappingScheme, _ = strconv.Atoi(l.args[0])
	case "EscapeChar":
		fm.EscapeChar, _ = strconv.Atoi(l.args[0])
	case "CharacterSet":
		fm.CharacterSet = l.args[0]
	case "Characters":
		fm.Characters, _ = strconv.Atoi(l.args[0])
	case "IsBaseFont":
		fm.IsBaseFont = NewBoolean(l.args[0])
	case "VVector":
		fm.VVector = [2]Number{}
		for i, arg := range l.args {
			fm.VVector[i] = NewNumber(arg)
		}
	case "IsFixedV":
		fm.IsFixedV = NewBoolean(l.args[0])
	case "IsCIDFont":
		fm.IsCIDFont = NewBoolean(l.args[0])
	case "CapHeight":
		fm.CapHeight = NewNumber(l.args[0])
	case "XHeight":
		fm.XHeight = NewNumber(l.args[0])
	case "Ascender":
		fm.Ascender = NewNumber(l.args[0])
	case "Descender":
		fm.Descender = NewNumber(l.args[0])
	case "StdHW":
		fm.StdHW = NewNumber(l.args[0])
	case "StdVW":
		fm.StdVW = NewNumber(l.args[0])
	case "UnderlinePosition":
		fm.UnderlinePosition = NewNumber(l.args[0])
	case "BlendAxisTypes":
		// TODO
	case "BlendDesignPositions":
		// TODO
	case "BlendDesignMap":
		// TODO
	case "WeightVector":
		// TODO
	}
}
