func generate(numRows int) [][]int {
    if numRows <= 0 {
        return nil
    }

    result := [][]int{{1}}

    for i := 1; i < numRows; i++ {
        newRow := []int{1}

        for j := 1; j < i; j++ {
            newValue := result[i-1][j-1] + result[i-1][j]
            newRow = append(newRow, newValue)
        }

        newRow = append(newRow, 1)

        result = append(result, newRow)
    }

    return result
}
