package sort

func InsertSort(nums []int) {
	var i,j,v int
	for i = 1; i < len(nums); i++ {
		v = nums[i]
		for j = i - 1; j >= 0 && nums[j] > v; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = nums[i]
	}
}