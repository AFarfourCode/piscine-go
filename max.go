package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	ifmax := a[0]
	for _, value := range a {
		if value > ifmax {
			ifmax = value
		}
	}
	return ifmax
}
