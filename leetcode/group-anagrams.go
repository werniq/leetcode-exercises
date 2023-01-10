func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	for _, str := range strs {
		isMatched := false
		for i, _ := range result {
			group := &result[i]
			if matched(str, (*group)[0]) {
				(*group) = append((*group), str)
				isMatched = true
				continue
			}
		}

		if !isMatched {
			result = append(result, []string{str})
		}
	}

	return result
}

func matched(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, s := range a {
		map1[s] += 1
	}
	for _, s := range b {
		map2[s] += 1
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}

	return true
}
