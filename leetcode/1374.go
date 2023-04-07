func generateTheString(n int) string {
    var result string
    if n%2 == 0 {
        result = strings.Repeat("a", n-1) + "b"
    } else {
        result = strings.Repeat("a", n)
    }
    return result
}
