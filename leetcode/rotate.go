func rotate(matrix [][]int) {
	l := len(matrix) / 2
	x := math.Floor(float64(l))
	for min := 0; min < int(x); min++ {
		max := len(matrix) - min - 1
		for i := min; i < max; i++ {
			offset := i - min
			top := matrix[i][min]
			top = matrix[min][i]
			matrix[min][i] = matrix[max - offset][min]
			matrix[max - offset][min] = matrix[max][max - offset]
			matrix[max][max - offset] = matrix[i][max]
			matrix[i][max] = top
		}
	}
}
