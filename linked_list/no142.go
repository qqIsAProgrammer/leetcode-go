package linked_list

// AC
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	flag := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			flag = true
			break
		}
	}

	if flag {
		for slow != head {
			slow = slow.Next
			head = head.Next
		}
		return head
	}
	return nil
}
