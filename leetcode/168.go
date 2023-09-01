import (
    "strings"
)

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    
    words := strings.Fields(s)
    
    reverse(words)
    
    result := strings.Join(words, " ")
    
    return result
}

func reverse(arr []string) {
    i, j := 0, len(arr)-1
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}

