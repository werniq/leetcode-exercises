func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	current := prev.Next
	next := current.Next

	for i := 0; i < right-left; i++ {
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = current.Next
	}

	return dummy.Next
}

