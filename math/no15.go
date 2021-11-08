package math

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for fir := 0; fir < n-2; fir++ {
		if fir > 0 && nums[fir] == nums[fir-1] {
			continue
		}
		target := -1 * nums[fir]
		thi := n - 1
		for sec := fir + 1; sec < n-1; sec++ {
			if sec > fir+1 && nums[sec] == nums[sec-1] {
				continue
			}
			for sec < thi && nums[sec]+nums[thi] > target {
				thi--
			}
			if sec == thi {
				break
			}
			if nums[sec]+nums[thi] == target {
				ans = append(ans, []int{nums[fir], nums[sec], nums[thi]})
			}
		}
	}
	return ans
}
