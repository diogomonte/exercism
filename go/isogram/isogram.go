package isogram

import (
	"unicode"
)

func IsIsogram(input string) bool {
	counter := make(map[rune]int)
	for _, value := range input {
		lowerValue := unicode.ToLower(value)

		count, found := counter[lowerValue]

		if found && unicode.IsLetter(value) {
			counter[lowerValue] = count + 1
		} else {
			counter[lowerValue] = 1
		}

		if counter[lowerValue] > 1 {
			return false
		}
	}
	return true
}