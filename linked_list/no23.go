package linked_list

import "container/heap"

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *Heap) Push(x interface{}) {
	node := x.(*ListNode)
	*h = append(*h, node)
}

func (h *Heap) Pop() interface{} {
	n := len(*h) - 1
	node := (*h)[n]
	*h = (*h)[:n]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	h := new(Heap)
	heap.Init(h)
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}
