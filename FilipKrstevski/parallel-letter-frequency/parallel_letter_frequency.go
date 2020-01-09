package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}


func ConcurrentFrequency(list []string) FreqMap {
	cha := make(chan FreqMap)
	result := FreqMap{}
	for _, text := range list {
		go func(text string) {
			cha <- Frequency(text)
		}(text)
	}
	for range list {
		for letter, count := range <-cha {
			result[letter] += count
		}
	}
	return result
}