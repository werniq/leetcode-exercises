/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    
    if root == nil {
        return []int{}
    }
    queue := []*TreeNode{root}
    result := []int{}
    for len(queue) > 0 {
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            if i == levelSize-1 {
                result = append(result, currentNode.Val)
            }
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }
    return result
}
