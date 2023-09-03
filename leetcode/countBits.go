func countBits(n int) []int {
    result := make([]int, n+1)

    for i := 1; i <= n; i++ {
        count := 0
        num := i
        for num > 0 {
            num = num & (num - 1)
            count++
        }
        result[i] = count
    }

    return result
}
