func dailyTemperatures(temperatures []int) []int {
	a := []int{} 
	sol := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(a) > 0 && temperatures[a[len(a)-1]] < v {
			topInd := len(a) - 1
			ele := a[topInd]
			a = a[0:topInd]
			sol[ele] = i - ele
		}
		a = append(a, i)
	}
	return sol
}
