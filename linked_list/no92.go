package linked_list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	begin := dummy
	for i := 0; i < left-1; i++ {
		begin = begin.Next
	}

	rnode := begin
	for i := 0; i < right-left+1; i++ {
		rnode = rnode.Next
	}

	lnode := begin.Next
	curr := rnode.Next

	begin.Next = nil
	rnode.Next = nil
	reverseList(lnode)

	begin.Next = rnode
	lnode.Next = curr
	return dummy.Next
}

func reverseList(head *ListNode) {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
}
