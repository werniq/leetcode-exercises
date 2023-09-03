func uniquePaths(m int, n int) int {
    result := 1
    for i := 1; i <= m-1; i++ {
        result *= (n - 1 + i)
        result /= i
    }
    
    return result
}
