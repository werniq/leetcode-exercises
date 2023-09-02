func numTrees(n int) int {
    if n <= 1 {
        return 1
    }
    
    dp := make([]int, n+1)
    // default values for 0 and 1
    dp[0], dp[1] = 1, 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            // The number of unique BSTs with j as the root is the product of the
            // number of BSTs on the left subtree (j-1 nodes) and the number of BSTs
            // on the right subtree (i-j nodes).
            dp[i] += dp[j-1] * dp[i-j]
        }
    }
    
    return dp[n]
}

