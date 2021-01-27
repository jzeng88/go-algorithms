// https://www.codewars.com/kata/59c5f4e9d751df43cf000035/train/go
package longest_vowel_chain

import "fmt"

func Solve(s string) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	maxCount := 0
	maxSoFar := 0

	for _, letter := range s {
		if contains(string(letter), vowels) {
			maxSoFar++
		} else {
			if maxSoFar > maxCount {
				maxCount = maxSoFar
			}
			maxSoFar = 0
		}
	}

	return maxCount
}

func contains(letter string, vowels []string) bool {
	for _, vowel := range vowels {
		fmt.Println(vowel)
		if vowel == letter {
			return true
		}
	}
	return false
}
