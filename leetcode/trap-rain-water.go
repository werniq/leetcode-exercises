func trap(height []int) int {

	lp := 0
	rp := len(height) - 1
	l_max := 0
	r_max := 0
	result := 0
	for lp < rp {
		if height[lp] < height[rp] {
			if height[lp] > l_max {
				l_max = height[lp]
			} else {
				result += l_max - height[lp]
			}
			lp++
		} else {
			if height[rp] > r_max {
				r_max = height[rp]
			} else {
				result += r_max - height[rp]
			}
			rp--
		}
	}
	return result
}
