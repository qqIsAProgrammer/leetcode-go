package heap

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for key, val := range freq {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).([2]int)[0])
	}
	return ans
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
