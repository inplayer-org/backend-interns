package isogram

import "strings"

func IsIsogram(str string) bool {
	var str1 string
	var isIsogram bool = true
	var specialCharacter string = "- "
	for _, value := range str {
		if strings.ContainsAny(strings.ToLower(string(value)), str1) && !strings.ContainsAny(strings.ToLower(string(value)), specialCharacter) {
			isIsogram = false
		}
		str1 += strings.ToLower(string(value))
	}
	return isIsogram
}

//!@#$%^&*()-_=+,.<>/?;:'[{]}

// package isogram

// import "strings"

// const testVersion = 1

// var ignorables = []rune{
// 	'-',
// 	' ',
// 	'	',
// 	'+',
// 	'!',
// 	'@',
// 	'#',
// 	'$',
// 	'%',
// 	'^',
// 	'&',
// 	'*',
// 	'(',
// 	')',
// }

// func IsIsogram(candidate string) bool {
// 	hit_table := make(map[rune]int)
// runes:
// 	for _, r := range strings.ToLower(candidate) {
// 		for _, ignorable := range ignorables {
// 			if r == ignorable {
// 				continue runes
// 			}
// 		}
// 		hit_table[r] += 1
// 	}
// 	for _, count := range hit_table {
// 		if count > 1 {
// 			return false
// 		}
// 	}
// 	return true
// }
