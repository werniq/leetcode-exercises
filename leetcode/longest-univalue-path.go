func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)
		l, r := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			l += left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			r += right + 1
		}
		res = max(res, l+r)
		return max(l, r)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
