package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	lastSplitIndex := 0

	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			word := str[lastSplitIndex:i]
			summary[word]++
			lastSplitIndex = i + 1
		}
	}

	return summary
}
