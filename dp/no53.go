package dp

func maxSubArray(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	ans := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(cur+nums[i], nums[i])
		ans = max(ans, cur)
	}
	return ans
}
