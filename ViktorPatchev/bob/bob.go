// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func charOnly(s string) bool {
	for _, char := range s {
		if !strings.Contains(alpha, string(char)) {
			return false
		}
	}
	return true
}

func numberOnly(s string) bool {
	for _, char := range s {
		if !strings.Contains(number, string(char)) {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var a string
	var x string = "1234567890"
	switch {

	case !numberOnly(remark) && (!strings.ContainsAny(remark, x) && remark == strings.ToUpper(remark) && remark[len(remark)-1:] == "?"):
		a = fmt.Sprintf("Calm down, I know what I'm doing!")

	case !numberOnly(remark) && remark[len(remark)-1:] == "?":
		a = fmt.Sprintf("Sure.")
	case charOnly(remark) && !numberOnly(remark) && (remark == strings.ToUpper(remark) && remark[len(remark)-1:] != "?"):
		a = fmt.Sprintf("Whoa, chill out!")

	case remark == "":
		a = fmt.Sprintf("Fine. Be that way!")
	case !charOnly(remark) && numberOnly(remark):
		a = fmt.Sprintf("Whatever.")
	default:
		a = fmt.Sprintf("Whatever.")
	}
	return a
}
