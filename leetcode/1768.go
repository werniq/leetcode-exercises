func mergeAlternately(word1 string, word2 string) string {
    return iterateThroughStrings(word1, word2)
}

//                         bigger      lower
func iterateThroughStrings(arr string, arr1 string) string {
    res := ""
    for i := 0; i <= len(arr1)-1 ; i++ {
        if i == len(arr1) - 1 && i < len(arr)-1 {
            res += string(arr[i])
            res += string(arr1[i])
            res += string(arr[i+1:])
            return res
        } else if i == len(arr)-1 && i < len(arr1)-1{
            res += string(arr[i])
            res += string(arr1[i:])
            return res
        }
        res += string(arr[i])
        res += string(arr1[i])
    }
    return res
}
