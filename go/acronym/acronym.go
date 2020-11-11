// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abbreviation := ""
	values := strings.FieldsFunc(s, split)
	for _, value := range values {
		l := []rune(value)[0]
		if unicode.IsLetter(l) {
			abbreviation += strings.ToUpper(string(l))
		}
	}
	return abbreviation
}

func split(s rune) bool {
	return s == ' ' || s == '-' || s == '_'
}
