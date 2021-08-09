package dp

import "math"

func maxProfit(prices []int) int {
	low := math.MaxInt64
	profit := 0
	for _, p := range prices {
		if p < low {
			low = p
		} else if p-low > profit {
			profit = p - low
		}
	}
	return profit
}
