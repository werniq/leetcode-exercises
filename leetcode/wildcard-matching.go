func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func initLookupTable(row, column int) [][]bool {
	lookup := make([][]bool, row)
	for i := range lookup {
		lookup[i] = make([]bool, column)
	}
	return lookup
}


func isMatch(str, pattern string) bool {
	s := StringToRuneSlice(str)
	p := StringToRuneSlice(pattern)

	if len(p) == 0 {
		return len(s) == 0
	}

	lookup := initLookupTable(len(s)+1, len(p)+1)

	lookup[0][0] = true

	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			lookup[0][j] = lookup[0][j-1]
		}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				lookup[i][j] = lookup[i][j-1] || lookup[i-1][j]

			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				lookup[i][j] = lookup[i-1][j-1]

			} else {
				lookup[i][j] = false
			}
		}
	}

	return lookup[len(s)][len(p)]
}
