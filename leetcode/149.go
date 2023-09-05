// Max points at a line 
func maxPoints(points [][]int) int {
    if len(points) <= 2 {
        return len(points)
    }

    maxCount := 0

    for i := 0; i < len(points); i++ {
        slopeMap := make(map[float64]int)
        samePointCount := 0
        verticalCount := 0
        localMax := 0

        for j := 0; j < len(points); j++ {
            if i == j {
                samePointCount++
                continue
            }

            dx := points[j][0] - points[i][0]
            dy := points[j][1] - points[i][1]

            if dx == 0 {
                verticalCount++
                continue
            }

            slope := float64(dy) / float64(dx)
            slopeMap[slope]++
            localMax = max(localMax, slopeMap[slope])
        }

        localMax = max(localMax, verticalCount) + samePointCount + 1
        maxCount = max(maxCount, localMax)
    }

    return maxCount - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
