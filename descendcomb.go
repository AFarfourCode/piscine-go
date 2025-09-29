package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	output := ""
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			output += string(rune(i/10 + '0'))
			output += string(rune(i%10 + '0'))
			output += string(' ')
			output += string(rune(j/10 + '0'))
			output += string(rune(j%10 + '0'))
			if !(i == 1 && j == 0) {
				output += ", "
			}
		}
	}
	for _, r := range output {
		z01.PrintRune(r)
	}
}
