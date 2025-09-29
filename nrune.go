package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	index := 1
	for _, char := range s {
		if index == n {
			return char
		}
		index++
	}
	return 0
}
