package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	var result []string

	for _, v := range original {
		if v != "" {
			result = append(result, v)
		}
	}

	*ptr = result
	return len(result)
}
