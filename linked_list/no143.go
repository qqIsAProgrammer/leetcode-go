package linked_list

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	l1 := head
	tmp := slow.Next
	slow.Next = nil
	l2 := reverse(tmp)

	for l2 != nil {
		next := l2.Next
		l2.Next = l1.Next
		l1.Next = l2
		l1 = l2.Next
		l2 = next
	}
}

func reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	var prev *ListNode
	for l != nil {
		next := l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}
