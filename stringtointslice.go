package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	intSlice := make([]int, 0, len(str))
	for _, charRune := range str {
		intSlice = append(intSlice, int(charRune))
	}
	return intSlice
}
