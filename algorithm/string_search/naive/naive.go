package naive

func StringSearch(s, p string) int {
	var i, j int
	sLen := len(s)
	pLen := len(p)

	for i < sLen && j < pLen {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == pLen {
		return i - j
	}
	return -1
}
