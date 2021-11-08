package sort

import "sync"

var wg sync.WaitGroup

func QuickSortParallel(nums []int) {
	quickSortParallel(nums, 0, len(nums)-1)
	wg.Wait()
}

func quickSortParallel(nums []int, lo, hi int) {
	if lo > hi {
		return
	}

	pivot := nums[lo]
	i, j := lo, hi
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[lo] = nums[lo], nums[i]

	wg.Add(2)
	go func() {
		quickSortParallel(nums, lo, i-1)
		wg.Done()
	}()
	go func() {
		quickSortParallel(nums, i+1, hi)
		wg.Done()
	}()
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, lo, hi int) {
	if lo > hi {
		return
	}

	pivot := nums[lo]
	i, j := lo, hi
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[lo] = nums[lo], nums[i]

	quickSort(nums, lo, i-1)
	quickSort(nums, i+1, hi)
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
