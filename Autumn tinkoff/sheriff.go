package main

func sheriffMostCounts(s string) int {
	targetWord := "sheriff"
	targetFreq := make(map[rune]int)
	for _, char := range targetWord {
		targetFreq[char]++
	}

	sFreq := make(map[rune]int)
	for _, char := range s {
		sFreq[char]++
	}

	maxOccurrences := len(s)
	for char, count := range targetFreq {
		if sCount, ok := sFreq[char]; ok && sCount >= count {
			maxOccurrences = min(maxOccurrences, sCount/count)
		} else {
			maxOccurrences = 0
			break
		}
	}

	return maxOccurrences
}
