func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	partSize := length / k
	extra := length % k

	result := make([]*ListNode, k)
	current = head
	for i := 0; i < k; i++ {
		result[i] = current

		size := partSize
		if i < extra {
			size++
		}

		for j := 0; j < size-1; j++ {
			current = current.Next
		}

		if current != nil {
			next := current.Next
			current.Next = nil
			current = next
		}
	}

	return result
}
