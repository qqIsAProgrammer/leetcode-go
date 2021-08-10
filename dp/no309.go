package dp

import "math"

func maxProfit309(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	f := make([][3]int, n)
	f[0][0] = -prices[0]
	f[0][1] = math.MinInt64
	f[0][2] = 0
	for i := 1; i < n; i++ {
		f[i][0] = maxInt(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = maxInt(f[i-1][1], f[i-1][2])
	}
	return maxInt(f[n-1][1], f[n-1][2])
}
