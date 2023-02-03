package goafm

func Fields(s string) []string {
	ret := make([]string, 0, 6)
	start := 0
	for i, v := range s {
		if isSpace(v) {
			if start < i {
				ret = append(ret, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		ret = append(ret, s[start:])
	}
	return ret
}

func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}
