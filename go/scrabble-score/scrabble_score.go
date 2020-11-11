package scrabble

import (
	"regexp"
	"strings"
)

func Score(word string) int {
	score := 0
	for _, letter := range strings.ToUpper(word) {
		s := string(letter)
		if match(".*A|E|I|O|U|L|N|R|S|T$", s) { score += 1}
		if match(".*D|G$", s) { score += 2}
		if match(".*B|C|M|P$", s) { score += 3}
		if match(".*F|H|V|W|Y$", s) { score += 4}
		if match(".*K$", s) { score += 5}
		if match(".*J|X$", s) { score += 8}
		if match(".*Q|Z$", s) { score += 10}
	}
	return score
}

func match(pattern string, s string) bool {
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}
