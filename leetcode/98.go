func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    return helper(root, &prev)
}

func helper(node *TreeNode, prev **TreeNode) bool {
    if node == nil {
        return true
    }
    
    if !helper(node.Left, prev) {
        return false
    }
    
    if *prev != nil && node.Val <= (*prev).Val {
        return false
    }
    
    *prev = node
    
    return helper(node.Right, prev)
}
