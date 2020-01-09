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
	str := []rune(s)
	var specialChars = "-_"
	//Couldn't find unicode function to check for special character
	var newString string
	for i, value := range s {
		if i < len(s)-1 {
			if unicode.IsSpace(value) || strings.ContainsAny(string(s[i]), specialChars) {
				if unicode.IsLetter(str[i+1]) {
					newString += strings.ToUpper(string(s[i+1]))
				}
			} else if i == 0 {
				newString += strings.ToUpper(string(s[i]))
			}
		}
	}

	return newString
}
