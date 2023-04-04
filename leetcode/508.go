func findFrequentTreeSum(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    freq := make(map[int]int)
    maxFreq := 0
    traverse(root, freq, &maxFreq)
    res := make([]int, 0)
    for k, v := range freq {
        if v == maxFreq {
            res = append(res, k)
        }
    }
    return res
}

func traverse(root *TreeNode, freq map[int]int, maxFreq *int) int {
    if root == nil {
        return 0
    }
    sum := traverse(root.Left, freq, maxFreq) + traverse(root.Right, freq, maxFreq) + root.Val
    freq[sum]++
    if freq[sum] > *maxFreq {
        *maxFreq = freq[sum]
    }
    return sum
}
