import "math"
func divide(dividend int, divisor int) int {
    if dividend == -2147483648 && divisor == -1{
        return 2147483647
    }
    var sol int = dividend / divisor
    return int(math.Floor(float64(sol)))
}
