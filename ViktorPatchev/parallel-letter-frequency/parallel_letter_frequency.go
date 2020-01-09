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
	freq := make(chan FreqMap)
	res := FreqMap{}
	for _, s := range list {
		go func(s string) {
			freq <- Frequency(s)
		}(s)
	}

	for range list {
		for letter, count := range <-freq {
			res[letter] += count
		}
	}
	return res

}
