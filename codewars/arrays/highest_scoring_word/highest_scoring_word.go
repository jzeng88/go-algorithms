// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/go
package highest_scoring_word

import "strings"

func High(s string) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	words := strings.Split(s, " ")

	var wordsSum []int

	for _, word := range words {
		letters := strings.Split(word, "")
		sum := 0

		for _, letter := range letters {
			sum += strings.Index(alphabets, letter) + 1
		}

		wordsSum = append(wordsSum, sum)
	}

	var maxSum int
	var maxSumIdx int

	for idx, wordSum := range wordsSum {
		if wordSum > maxSum {
			maxSum = wordSum
			maxSumIdx = idx
		}
	}

	return words[maxSumIdx]
}
