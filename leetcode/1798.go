func getMaximumConsecutive(coins []int) int {
    sort.Ints(coins)  
    res := 1  
    for i := 0; i < len(coins); i++ {
        if coins[i] > res { 
            break  
        }
        res += coins[i] 
    return res
}
