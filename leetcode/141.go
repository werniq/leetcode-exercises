/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    tortoise := head
    hare := head

    for hare != nil && hare.Next != nil {
        tortoise = tortoise.Next
        hare = hare.Next.Next

        if tortoise == hare {
            return true 
        }
    }

    return false 
}
