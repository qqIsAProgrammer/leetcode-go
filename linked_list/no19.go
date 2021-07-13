package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	prev := dummy
	curr := dummy

	for i := 0; i < n; i++ {
		prev = prev.Next
	}

	for prev.Next != nil {
		prev = prev.Next
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return dummy.Next
}
