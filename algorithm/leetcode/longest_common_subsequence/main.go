package main

func main() {
	println(longestCommonSubSequence("abcde", "ace"))
}

func longestCommonSubSequence(s1, s2 string) int {
	// 初始化a[i][j]数组
	a := make([][]int, len(s1) + 1)
	for i := 0; i < len(s1) + 1; i++ {
		a[i] = make([]int, len(s2)  + 1)
	}

	max := 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if []byte(s1)[i-1] == []byte(s2)[j-1] {
				a[i][j] = a[i-1][j-1] + 1
			} else {
				if a[i-1][j] > a[i][j-1] {
					a[i][j] = a[i-1][j]
				} else {
					a[i][j] = a[i][j-1]
				}
			}
			if a[i][j] > max {
				max = a[i][j]
			}
		}
	}
	return max
}
