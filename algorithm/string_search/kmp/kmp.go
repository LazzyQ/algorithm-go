package kmp

func GetNext(p string) []int {
	pLen := len(p)
	i, j := 0, -1
	next := make([]int, pLen + 1)
	next[0] = -1
	for ; i < pLen; {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
