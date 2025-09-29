package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if newWord {
				if c >= 'a' && c <= 'z' {
					runes[i] = c - 32
				}
				newWord = false
			} else {
				if c >= 'A' && c <= 'Z' {
					runes[i] = c + 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(runes)
}
