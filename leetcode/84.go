func largestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }
    stack := make([]int, 0)
    maxArea := 0

    calculateArea := func(height, width int) {
        area := height * width
        if area > maxArea {
            maxArea = area
        }
    }

    for i := 0; i <= n; i++ {
        var h int
        if i == n {
            h = 0
        } else {
            h = heights[i]
        }

        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            barIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            var width int
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            } else {
                width = i
            }

            calculateArea(heights[barIndex], width)
        }

        stack = append(stack, i)
    }

    return maxArea
}
