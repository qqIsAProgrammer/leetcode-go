package dp

func rob1(nums []int) int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	total := 0
	prev := 0
	for _, num := range nums {
		tmp := total
		total = maxInt(prev+num, total)
		prev = tmp
	}
	return total
}
