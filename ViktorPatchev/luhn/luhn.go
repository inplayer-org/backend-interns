package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	var isValid bool
	var sum int
	var pom int
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), "")
	if len(input) <= 1 {
		isValid = false
	} else {
		for i := len(input) - 1; i >= 0; i-- {
			number, err := strconv.Atoi(string(input[i]))
			if err == nil {
				if pom%2 == 1 {

					number = number * 2
					if number > 9 {
						number = number - 9
					}
				}
			} else {
				return false
			}
			pom++
			sum += number

		}
		if sum%10 == 0 {
			isValid = true
		} else {
			isValid = false
		}
	}

	return isValid
}
