func numEnclaves(grid [][]int) int {
    var count int

    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
            return
        }

        grid[i][j] = -1 

        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    for i := 0; i < len(grid); i++ {
        dfs(i, 0)        
        dfs(i, len(grid[0])-1)
    }
    for j := 0; j < len(grid[0]); j++ {
        dfs(0, j)          
        dfs(len(grid)-1, j) 
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                count++
            }
        }
    }

    return count
}
