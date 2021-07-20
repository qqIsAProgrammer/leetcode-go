package linked_list

// AC
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fir := head

	dum := new(ListNode)
	dum.Next = head
	curr := dum
	for fir != nil && fir.Next != nil {
		sec := fir.Next
		curr.Next = sec
		fir.Next = sec.Next
		sec.Next = fir

		curr = fir
		fir = fir.Next
	}

	return dum.Next
}
