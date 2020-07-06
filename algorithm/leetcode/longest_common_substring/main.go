package main

func main() {
	println(longestCommonSubstring("acdfg", "akdfc"))
}

func longestCommonSubstring(s1, s2 string) string {
	// 初始化a[i][j]数组
	a := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		a[i] = make([]int, len(s2))
	}

	// 如果s1的首字母和s2中[0:j]的某一个相同，那么a[0][0:j]相应位置就是1
	for i := 0; i < len(s2); i++ {
		if []byte(s1)[0] == []byte(s2)[i] {
			a[0][i] = 1
		}
	}

	// 如果s2的首字母和s1中[0:i]的某一个相同，那么a[0:i][0]相应位置就是1
	for i := 0; i < len(s1); i++ {
		if []byte(s2)[0] == []byte(s1)[i] {
			a[i][0] = 1
		}
	}

	max := 0
	ii := 0
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if []byte(s1)[i] == []byte(s2)[j] {
				// 状态转移方程的表示
				a[i][j] = a[i-1][j-1] + 1
				if a[i][j] > max {
					max = a[i][j]
					ii = i
				}
			}
		}
	}
	return string([]byte(s1)[ii+1-max : ii+1])
}
