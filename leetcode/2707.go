func minExtraChar(s string, dictionary []string) int {
    dictionarySet := make(map[string]bool)
    for _, word := range dictionary {
        dictionarySet[word] = true
    }

    n := len(s)
    dp := make([]int, n+1)

    dp[0] = 0

    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] + 1 

        for j := i - 1; j >= 0; j-- {
            if dictionarySet[s[j:i]] {
                dp[i] = min(dp[i], dp[j])
            }
        }
    }

    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
