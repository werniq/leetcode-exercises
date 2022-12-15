func letterCombinations(digits string) []string {
	
	if digits == ""{
        return []string{}
	}

	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
    r := []string{""}

	for i := 0; i < len(digits); i++ {
		retLen := len(r)
        ll := m[string(digits[i])]
		for j := 0; j < retLen; j++ {
			for _, l := range ll {
				r = append(r, r[0]+l)
			}
			r = r[1:]
		}

	}

	return r

}
