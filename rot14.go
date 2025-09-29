package piscine

func Rot14(s string) string {
	var result string

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = (c-'a'+14)%26 + 'a'
		}
		if c >= 'A' && c <= 'Z' {
			c = (c-'A'+14)%26 + 'A'
		}
		result += string(c)
	}

	return result
}
