package main

func convert(str string, row int) string {
	if row == 1 {
		return str
	}
	item_len := 2*row - 2
	res := make([][]string, row, row)

	for index, v := range str {

		mod := index % item_len

		if mod < row {
			res[mod] = append(res[mod], string(v))
		} else {
			i := row - (mod - row) - 2
			res[i] = append(res[i], string(v))
		}
	}

	var s string

	for _, arr := range res {
		for _, v := range arr {
			s += v
		}
	}

	return s

}
