package math

func maxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	ans := max
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := max
			max = maxInt(nums[i], nums[i]*min)
			min = minInt(nums[i], nums[i]*tmp)
		} else {
			max = maxInt(nums[i], nums[i]*max)
			min = minInt(nums[i], nums[i]*min)
		}
		ans = maxInt(ans, max)
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
