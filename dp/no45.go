package dp

func jump(nums []int) int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	pos := 0
	step := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		pos = maxInt(pos, i+nums[i])
		if i == end {
			end = pos
			step++
		}
	}
	return step
}
