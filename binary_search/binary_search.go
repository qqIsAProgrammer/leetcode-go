package binary_search

func BinarySearch(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
