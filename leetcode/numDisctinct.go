func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([][]int, lt+1)
	for k := range dp {
		dp[k] = make([]int, ls+1)
	}
	for k := range dp[0] {
		dp[0][k] = 1
	}
	for i := 1; i <= lt; i++ {
		for j := 1; j <= ls; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[lt][ls]
}
