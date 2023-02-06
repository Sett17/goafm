package goafm
// Code generated by aglfn2go.ipynb. DO NOT EDIT.

// -----------------------------------------------------------
// Copyright 2002-2019 Adobe (http://www.adobe.com/).
//
// Redistribution and use in source and binary forms, with or
// without modification, are permitted provided that the
// following conditions are met:
//
// Redistributions of source code must retain the above
// copyright notice, this list of conditions and the following
// disclaimer.
//
// Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following
// disclaimer in the documentation and/or other materials
// provided with the distribution.
//
// Neither the name of Adobe nor the names of its contributors
// may be used to endorse or promote products derived from this
// software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
// CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
// OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// -----------------------------------------------------------
// Name:          Adobe Glyph List For New Fonts
// Table version: 1.7
// Date:          November 6, 2008
// URL:           https://github.com/adobe-type-tools/agl-aglfn
//
// Description:
//
// AGLFN (Adobe Glyph List For New Fonts) provides a list of base glyph
// names that are recommended for new fonts, which are compatible with
// the AGL (Adobe Glyph List) Specification, and which should be used
// as described in Section 6 of that document. AGLFN comprises the set
// of glyph names from AGL that map via the AGL Specification rules to
// the semantically correct UV (Unicode Value). For example, "Asmall"
// is omitted because AGL maps this glyph name to the PUA (Private Use
// Area) value U+F761, rather than to the UV that maps from the glyph
// name "A." Also omitted is "ffi," because AGL maps this to the
// Alphabetic Presentation Forms value U+FB03, rather than decomposing
// it into the following sequence of three UVs: U+0066, U+0066, and
// U+0069. The name "arrowvertex" has been omitted because this glyph
// now has a real UV, and AGL is now incorrect in mapping it to the PUA
// value U+F8E6. If you do not find an appropriate name for your glyph
// in this list, then please refer to Section 6 of the AGL
// Specification.
//
// Format: three semicolon-delimited fields:
//   (1) Standard UV or CUS UV--four uppercase hexadecimal digits
//   (2) Glyph name--upper/lowercase letters and digits
//   (3) Character names: Unicode character names for standard UVs, and
//       descriptive names for CUS UVs--uppercase letters, hyphen, and
//       space
//
// The records are sorted by glyph name in increasing ASCII order,
// entries with the same glyph name are sorted in decreasing priority
// order, the UVs and Unicode character names are provided for
// convenience, lines starting with "#" are comments, and blank lines
// should be ignored.
//
// Revision History:
//
// 1.7 [6 November 2008]
// - Reverted to the original 1.4 and earlier mappings for Delta,
//   Omega, and mu.
// - Removed mappings for "afii" names. These should now be assigned
//   "uni" names.
// - Removed mappings for "commaaccent" names. These should now be
//   assigned "uni" names.
//
// 1.6 [30 January 2006]
// - Completed work intended in 1.5.
//
// 1.5 [23 November 2005]
// - Removed duplicated block at end of file.
// - Changed mappings:
//   2206;Delta;INCREMENT changed to 0394;Delta;GREEK CAPITAL LETTER DELTA
//   2126;Omega;OHM SIGN changed to 03A9;Omega;GREEK CAPITAL LETTER OMEGA
//   03BC;mu;MICRO SIGN changed to 03BC;mu;GREEK SMALL LETTER MU
// - Corrected statement above about why "ffi" is omitted.
//
// 1.4 [24 September 2003]
// - Changed version to 1.4, to avoid confusion with the AGL 1.3.
// - Fixed spelling errors in the header.
// - Fully removed "arrowvertex," as it is mapped only to a PUA Unicode
//   value in some fonts.
//
// 1.1 [17 April 2003]
// - Renamed [Tt]cedilla back to [Tt]commaaccent.
//
// 1.0 [31 January 2003]
// - Original version.
// - Derived from the AGLv1.2 by:
//   removing the PUA area codes;
//   removing duplicate Unicode mappings; and
//   renaming "tcommaaccent" to "tcedilla" and "Tcommaaccent" to "Tcedilla"
//


// NamePair is a pair of glyph name and character name.
type NamePair struct {
	short string
	long  string
}

// contains returns true if the pair contains the given string.
func (np NamePair) contains(s string) bool {
	return np.short == s || np.long == s
}

// AGLFN is the Adobe Glyph List for New Fonts.
var AGLFN = map[uint16]NamePair{
	0x0041: {"A", "LATIN CAPITAL LETTER A"},
	0x00C6: {"AE", "LATIN CAPITAL LETTER AE"},
	0x01FC: {"AEacute", "LATIN CAPITAL LETTER AE WITH ACUTE"},
	0x00C1: {"Aacute", "LATIN CAPITAL LETTER A WITH ACUTE"},
	0x0102: {"Abreve", "LATIN CAPITAL LETTER A WITH BREVE"},
	0x00C2: {"Acircumflex", "LATIN CAPITAL LETTER A WITH CIRCUMFLEX"},
	0x00C4: {"Adieresis", "LATIN CAPITAL LETTER A WITH DIAERESIS"},
	0x00C0: {"Agrave", "LATIN CAPITAL LETTER A WITH GRAVE"},
	0x0391: {"Alpha", "GREEK CAPITAL LETTER ALPHA"},
	0x0386: {"Alphatonos", "GREEK CAPITAL LETTER ALPHA WITH TONOS"},
	0x0100: {"Amacron", "LATIN CAPITAL LETTER A WITH MACRON"},
	0x0104: {"Aogonek", "LATIN CAPITAL LETTER A WITH OGONEK"},
	0x00C5: {"Aring", "LATIN CAPITAL LETTER A WITH RING ABOVE"},
	0x01FA: {"Aringacute", "LATIN CAPITAL LETTER A WITH RING ABOVE AND ACUTE"},
	0x00C3: {"Atilde", "LATIN CAPITAL LETTER A WITH TILDE"},
	0x0042: {"B", "LATIN CAPITAL LETTER B"},
	0x0392: {"Beta", "GREEK CAPITAL LETTER BETA"},
	0x0043: {"C", "LATIN CAPITAL LETTER C"},
	0x0106: {"Cacute", "LATIN CAPITAL LETTER C WITH ACUTE"},
	0x010C: {"Ccaron", "LATIN CAPITAL LETTER C WITH CARON"},
	0x00C7: {"Ccedilla", "LATIN CAPITAL LETTER C WITH CEDILLA"},
	0x0108: {"Ccircumflex", "LATIN CAPITAL LETTER C WITH CIRCUMFLEX"},
	0x010A: {"Cdotaccent", "LATIN CAPITAL LETTER C WITH DOT ABOVE"},
	0x03A7: {"Chi", "GREEK CAPITAL LETTER CHI"},
	0x0044: {"D", "LATIN CAPITAL LETTER D"},
	0x010E: {"Dcaron", "LATIN CAPITAL LETTER D WITH CARON"},
	0x0110: {"Dcroat", "LATIN CAPITAL LETTER D WITH STROKE"},
	0x2206: {"Delta", "INCREMENT"},
	0x0045: {"E", "LATIN CAPITAL LETTER E"},
	0x00C9: {"Eacute", "LATIN CAPITAL LETTER E WITH ACUTE"},
	0x0114: {"Ebreve", "LATIN CAPITAL LETTER E WITH BREVE"},
	0x011A: {"Ecaron", "LATIN CAPITAL LETTER E WITH CARON"},
	0x00CA: {"Ecircumflex", "LATIN CAPITAL LETTER E WITH CIRCUMFLEX"},
	0x00CB: {"Edieresis", "LATIN CAPITAL LETTER E WITH DIAERESIS"},
	0x0116: {"Edotaccent", "LATIN CAPITAL LETTER E WITH DOT ABOVE"},
	0x00C8: {"Egrave", "LATIN CAPITAL LETTER E WITH GRAVE"},
	0x0112: {"Emacron", "LATIN CAPITAL LETTER E WITH MACRON"},
	0x014A: {"Eng", "LATIN CAPITAL LETTER ENG"},
	0x0118: {"Eogonek", "LATIN CAPITAL LETTER E WITH OGONEK"},
	0x0395: {"Epsilon", "GREEK CAPITAL LETTER EPSILON"},
	0x0388: {"Epsilontonos", "GREEK CAPITAL LETTER EPSILON WITH TONOS"},
	0x0397: {"Eta", "GREEK CAPITAL LETTER ETA"},
	0x0389: {"Etatonos", "GREEK CAPITAL LETTER ETA WITH TONOS"},
	0x00D0: {"Eth", "LATIN CAPITAL LETTER ETH"},
	0x20AC: {"Euro", "EURO SIGN"},
	0x0046: {"F", "LATIN CAPITAL LETTER F"},
	0x0047: {"G", "LATIN CAPITAL LETTER G"},
	0x0393: {"Gamma", "GREEK CAPITAL LETTER GAMMA"},
	0x011E: {"Gbreve", "LATIN CAPITAL LETTER G WITH BREVE"},
	0x01E6: {"Gcaron", "LATIN CAPITAL LETTER G WITH CARON"},
	0x011C: {"Gcircumflex", "LATIN CAPITAL LETTER G WITH CIRCUMFLEX"},
	0x0120: {"Gdotaccent", "LATIN CAPITAL LETTER G WITH DOT ABOVE"},
	0x0048: {"H", "LATIN CAPITAL LETTER H"},
	0x25CF: {"H18533", "BLACK CIRCLE"},
	0x25AA: {"H18543", "BLACK SMALL SQUARE"},
	0x25AB: {"H18551", "WHITE SMALL SQUARE"},
	0x25A1: {"H22073", "WHITE SQUARE"},
	0x0126: {"Hbar", "LATIN CAPITAL LETTER H WITH STROKE"},
	0x0124: {"Hcircumflex", "LATIN CAPITAL LETTER H WITH CIRCUMFLEX"},
	0x0049: {"I", "LATIN CAPITAL LETTER I"},
	0x0132: {"IJ", "LATIN CAPITAL LIGATURE IJ"},
	0x00CD: {"Iacute", "LATIN CAPITAL LETTER I WITH ACUTE"},
	0x012C: {"Ibreve", "LATIN CAPITAL LETTER I WITH BREVE"},
	0x00CE: {"Icircumflex", "LATIN CAPITAL LETTER I WITH CIRCUMFLEX"},
	0x00CF: {"Idieresis", "LATIN CAPITAL LETTER I WITH DIAERESIS"},
	0x0130: {"Idotaccent", "LATIN CAPITAL LETTER I WITH DOT ABOVE"},
	0x2111: {"Ifraktur", "BLACK-LETTER CAPITAL I"},
	0x00CC: {"Igrave", "LATIN CAPITAL LETTER I WITH GRAVE"},
	0x012A: {"Imacron", "LATIN CAPITAL LETTER I WITH MACRON"},
	0x012E: {"Iogonek", "LATIN CAPITAL LETTER I WITH OGONEK"},
	0x0399: {"Iota", "GREEK CAPITAL LETTER IOTA"},
	0x03AA: {"Iotadieresis", "GREEK CAPITAL LETTER IOTA WITH DIALYTIKA"},
	0x038A: {"Iotatonos", "GREEK CAPITAL LETTER IOTA WITH TONOS"},
	0x0128: {"Itilde", "LATIN CAPITAL LETTER I WITH TILDE"},
	0x004A: {"J", "LATIN CAPITAL LETTER J"},
	0x0134: {"Jcircumflex", "LATIN CAPITAL LETTER J WITH CIRCUMFLEX"},
	0x004B: {"K", "LATIN CAPITAL LETTER K"},
	0x039A: {"Kappa", "GREEK CAPITAL LETTER KAPPA"},
	0x004C: {"L", "LATIN CAPITAL LETTER L"},
	0x0139: {"Lacute", "LATIN CAPITAL LETTER L WITH ACUTE"},
	0x039B: {"Lambda", "GREEK CAPITAL LETTER LAMDA"},
	0x013D: {"Lcaron", "LATIN CAPITAL LETTER L WITH CARON"},
	0x013F: {"Ldot", "LATIN CAPITAL LETTER L WITH MIDDLE DOT"},
	0x0141: {"Lslash", "LATIN CAPITAL LETTER L WITH STROKE"},
	0x004D: {"M", "LATIN CAPITAL LETTER M"},
	0x039C: {"Mu", "GREEK CAPITAL LETTER MU"},
	0x004E: {"N", "LATIN CAPITAL LETTER N"},
	0x0143: {"Nacute", "LATIN CAPITAL LETTER N WITH ACUTE"},
	0x0147: {"Ncaron", "LATIN CAPITAL LETTER N WITH CARON"},
	0x00D1: {"Ntilde", "LATIN CAPITAL LETTER N WITH TILDE"},
	0x039D: {"Nu", "GREEK CAPITAL LETTER NU"},
	0x004F: {"O", "LATIN CAPITAL LETTER O"},
	0x0152: {"OE", "LATIN CAPITAL LIGATURE OE"},
	0x00D3: {"Oacute", "LATIN CAPITAL LETTER O WITH ACUTE"},
	0x014E: {"Obreve", "LATIN CAPITAL LETTER O WITH BREVE"},
	0x00D4: {"Ocircumflex", "LATIN CAPITAL LETTER O WITH CIRCUMFLEX"},
	0x00D6: {"Odieresis", "LATIN CAPITAL LETTER O WITH DIAERESIS"},
	0x00D2: {"Ograve", "LATIN CAPITAL LETTER O WITH GRAVE"},
	0x01A0: {"Ohorn", "LATIN CAPITAL LETTER O WITH HORN"},
	0x0150: {"Ohungarumlaut", "LATIN CAPITAL LETTER O WITH DOUBLE ACUTE"},
	0x014C: {"Omacron", "LATIN CAPITAL LETTER O WITH MACRON"},
	0x2126: {"Omega", "OHM SIGN"},
	0x038F: {"Omegatonos", "GREEK CAPITAL LETTER OMEGA WITH TONOS"},
	0x039F: {"Omicron", "GREEK CAPITAL LETTER OMICRON"},
	0x038C: {"Omicrontonos", "GREEK CAPITAL LETTER OMICRON WITH TONOS"},
	0x00D8: {"Oslash", "LATIN CAPITAL LETTER O WITH STROKE"},
	0x01FE: {"Oslashacute", "LATIN CAPITAL LETTER O WITH STROKE AND ACUTE"},
	0x00D5: {"Otilde", "LATIN CAPITAL LETTER O WITH TILDE"},
	0x0050: {"P", "LATIN CAPITAL LETTER P"},
	0x03A6: {"Phi", "GREEK CAPITAL LETTER PHI"},
	0x03A0: {"Pi", "GREEK CAPITAL LETTER PI"},
	0x03A8: {"Psi", "GREEK CAPITAL LETTER PSI"},
	0x0051: {"Q", "LATIN CAPITAL LETTER Q"},
	0x0052: {"R", "LATIN CAPITAL LETTER R"},
	0x0154: {"Racute", "LATIN CAPITAL LETTER R WITH ACUTE"},
	0x0158: {"Rcaron", "LATIN CAPITAL LETTER R WITH CARON"},
	0x211C: {"Rfraktur", "BLACK-LETTER CAPITAL R"},
	0x03A1: {"Rho", "GREEK CAPITAL LETTER RHO"},
	0x0053: {"S", "LATIN CAPITAL LETTER S"},
	0x250C: {"SF010000", "BOX DRAWINGS LIGHT DOWN AND RIGHT"},
	0x2514: {"SF020000", "BOX DRAWINGS LIGHT UP AND RIGHT"},
	0x2510: {"SF030000", "BOX DRAWINGS LIGHT DOWN AND LEFT"},
	0x2518: {"SF040000", "BOX DRAWINGS LIGHT UP AND LEFT"},
	0x253C: {"SF050000", "BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL"},
	0x252C: {"SF060000", "BOX DRAWINGS LIGHT DOWN AND HORIZONTAL"},
	0x2534: {"SF070000", "BOX DRAWINGS LIGHT UP AND HORIZONTAL"},
	0x251C: {"SF080000", "BOX DRAWINGS LIGHT VERTICAL AND RIGHT"},
	0x2524: {"SF090000", "BOX DRAWINGS LIGHT VERTICAL AND LEFT"},
	0x2500: {"SF100000", "BOX DRAWINGS LIGHT HORIZONTAL"},
	0x2502: {"SF110000", "BOX DRAWINGS LIGHT VERTICAL"},
	0x2561: {"SF190000", "BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE"},
	0x2562: {"SF200000", "BOX DRAWINGS VERTICAL DOUBLE AND LEFT SINGLE"},
	0x2556: {"SF210000", "BOX DRAWINGS DOWN DOUBLE AND LEFT SINGLE"},
	0x2555: {"SF220000", "BOX DRAWINGS DOWN SINGLE AND LEFT DOUBLE"},
	0x2563: {"SF230000", "BOX DRAWINGS DOUBLE VERTICAL AND LEFT"},
	0x2551: {"SF240000", "BOX DRAWINGS DOUBLE VERTICAL"},
	0x2557: {"SF250000", "BOX DRAWINGS DOUBLE DOWN AND LEFT"},
	0x255D: {"SF260000", "BOX DRAWINGS DOUBLE UP AND LEFT"},
	0x255C: {"SF270000", "BOX DRAWINGS UP DOUBLE AND LEFT SINGLE"},
	0x255B: {"SF280000", "BOX DRAWINGS UP SINGLE AND LEFT DOUBLE"},
	0x255E: {"SF360000", "BOX DRAWINGS VERTICAL SINGLE AND RIGHT DOUBLE"},
	0x255F: {"SF370000", "BOX DRAWINGS VERTICAL DOUBLE AND RIGHT SINGLE"},
	0x255A: {"SF380000", "BOX DRAWINGS DOUBLE UP AND RIGHT"},
	0x2554: {"SF390000", "BOX DRAWINGS DOUBLE DOWN AND RIGHT"},
	0x2569: {"SF400000", "BOX DRAWINGS DOUBLE UP AND HORIZONTAL"},
	0x2566: {"SF410000", "BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL"},
	0x2560: {"SF420000", "BOX DRAWINGS DOUBLE VERTICAL AND RIGHT"},
	0x2550: {"SF430000", "BOX DRAWINGS DOUBLE HORIZONTAL"},
	0x256C: {"SF440000", "BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL"},
	0x2567: {"SF450000", "BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE"},
	0x2568: {"SF460000", "BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE"},
	0x2564: {"SF470000", "BOX DRAWINGS DOWN SINGLE AND HORIZONTAL DOUBLE"},
	0x2565: {"SF480000", "BOX DRAWINGS DOWN DOUBLE AND HORIZONTAL SINGLE"},
	0x2559: {"SF490000", "BOX DRAWINGS UP DOUBLE AND RIGHT SINGLE"},
	0x2558: {"SF500000", "BOX DRAWINGS UP SINGLE AND RIGHT DOUBLE"},
	0x2552: {"SF510000", "BOX DRAWINGS DOWN SINGLE AND RIGHT DOUBLE"},
	0x2553: {"SF520000", "BOX DRAWINGS DOWN DOUBLE AND RIGHT SINGLE"},
	0x256B: {"SF530000", "BOX DRAWINGS VERTICAL DOUBLE AND HORIZONTAL SINGLE"},
	0x256A: {"SF540000", "BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE"},
	0x015A: {"Sacute", "LATIN CAPITAL LETTER S WITH ACUTE"},
	0x0160: {"Scaron", "LATIN CAPITAL LETTER S WITH CARON"},
	0x015E: {"Scedilla", "LATIN CAPITAL LETTER S WITH CEDILLA"},
	0x015C: {"Scircumflex", "LATIN CAPITAL LETTER S WITH CIRCUMFLEX"},
	0x03A3: {"Sigma", "GREEK CAPITAL LETTER SIGMA"},
	0x0054: {"T", "LATIN CAPITAL LETTER T"},
	0x03A4: {"Tau", "GREEK CAPITAL LETTER TAU"},
	0x0166: {"Tbar", "LATIN CAPITAL LETTER T WITH STROKE"},
	0x0164: {"Tcaron", "LATIN CAPITAL LETTER T WITH CARON"},
	0x0398: {"Theta", "GREEK CAPITAL LETTER THETA"},
	0x00DE: {"Thorn", "LATIN CAPITAL LETTER THORN"},
	0x0055: {"U", "LATIN CAPITAL LETTER U"},
	0x00DA: {"Uacute", "LATIN CAPITAL LETTER U WITH ACUTE"},
	0x016C: {"Ubreve", "LATIN CAPITAL LETTER U WITH BREVE"},
	0x00DB: {"Ucircumflex", "LATIN CAPITAL LETTER U WITH CIRCUMFLEX"},
	0x00DC: {"Udieresis", "LATIN CAPITAL LETTER U WITH DIAERESIS"},
	0x00D9: {"Ugrave", "LATIN CAPITAL LETTER U WITH GRAVE"},
	0x01AF: {"Uhorn", "LATIN CAPITAL LETTER U WITH HORN"},
	0x0170: {"Uhungarumlaut", "LATIN CAPITAL LETTER U WITH DOUBLE ACUTE"},
	0x016A: {"Umacron", "LATIN CAPITAL LETTER U WITH MACRON"},
	0x0172: {"Uogonek", "LATIN CAPITAL LETTER U WITH OGONEK"},
	0x03A5: {"Upsilon", "GREEK CAPITAL LETTER UPSILON"},
	0x03D2: {"Upsilon1", "GREEK UPSILON WITH HOOK SYMBOL"},
	0x03AB: {"Upsilondieresis", "GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA"},
	0x038E: {"Upsilontonos", "GREEK CAPITAL LETTER UPSILON WITH TONOS"},
	0x016E: {"Uring", "LATIN CAPITAL LETTER U WITH RING ABOVE"},
	0x0168: {"Utilde", "LATIN CAPITAL LETTER U WITH TILDE"},
	0x0056: {"V", "LATIN CAPITAL LETTER V"},
	0x0057: {"W", "LATIN CAPITAL LETTER W"},
	0x1E82: {"Wacute", "LATIN CAPITAL LETTER W WITH ACUTE"},
	0x0174: {"Wcircumflex", "LATIN CAPITAL LETTER W WITH CIRCUMFLEX"},
	0x1E84: {"Wdieresis", "LATIN CAPITAL LETTER W WITH DIAERESIS"},
	0x1E80: {"Wgrave", "LATIN CAPITAL LETTER W WITH GRAVE"},
	0x0058: {"X", "LATIN CAPITAL LETTER X"},
	0x039E: {"Xi", "GREEK CAPITAL LETTER XI"},
	0x0059: {"Y", "LATIN CAPITAL LETTER Y"},
	0x00DD: {"Yacute", "LATIN CAPITAL LETTER Y WITH ACUTE"},
	0x0176: {"Ycircumflex", "LATIN CAPITAL LETTER Y WITH CIRCUMFLEX"},
	0x0178: {"Ydieresis", "LATIN CAPITAL LETTER Y WITH DIAERESIS"},
	0x1EF2: {"Ygrave", "LATIN CAPITAL LETTER Y WITH GRAVE"},
	0x005A: {"Z", "LATIN CAPITAL LETTER Z"},
	0x0179: {"Zacute", "LATIN CAPITAL LETTER Z WITH ACUTE"},
	0x017D: {"Zcaron", "LATIN CAPITAL LETTER Z WITH CARON"},
	0x017B: {"Zdotaccent", "LATIN CAPITAL LETTER Z WITH DOT ABOVE"},
	0x0396: {"Zeta", "GREEK CAPITAL LETTER ZETA"},
	0x0061: {"a", "LATIN SMALL LETTER A"},
	0x00E1: {"aacute", "LATIN SMALL LETTER A WITH ACUTE"},
	0x0103: {"abreve", "LATIN SMALL LETTER A WITH BREVE"},
	0x00E2: {"acircumflex", "LATIN SMALL LETTER A WITH CIRCUMFLEX"},
	0x00B4: {"acute", "ACUTE ACCENT"},
	0x0301: {"acutecomb", "COMBINING ACUTE ACCENT"},
	0x00E4: {"adieresis", "LATIN SMALL LETTER A WITH DIAERESIS"},
	0x00E6: {"ae", "LATIN SMALL LETTER AE"},
	0x01FD: {"aeacute", "LATIN SMALL LETTER AE WITH ACUTE"},
	0x00E0: {"agrave", "LATIN SMALL LETTER A WITH GRAVE"},
	0x2135: {"aleph", "ALEF SYMBOL"},
	0x03B1: {"alpha", "GREEK SMALL LETTER ALPHA"},
	0x03AC: {"alphatonos", "GREEK SMALL LETTER ALPHA WITH TONOS"},
	0x0101: {"amacron", "LATIN SMALL LETTER A WITH MACRON"},
	0x0026: {"ampersand", "AMPERSAND"},
	0x2220: {"angle", "ANGLE"},
	0x2329: {"angleleft", "LEFT-POINTING ANGLE BRACKET"},
	0x232A: {"angleright", "RIGHT-POINTING ANGLE BRACKET"},
	0x0387: {"anoteleia", "GREEK ANO TELEIA"},
	0x0105: {"aogonek", "LATIN SMALL LETTER A WITH OGONEK"},
	0x2248: {"approxequal", "ALMOST EQUAL TO"},
	0x00E5: {"aring", "LATIN SMALL LETTER A WITH RING ABOVE"},
	0x01FB: {"aringacute", "LATIN SMALL LETTER A WITH RING ABOVE AND ACUTE"},
	0x2194: {"arrowboth", "LEFT RIGHT ARROW"},
	0x21D4: {"arrowdblboth", "LEFT RIGHT DOUBLE ARROW"},
	0x21D3: {"arrowdbldown", "DOWNWARDS DOUBLE ARROW"},
	0x21D0: {"arrowdblleft", "LEFTWARDS DOUBLE ARROW"},
	0x21D2: {"arrowdblright", "RIGHTWARDS DOUBLE ARROW"},
	0x21D1: {"arrowdblup", "UPWARDS DOUBLE ARROW"},
	0x2193: {"arrowdown", "DOWNWARDS ARROW"},
	0x2190: {"arrowleft", "LEFTWARDS ARROW"},
	0x2192: {"arrowright", "RIGHTWARDS ARROW"},
	0x2191: {"arrowup", "UPWARDS ARROW"},
	0x2195: {"arrowupdn", "UP DOWN ARROW"},
	0x21A8: {"arrowupdnbse", "UP DOWN ARROW WITH BASE"},
	0x005E: {"asciicircum", "CIRCUMFLEX ACCENT"},
	0x007E: {"asciitilde", "TILDE"},
	0x002A: {"asterisk", "ASTERISK"},
	0x2217: {"asteriskmath", "ASTERISK OPERATOR"},
	0x0040: {"at", "COMMERCIAL AT"},
	0x00E3: {"atilde", "LATIN SMALL LETTER A WITH TILDE"},
	0x0062: {"b", "LATIN SMALL LETTER B"},
	0x005C: {"backslash", "REVERSE SOLIDUS"},
	0x007C: {"bar", "VERTICAL LINE"},
	0x03B2: {"beta", "GREEK SMALL LETTER BETA"},
	0x2588: {"block", "FULL BLOCK"},
	0x007B: {"braceleft", "LEFT CURLY BRACKET"},
	0x007D: {"braceright", "RIGHT CURLY BRACKET"},
	0x005B: {"bracketleft", "LEFT SQUARE BRACKET"},
	0x005D: {"bracketright", "RIGHT SQUARE BRACKET"},
	0x02D8: {"breve", "BREVE"},
	0x00A6: {"brokenbar", "BROKEN BAR"},
	0x2022: {"bullet", "BULLET"},
	0x0063: {"c", "LATIN SMALL LETTER C"},
	0x0107: {"cacute", "LATIN SMALL LETTER C WITH ACUTE"},
	0x02C7: {"caron", "CARON"},
	0x21B5: {"carriagereturn", "DOWNWARDS ARROW WITH CORNER LEFTWARDS"},
	0x010D: {"ccaron", "LATIN SMALL LETTER C WITH CARON"},
	0x00E7: {"ccedilla", "LATIN SMALL LETTER C WITH CEDILLA"},
	0x0109: {"ccircumflex", "LATIN SMALL LETTER C WITH CIRCUMFLEX"},
	0x010B: {"cdotaccent", "LATIN SMALL LETTER C WITH DOT ABOVE"},
	0x00B8: {"cedilla", "CEDILLA"},
	0x00A2: {"cent", "CENT SIGN"},
	0x03C7: {"chi", "GREEK SMALL LETTER CHI"},
	0x25CB: {"circle", "WHITE CIRCLE"},
	0x2297: {"circlemultiply", "CIRCLED TIMES"},
	0x2295: {"circleplus", "CIRCLED PLUS"},
	0x02C6: {"circumflex", "MODIFIER LETTER CIRCUMFLEX ACCENT"},
	0x2663: {"club", "BLACK CLUB SUIT"},
	0x003A: {"colon", "COLON"},
	0x20A1: {"colonmonetary", "COLON SIGN"},
	0x002C: {"comma", "COMMA"},
	0x2245: {"congruent", "APPROXIMATELY EQUAL TO"},
	0x00A9: {"copyright", "COPYRIGHT SIGN"},
	0x00A4: {"currency", "CURRENCY SIGN"},
	0x0064: {"d", "LATIN SMALL LETTER D"},
	0x2020: {"dagger", "DAGGER"},
	0x2021: {"daggerdbl", "DOUBLE DAGGER"},
	0x010F: {"dcaron", "LATIN SMALL LETTER D WITH CARON"},
	0x0111: {"dcroat", "LATIN SMALL LETTER D WITH STROKE"},
	0x00B0: {"degree", "DEGREE SIGN"},
	0x03B4: {"delta", "GREEK SMALL LETTER DELTA"},
	0x2666: {"diamond", "BLACK DIAMOND SUIT"},
	0x00A8: {"dieresis", "DIAERESIS"},
	0x0385: {"dieresistonos", "GREEK DIALYTIKA TONOS"},
	0x00F7: {"divide", "DIVISION SIGN"},
	0x2593: {"dkshade", "DARK SHADE"},
	0x2584: {"dnblock", "LOWER HALF BLOCK"},
	0x0024: {"dollar", "DOLLAR SIGN"},
	0x20AB: {"dong", "DONG SIGN"},
	0x02D9: {"dotaccent", "DOT ABOVE"},
	0x0323: {"dotbelowcomb", "COMBINING DOT BELOW"},
	0x0131: {"dotlessi", "LATIN SMALL LETTER DOTLESS I"},
	0x22C5: {"dotmath", "DOT OPERATOR"},
	0x0065: {"e", "LATIN SMALL LETTER E"},
	0x00E9: {"eacute", "LATIN SMALL LETTER E WITH ACUTE"},
	0x0115: {"ebreve", "LATIN SMALL LETTER E WITH BREVE"},
	0x011B: {"ecaron", "LATIN SMALL LETTER E WITH CARON"},
	0x00EA: {"ecircumflex", "LATIN SMALL LETTER E WITH CIRCUMFLEX"},
	0x00EB: {"edieresis", "LATIN SMALL LETTER E WITH DIAERESIS"},
	0x0117: {"edotaccent", "LATIN SMALL LETTER E WITH DOT ABOVE"},
	0x00E8: {"egrave", "LATIN SMALL LETTER E WITH GRAVE"},
	0x0038: {"eight", "DIGIT EIGHT"},
	0x2208: {"element", "ELEMENT OF"},
	0x2026: {"ellipsis", "HORIZONTAL ELLIPSIS"},
	0x0113: {"emacron", "LATIN SMALL LETTER E WITH MACRON"},
	0x2014: {"emdash", "EM DASH"},
	0x2205: {"emptyset", "EMPTY SET"},
	0x2013: {"endash", "EN DASH"},
	0x014B: {"eng", "LATIN SMALL LETTER ENG"},
	0x0119: {"eogonek", "LATIN SMALL LETTER E WITH OGONEK"},
	0x03B5: {"epsilon", "GREEK SMALL LETTER EPSILON"},
	0x03AD: {"epsilontonos", "GREEK SMALL LETTER EPSILON WITH TONOS"},
	0x003D: {"equal", "EQUALS SIGN"},
	0x2261: {"equivalence", "IDENTICAL TO"},
	0x212E: {"estimated", "ESTIMATED SYMBOL"},
	0x03B7: {"eta", "GREEK SMALL LETTER ETA"},
	0x03AE: {"etatonos", "GREEK SMALL LETTER ETA WITH TONOS"},
	0x00F0: {"eth", "LATIN SMALL LETTER ETH"},
	0x0021: {"exclam", "EXCLAMATION MARK"},
	0x203C: {"exclamdbl", "DOUBLE EXCLAMATION MARK"},
	0x00A1: {"exclamdown", "INVERTED EXCLAMATION MARK"},
	0x2203: {"existential", "THERE EXISTS"},
	0x0066: {"f", "LATIN SMALL LETTER F"},
	0x2640: {"female", "FEMALE SIGN"},
	0x2012: {"figuredash", "FIGURE DASH"},
	0x25A0: {"filledbox", "BLACK SQUARE"},
	0x25AC: {"filledrect", "BLACK RECTANGLE"},
	0x0035: {"five", "DIGIT FIVE"},
	0x215D: {"fiveeighths", "VULGAR FRACTION FIVE EIGHTHS"},
	0x0192: {"florin", "LATIN SMALL LETTER F WITH HOOK"},
	0x0034: {"four", "DIGIT FOUR"},
	0x2044: {"fraction", "FRACTION SLASH"},
	0x20A3: {"franc", "FRENCH FRANC SIGN"},
	0x0067: {"g", "LATIN SMALL LETTER G"},
	0x03B3: {"gamma", "GREEK SMALL LETTER GAMMA"},
	0x011F: {"gbreve", "LATIN SMALL LETTER G WITH BREVE"},
	0x01E7: {"gcaron", "LATIN SMALL LETTER G WITH CARON"},
	0x011D: {"gcircumflex", "LATIN SMALL LETTER G WITH CIRCUMFLEX"},
	0x0121: {"gdotaccent", "LATIN SMALL LETTER G WITH DOT ABOVE"},
	0x00DF: {"germandbls", "LATIN SMALL LETTER SHARP S"},
	0x2207: {"gradient", "NABLA"},
	0x0060: {"grave", "GRAVE ACCENT"},
	0x0300: {"gravecomb", "COMBINING GRAVE ACCENT"},
	0x003E: {"greater", "GREATER-THAN SIGN"},
	0x2265: {"greaterequal", "GREATER-THAN OR EQUAL TO"},
	0x00AB: {"guillemotleft", "LEFT-POINTING DOUBLE ANGLE QUOTATION MARK"},
	0x00BB: {"guillemotright", "RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK"},
	0x2039: {"guilsinglleft", "SINGLE LEFT-POINTING ANGLE QUOTATION MARK"},
	0x203A: {"guilsinglright", "SINGLE RIGHT-POINTING ANGLE QUOTATION MARK"},
	0x0068: {"h", "LATIN SMALL LETTER H"},
	0x0127: {"hbar", "LATIN SMALL LETTER H WITH STROKE"},
	0x0125: {"hcircumflex", "LATIN SMALL LETTER H WITH CIRCUMFLEX"},
	0x2665: {"heart", "BLACK HEART SUIT"},
	0x0309: {"hookabovecomb", "COMBINING HOOK ABOVE"},
	0x2302: {"house", "HOUSE"},
	0x02DD: {"hungarumlaut", "DOUBLE ACUTE ACCENT"},
	0x002D: {"hyphen", "HYPHEN-MINUS"},
	0x0069: {"i", "LATIN SMALL LETTER I"},
	0x00ED: {"iacute", "LATIN SMALL LETTER I WITH ACUTE"},
	0x012D: {"ibreve", "LATIN SMALL LETTER I WITH BREVE"},
	0x00EE: {"icircumflex", "LATIN SMALL LETTER I WITH CIRCUMFLEX"},
	0x00EF: {"idieresis", "LATIN SMALL LETTER I WITH DIAERESIS"},
	0x00EC: {"igrave", "LATIN SMALL LETTER I WITH GRAVE"},
	0x0133: {"ij", "LATIN SMALL LIGATURE IJ"},
	0x012B: {"imacron", "LATIN SMALL LETTER I WITH MACRON"},
	0x221E: {"infinity", "INFINITY"},
	0x222B: {"integral", "INTEGRAL"},
	0x2321: {"integralbt", "BOTTOM HALF INTEGRAL"},
	0x2320: {"integraltp", "TOP HALF INTEGRAL"},
	0x2229: {"intersection", "INTERSECTION"},
	0x25D8: {"invbullet", "INVERSE BULLET"},
	0x25D9: {"invcircle", "INVERSE WHITE CIRCLE"},
	0x263B: {"invsmileface", "BLACK SMILING FACE"},
	0x012F: {"iogonek", "LATIN SMALL LETTER I WITH OGONEK"},
	0x03B9: {"iota", "GREEK SMALL LETTER IOTA"},
	0x03CA: {"iotadieresis", "GREEK SMALL LETTER IOTA WITH DIALYTIKA"},
	0x0390: {"iotadieresistonos", "GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS"},
	0x03AF: {"iotatonos", "GREEK SMALL LETTER IOTA WITH TONOS"},
	0x0129: {"itilde", "LATIN SMALL LETTER I WITH TILDE"},
	0x006A: {"j", "LATIN SMALL LETTER J"},
	0x0135: {"jcircumflex", "LATIN SMALL LETTER J WITH CIRCUMFLEX"},
	0x006B: {"k", "LATIN SMALL LETTER K"},
	0x03BA: {"kappa", "GREEK SMALL LETTER KAPPA"},
	0x0138: {"kgreenlandic", "LATIN SMALL LETTER KRA"},
	0x006C: {"l", "LATIN SMALL LETTER L"},
	0x013A: {"lacute", "LATIN SMALL LETTER L WITH ACUTE"},
	0x03BB: {"lambda", "GREEK SMALL LETTER LAMDA"},
	0x013E: {"lcaron", "LATIN SMALL LETTER L WITH CARON"},
	0x0140: {"ldot", "LATIN SMALL LETTER L WITH MIDDLE DOT"},
	0x003C: {"less", "LESS-THAN SIGN"},
	0x2264: {"lessequal", "LESS-THAN OR EQUAL TO"},
	0x258C: {"lfblock", "LEFT HALF BLOCK"},
	0x20A4: {"lira", "LIRA SIGN"},
	0x2227: {"logicaland", "LOGICAL AND"},
	0x00AC: {"logicalnot", "NOT SIGN"},
	0x2228: {"logicalor", "LOGICAL OR"},
	0x017F: {"longs", "LATIN SMALL LETTER LONG S"},
	0x25CA: {"lozenge", "LOZENGE"},
	0x0142: {"lslash", "LATIN SMALL LETTER L WITH STROKE"},
	0x2591: {"ltshade", "LIGHT SHADE"},
	0x006D: {"m", "LATIN SMALL LETTER M"},
	0x00AF: {"macron", "MACRON"},
	0x2642: {"male", "MALE SIGN"},
	0x2212: {"minus", "MINUS SIGN"},
	0x2032: {"minute", "PRIME"},
	0x00B5: {"mu", "MICRO SIGN"},
	0x00D7: {"multiply", "MULTIPLICATION SIGN"},
	0x266A: {"musicalnote", "EIGHTH NOTE"},
	0x266B: {"musicalnotedbl", "BEAMED EIGHTH NOTES"},
	0x006E: {"n", "LATIN SMALL LETTER N"},
	0x0144: {"nacute", "LATIN SMALL LETTER N WITH ACUTE"},
	0x0149: {"napostrophe", "LATIN SMALL LETTER N PRECEDED BY APOSTROPHE"},
	0x0148: {"ncaron", "LATIN SMALL LETTER N WITH CARON"},
	0x0039: {"nine", "DIGIT NINE"},
	0x2209: {"notelement", "NOT AN ELEMENT OF"},
	0x2260: {"notequal", "NOT EQUAL TO"},
	0x2284: {"notsubset", "NOT A SUBSET OF"},
	0x00F1: {"ntilde", "LATIN SMALL LETTER N WITH TILDE"},
	0x03BD: {"nu", "GREEK SMALL LETTER NU"},
	0x0023: {"numbersign", "NUMBER SIGN"},
	0x006F: {"o", "LATIN SMALL LETTER O"},
	0x00F3: {"oacute", "LATIN SMALL LETTER O WITH ACUTE"},
	0x014F: {"obreve", "LATIN SMALL LETTER O WITH BREVE"},
	0x00F4: {"ocircumflex", "LATIN SMALL LETTER O WITH CIRCUMFLEX"},
	0x00F6: {"odieresis", "LATIN SMALL LETTER O WITH DIAERESIS"},
	0x0153: {"oe", "LATIN SMALL LIGATURE OE"},
	0x02DB: {"ogonek", "OGONEK"},
	0x00F2: {"ograve", "LATIN SMALL LETTER O WITH GRAVE"},
	0x01A1: {"ohorn", "LATIN SMALL LETTER O WITH HORN"},
	0x0151: {"ohungarumlaut", "LATIN SMALL LETTER O WITH DOUBLE ACUTE"},
	0x014D: {"omacron", "LATIN SMALL LETTER O WITH MACRON"},
	0x03C9: {"omega", "GREEK SMALL LETTER OMEGA"},
	0x03D6: {"omega1", "GREEK PI SYMBOL"},
	0x03CE: {"omegatonos", "GREEK SMALL LETTER OMEGA WITH TONOS"},
	0x03BF: {"omicron", "GREEK SMALL LETTER OMICRON"},
	0x03CC: {"omicrontonos", "GREEK SMALL LETTER OMICRON WITH TONOS"},
	0x0031: {"one", "DIGIT ONE"},
	0x2024: {"onedotenleader", "ONE DOT LEADER"},
	0x215B: {"oneeighth", "VULGAR FRACTION ONE EIGHTH"},
	0x00BD: {"onehalf", "VULGAR FRACTION ONE HALF"},
	0x00BC: {"onequarter", "VULGAR FRACTION ONE QUARTER"},
	0x2153: {"onethird", "VULGAR FRACTION ONE THIRD"},
	0x25E6: {"openbullet", "WHITE BULLET"},
	0x00AA: {"ordfeminine", "FEMININE ORDINAL INDICATOR"},
	0x00BA: {"ordmasculine", "MASCULINE ORDINAL INDICATOR"},
	0x221F: {"orthogonal", "RIGHT ANGLE"},
	0x00F8: {"oslash", "LATIN SMALL LETTER O WITH STROKE"},
	0x01FF: {"oslashacute", "LATIN SMALL LETTER O WITH STROKE AND ACUTE"},
	0x00F5: {"otilde", "LATIN SMALL LETTER O WITH TILDE"},
	0x0070: {"p", "LATIN SMALL LETTER P"},
	0x00B6: {"paragraph", "PILCROW SIGN"},
	0x0028: {"parenleft", "LEFT PARENTHESIS"},
	0x0029: {"parenright", "RIGHT PARENTHESIS"},
	0x2202: {"partialdiff", "PARTIAL DIFFERENTIAL"},
	0x0025: {"percent", "PERCENT SIGN"},
	0x002E: {"period", "FULL STOP"},
	0x00B7: {"periodcentered", "MIDDLE DOT"},
	0x22A5: {"perpendicular", "UP TACK"},
	0x2030: {"perthousand", "PER MILLE SIGN"},
	0x20A7: {"peseta", "PESETA SIGN"},
	0x03C6: {"phi", "GREEK SMALL LETTER PHI"},
	0x03D5: {"phi1", "GREEK PHI SYMBOL"},
	0x03C0: {"pi", "GREEK SMALL LETTER PI"},
	0x002B: {"plus", "PLUS SIGN"},
	0x00B1: {"plusminus", "PLUS-MINUS SIGN"},
	0x211E: {"prescription", "PRESCRIPTION TAKE"},
	0x220F: {"product", "N-ARY PRODUCT"},
	0x2282: {"propersubset", "SUBSET OF"},
	0x2283: {"propersuperset", "SUPERSET OF"},
	0x221D: {"proportional", "PROPORTIONAL TO"},
	0x03C8: {"psi", "GREEK SMALL LETTER PSI"},
	0x0071: {"q", "LATIN SMALL LETTER Q"},
	0x003F: {"question", "QUESTION MARK"},
	0x00BF: {"questiondown", "INVERTED QUESTION MARK"},
	0x0022: {"quotedbl", "QUOTATION MARK"},
	0x201E: {"quotedblbase", "DOUBLE LOW-9 QUOTATION MARK"},
	0x201C: {"quotedblleft", "LEFT DOUBLE QUOTATION MARK"},
	0x201D: {"quotedblright", "RIGHT DOUBLE QUOTATION MARK"},
	0x2018: {"quoteleft", "LEFT SINGLE QUOTATION MARK"},
	0x201B: {"quotereversed", "SINGLE HIGH-REVERSED-9 QUOTATION MARK"},
	0x2019: {"quoteright", "RIGHT SINGLE QUOTATION MARK"},
	0x201A: {"quotesinglbase", "SINGLE LOW-9 QUOTATION MARK"},
	0x0027: {"quotesingle", "APOSTROPHE"},
	0x0072: {"r", "LATIN SMALL LETTER R"},
	0x0155: {"racute", "LATIN SMALL LETTER R WITH ACUTE"},
	0x221A: {"radical", "SQUARE ROOT"},
	0x0159: {"rcaron", "LATIN SMALL LETTER R WITH CARON"},
	0x2286: {"reflexsubset", "SUBSET OF OR EQUAL TO"},
	0x2287: {"reflexsuperset", "SUPERSET OF OR EQUAL TO"},
	0x00AE: {"registered", "REGISTERED SIGN"},
	0x2310: {"revlogicalnot", "REVERSED NOT SIGN"},
	0x03C1: {"rho", "GREEK SMALL LETTER RHO"},
	0x02DA: {"ring", "RING ABOVE"},
	0x2590: {"rtblock", "RIGHT HALF BLOCK"},
	0x0073: {"s", "LATIN SMALL LETTER S"},
	0x015B: {"sacute", "LATIN SMALL LETTER S WITH ACUTE"},
	0x0161: {"scaron", "LATIN SMALL LETTER S WITH CARON"},
	0x015F: {"scedilla", "LATIN SMALL LETTER S WITH CEDILLA"},
	0x015D: {"scircumflex", "LATIN SMALL LETTER S WITH CIRCUMFLEX"},
	0x2033: {"second", "DOUBLE PRIME"},
	0x00A7: {"section", "SECTION SIGN"},
	0x003B: {"semicolon", "SEMICOLON"},
	0x0037: {"seven", "DIGIT SEVEN"},
	0x215E: {"seveneighths", "VULGAR FRACTION SEVEN EIGHTHS"},
	0x2592: {"shade", "MEDIUM SHADE"},
	0x03C3: {"sigma", "GREEK SMALL LETTER SIGMA"},
	0x03C2: {"sigma1", "GREEK SMALL LETTER FINAL SIGMA"},
	0x223C: {"similar", "TILDE OPERATOR"},
	0x0036: {"six", "DIGIT SIX"},
	0x002F: {"slash", "SOLIDUS"},
	0x263A: {"smileface", "WHITE SMILING FACE"},
	0x0020: {"space", "SPACE"},
	0x2660: {"spade", "BLACK SPADE SUIT"},
	0x00A3: {"sterling", "POUND SIGN"},
	0x220B: {"suchthat", "CONTAINS AS MEMBER"},
	0x2211: {"summation", "N-ARY SUMMATION"},
	0x263C: {"sun", "WHITE SUN WITH RAYS"},
	0x0074: {"t", "LATIN SMALL LETTER T"},
	0x03C4: {"tau", "GREEK SMALL LETTER TAU"},
	0x0167: {"tbar", "LATIN SMALL LETTER T WITH STROKE"},
	0x0165: {"tcaron", "LATIN SMALL LETTER T WITH CARON"},
	0x2234: {"therefore", "THEREFORE"},
	0x03B8: {"theta", "GREEK SMALL LETTER THETA"},
	0x03D1: {"theta1", "GREEK THETA SYMBOL"},
	0x00FE: {"thorn", "LATIN SMALL LETTER THORN"},
	0x0033: {"three", "DIGIT THREE"},
	0x215C: {"threeeighths", "VULGAR FRACTION THREE EIGHTHS"},
	0x00BE: {"threequarters", "VULGAR FRACTION THREE QUARTERS"},
	0x02DC: {"tilde", "SMALL TILDE"},
	0x0303: {"tildecomb", "COMBINING TILDE"},
	0x0384: {"tonos", "GREEK TONOS"},
	0x2122: {"trademark", "TRADE MARK SIGN"},
	0x25BC: {"triagdn", "BLACK DOWN-POINTING TRIANGLE"},
	0x25C4: {"triaglf", "BLACK LEFT-POINTING POINTER"},
	0x25BA: {"triagrt", "BLACK RIGHT-POINTING POINTER"},
	0x25B2: {"triagup", "BLACK UP-POINTING TRIANGLE"},
	0x0032: {"two", "DIGIT TWO"},
	0x2025: {"twodotenleader", "TWO DOT LEADER"},
	0x2154: {"twothirds", "VULGAR FRACTION TWO THIRDS"},
	0x0075: {"u", "LATIN SMALL LETTER U"},
	0x00FA: {"uacute", "LATIN SMALL LETTER U WITH ACUTE"},
	0x016D: {"ubreve", "LATIN SMALL LETTER U WITH BREVE"},
	0x00FB: {"ucircumflex", "LATIN SMALL LETTER U WITH CIRCUMFLEX"},
	0x00FC: {"udieresis", "LATIN SMALL LETTER U WITH DIAERESIS"},
	0x00F9: {"ugrave", "LATIN SMALL LETTER U WITH GRAVE"},
	0x01B0: {"uhorn", "LATIN SMALL LETTER U WITH HORN"},
	0x0171: {"uhungarumlaut", "LATIN SMALL LETTER U WITH DOUBLE ACUTE"},
	0x016B: {"umacron", "LATIN SMALL LETTER U WITH MACRON"},
	0x005F: {"underscore", "LOW LINE"},
	0x2017: {"underscoredbl", "DOUBLE LOW LINE"},
	0x222A: {"union", "UNION"},
	0x2200: {"universal", "FOR ALL"},
	0x0173: {"uogonek", "LATIN SMALL LETTER U WITH OGONEK"},
	0x2580: {"upblock", "UPPER HALF BLOCK"},
	0x03C5: {"upsilon", "GREEK SMALL LETTER UPSILON"},
	0x03CB: {"upsilondieresis", "GREEK SMALL LETTER UPSILON WITH DIALYTIKA"},
	0x03B0: {"upsilondieresistonos", "GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS"},
	0x03CD: {"upsilontonos", "GREEK SMALL LETTER UPSILON WITH TONOS"},
	0x016F: {"uring", "LATIN SMALL LETTER U WITH RING ABOVE"},
	0x0169: {"utilde", "LATIN SMALL LETTER U WITH TILDE"},
	0x0076: {"v", "LATIN SMALL LETTER V"},
	0x0077: {"w", "LATIN SMALL LETTER W"},
	0x1E83: {"wacute", "LATIN SMALL LETTER W WITH ACUTE"},
	0x0175: {"wcircumflex", "LATIN SMALL LETTER W WITH CIRCUMFLEX"},
	0x1E85: {"wdieresis", "LATIN SMALL LETTER W WITH DIAERESIS"},
	0x2118: {"weierstrass", "SCRIPT CAPITAL P"},
	0x1E81: {"wgrave", "LATIN SMALL LETTER W WITH GRAVE"},
	0x0078: {"x", "LATIN SMALL LETTER X"},
	0x03BE: {"xi", "GREEK SMALL LETTER XI"},
	0x0079: {"y", "LATIN SMALL LETTER Y"},
	0x00FD: {"yacute", "LATIN SMALL LETTER Y WITH ACUTE"},
	0x0177: {"ycircumflex", "LATIN SMALL LETTER Y WITH CIRCUMFLEX"},
	0x00FF: {"ydieresis", "LATIN SMALL LETTER Y WITH DIAERESIS"},
	0x00A5: {"yen", "YEN SIGN"},
	0x1EF3: {"ygrave", "LATIN SMALL LETTER Y WITH GRAVE"},
	0x007A: {"z", "LATIN SMALL LETTER Z"},
	0x017A: {"zacute", "LATIN SMALL LETTER Z WITH ACUTE"},
	0x017E: {"zcaron", "LATIN SMALL LETTER Z WITH CARON"},
	0x017C: {"zdotaccent", "LATIN SMALL LETTER Z WITH DOT ABOVE"},
	0x0030: {"zero", "DIGIT ZERO"},
	0x03B6: {"zeta", "GREEK SMALL LETTER ZETA"},
}
