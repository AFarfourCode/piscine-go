package piscine

func TrimAtoi(s string) int {
	sign := 1
	foundDigit := false
	result := 0

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
			foundDigit = true
		} else if ch == '-' && !foundDigit {
			sign = -1
		}
	}

	return sign * result
}
