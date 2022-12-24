/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type NodeHeap []*ListNode
func (h NodeHeap) Len() int {
    return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
    return h[i].Val <= h[j].Val
}
func (h NodeHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *NodeHeap) Pop() interface{} {
    old := *h
    value := old[len(old)-1]
    *h = old[:len(old)-1]
    return value
}
func (h *NodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
    // loop for find current smallest
    var head *ListNode
    var currentNode *ListNode
    q := &NodeHeap{}
    for _, node := range lists {
        if node != nil {
            heap.Push(q, node)
        }
    }
    for q.Len() > 0 {
        n := heap.Pop(q).(*ListNode)
        
        if n.Next != nil {
            heap.Push(q, n.Next)
        }
        
        if (head == nil) {
            head = n
            currentNode = n
        } else {
            currentNode.Next = n
            currentNode = currentNode.Next
        }
    }
    return head
}
