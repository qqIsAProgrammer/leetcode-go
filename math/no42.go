package math

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	l := make([]int, n)
	r := make([]int, n)
	l[0] = height[0]
	r[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		l[i] = max(l[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		r[i] = max(r[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += min(l[i], r[i]) - height[i]
	}
	return ans
}
