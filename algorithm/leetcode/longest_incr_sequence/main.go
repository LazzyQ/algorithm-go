package main

func main() {

}

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		cnt[i] = 1
	}

	max := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // 递增
				if dp[i] < dp[j] + 1 { // 如果以nums[i]结尾的最长递增序列 < 以nums[i]结尾的最长递增序列 + 1
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i] ==  dp[j] + 1 {
					cnt[i] = cnt[j] + 1
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == max {
			sum += cnt[i]
		}
	}
	return sum
}
