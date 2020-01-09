package bob
import "strings"
const minLetter, maxLetter = 'a', 'z'
//Funkcija koja proveruva dali karakterot e bukva
func hasLetters(s *string) bool {
	toLower := strings.ToLower(*s)
	for _, l := range toLower {
		if l <= maxLetter && l >= minLetter {
			return true
		}
	}
	return false
}
func Hey(remark string) string {
	var response string
	remark = strings.TrimSpace(remark)
	switch {
	case (strings.HasSuffix(remark, "?") && !(strings.ToUpper(remark) == remark)) || (!hasLetters(&remark) && strings.HasSuffix(remark, "?")):
		response = "Sure."
	case strings.ToUpper(remark) == remark && !(strings.HasSuffix(remark, "?")) && hasLetters(&remark):
		response = "Whoa, chill out!"
	case strings.ToUpper(remark) == remark && strings.HasSuffix(remark, "?") && hasLetters(&remark):
		response = "Calm down, I know what I'm doing!"
	case len(remark) == 0:
		response = "Fine. Be that way!"
	default:
		response = "Whatever."
	}
	return response
}