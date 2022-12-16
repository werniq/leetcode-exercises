/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * 0 <-> 1 | 2 <-> 3 
 }
 */ 
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	head = curr.Next
	for ; curr != nil && curr.Next != nil; curr = curr.Next {
		next := curr.Next
		if prev != nil {
			prev.Next = next
		}
		curr.Next, next.Next, prev = next.Next, curr, curr
	}

	return head
}
