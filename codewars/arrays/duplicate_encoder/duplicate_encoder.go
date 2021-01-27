package duplicate_encoder

import "strings"

func DuplicateEncode(word string) (result string) {
	letterCount := make(map[string]int)
	word = strings.ToLower(word)

	for _, letter := range word {
		letterCount[string(letter)] += 1
	}

	for _, letter := range word {
		if letterCount[string(letter)] > 1 {
			result += ")"
		} else {
			result += "("
		}
	}

	return result
}
