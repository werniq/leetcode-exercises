var avail = 0
var noAvail = 1

type Solution struct {
	n int
	placedQueens []int
	step         int
}

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}

	var finalAns = 0
	for i := 0; i < n; i++ {
		var placedQueens = make([]int, n)
		placedQueens[0] = i
		var solution = Solution{n, placedQueens, 1}
		var ans = recSolution(solution)
		finalAns = finalAns + len(ans)
	}

	return finalAns
}

func recSolution(solution Solution) [][]int {
	var step = solution.step + 1
	var rowNum = solution.step
	var availablePos = calculateAvailPos(solution.placedQueens, rowNum, solution.n)

	var finalAns = [][]int{}

	if step == solution.n {
		for i := 0; i < solution.n; i++ {
			if availablePos[i] == avail {
				solution.placedQueens[rowNum] = i
				var successAns = make([]int, solution.n)
				for j := 0; j < solution.n; j++ {
					successAns[j] = solution.placedQueens[j]
				}
				return append(finalAns, successAns)
			}
		}
		return finalAns
	} else {
		for i := 0; i < solution.n; i++ {
			if availablePos[i] == avail {
				solution.placedQueens[rowNum] = i
				solution.step = step
				var ans = recSolution(solution)
				for j := 0; j < len(ans); j++ {
					finalAns = append(finalAns, ans[j])
				}
			}
		}
		return finalAns
	}
}

// Given a board that has placed Queen, update board with attacked spots.
func calculateAvailPos(placedQueens []int, rowNum int, n int) []int {
	var row = make([]int, n)
	if rowNum == 0 {
		return row
	}

	for i := 0; i < rowNum; i++ {
		var prevQueen = placedQueens[i]
		row[prevQueen] = noAvail
		var diff = rowNum - i
		if prevQueen-diff >= 0 {
			row[prevQueen-diff] = noAvail
		}
		if prevQueen+diff <= n-1 {
			row[prevQueen+diff] = noAvail
		}
	}
	return row
}
