func findSubstring(s string, words []string) []int {
    lenOne := len(words[0])
    totLen := lenOne * len(words)
    if len(s) < totLen {
        return []int{}
    }
    wordsMap := make(map[string]int)
    for _, word := range words {
        if v, ok := wordsMap[word]; ok {
            wordsMap[word] = v + 1
        } else {
            wordsMap[word] = 1
        }
    }

    startIndex := 0 
    endIndex := startIndex + totLen - 1
    res := []int{}

    for endIndex < len(s) {
        copiedWordsMap := copyMap(wordsMap)
        i := 0
        ex := true 
        for i < len(words) {
            subStrStarted := startIndex + i * lenOne 
            subStrEnded := startIndex + i * lenOne + lenOne
            if v, ok := copiedWordsMap[s[subStrStarted:subStrEnded]]; !ok {
                ex = false
                break
            } else {
                if v == 1 {
                    delete(copiedWordsMap, s[subStrStarted:subStrEnded])
                } else if v > 1 {
                    copiedWordsMap[s[subStrStarted:subStrEnded]] = v - 1
                }
            }
            i++
        }
        if ex {
            res = append(res, startIndex)
        }
        startIndex++
        endIndex++
    }
    return res
}


func copyMap(m map[string]int) map[string]int {
    copiedMap := make(map[string]int)

    for k, v := range m {
        copiedMap[k] = v
    }
    return copiedMap
}
