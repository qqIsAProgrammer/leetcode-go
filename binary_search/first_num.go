package binary_search

// 给定一个非降序数组和一个目标值
// 找到数组中第一个大于等于目标值的元素对应的索引
func Search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			if mid > 0 && nums[mid-1] < target {
				return mid
			} else {
				hi = mid - 1
			}
		} else {
			if mid == 0 || (mid > 0 && nums[mid-1] < target) {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
