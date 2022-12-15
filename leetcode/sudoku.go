func solveSudoku(board [][]byte)  {
    backTrack(board, 0)
}

func backTrack(tracker [][]byte, count int) bool {
    
    if count == 81 {
        return true
    }
    x := count/9
    y := count%9
    if tracker[x][y] != '.' {
        return backTrack(tracker, count+1)
    }
    
    var i byte
    for i=49;i<=57;i++{
        if numNotPresentInRow(tracker[x], i) && numNotPresentInColumn(tracker, y, i) &&
        numNotPresentInSquare(tracker, x, y, i) && tracker[x][y] == '.' {
            
            tracker[x][y] = i
            if backTrack(tracker, count+1) {
                return true
            }
            tracker[x][y] = '.'
        }
    }
    return false
}

//These are helpers to check if the number is present in row, column or 3X3 square region
func numNotPresentInRow(trackerRow []byte, i byte) bool {
    for _, val := range trackerRow {
        if i == val {
            return false
        }
    }
    return true
}

func numNotPresentInColumn(tracker [][]byte, y int, i byte) bool {
    for _, row := range tracker {
        if row[y] == i {
            return false
        }
    }
    return true
}

func numNotPresentInSquare(tracker [][]byte, x, y int, i byte) bool {
    a1 := 0
    a2 := 0
    
    b1 := 0
    b2 := 0
    if x >= 0 && x < 3 {
        a1 = 0
        a2 = 3
    }else if x >= 3 && x < 6 {
        a1 = 3
        a2 = 6
    } else {
        a1 = 6
        a2 = 9
    }
    
    if y >= 0 && y < 3 {
        b1 = 0
        b2 = 3
    }else if y >= 3 && y < 6 {
        b1 = 3
        b2 = 6
    } else {
        b1 = 6
        b2 = 9
    }
    for m := a1;m<a2;m++{
        for n := b1;n<b2;n++{
            if tracker[m][n] == i {
                return false
            }
        }
    }
    return true
}
