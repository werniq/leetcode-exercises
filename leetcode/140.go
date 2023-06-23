func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make(map[int][]string)
	return wordBreakHelper(s, 0, wordSet, memo)
}

func wordBreakHelper(s string, start int, wordSet map[string]bool, memo map[int][]string) []string {
	if result, ok := memo[start]; ok {
		return result
	}
	result := make([]string, 0)
	if start == len(s) {
		result = append(result, "")
	}
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordSet[word] {
			nextResults := wordBreakHelper(s, end, wordSet, memo)
			for _, nextResult := range nextResults {
				if nextResult == "" {
					result = append(result, word)
				} else {
					result = append(result, word+" "+nextResult)
				}
			}
		}
	}
	memo[start] = result
	return result
}
