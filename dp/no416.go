package dp

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}

	tar := sum / 2
	if tar < max {
		return false
	}

	// dp[i][j]: [0, i] 区间取若干个数之和是否存在等于 j.
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, tar+1)
	}
	dp[0][nums[0]] = true
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	for i := 1; i < n; i++ {
		for j := 1; j < tar+1; j++ {
			if j >= nums[i] { // 当 j >= nums[i] 可以选取 nums[i] 也可以不选取.
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else { // 当 j < nums[i] 不选取 nums[i].
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][tar]
}
