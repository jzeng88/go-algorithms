// https://www.codewars.com/kata/59d9ff9f7905dfeed50000b0/train/go
package alphabet_symmetry

import "strings"

func Solve(slice []string) []int {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	result := []int{}

	for _, word := range slice {
		count := 0

		// idx = int
		// letter = int32 (ASCII code of letter)
		for idx, letter := range strings.ToLower(word) {
			// reflect.TypeOf(alphabets[idx]) = uint8
			if int32(alphabets[idx]) == letter {
				count++
			}
		}

		result = append(result, count)
	}

	return result
}
