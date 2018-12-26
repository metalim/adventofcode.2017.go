package numbers

// MinMax values of []int.
func MinMax(sn []int) (min, max int) {
	if len(sn) == 0 {
		return
	}
	min = sn[0]
	max = min
	for _, n := range sn[1:] {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
	}
	return
}
