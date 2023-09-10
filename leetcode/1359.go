func countOrders(n int) int {
    const mod = 1000000007
    
    dp := make([]int, n+1)
    dp[0] = 1
    
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] * (2*i - 1) * (2*i) / 2 % mod
    }
    
    return dp[n]
}
