// https://www.codewars.com/kata/586d6cefbcc21eed7a001155/train/go
package longest_consecutive_rep

func LongestRepetition(text string) Result {
	if len(text) == 0 {
		return Result{}
	}

	var maxCountSoFar, maxCount int = 1, 1
	var maxLetterSoFar, maxLetter int32 = int32(text[0]), int32(text[0])

	for _, letter := range text[1:] {
		if letter == maxLetterSoFar {
			maxCountSoFar++
		} else {
			maxLetterSoFar = letter
			maxCountSoFar = 1
		}

		if maxCountSoFar > maxCount {
			maxCount = maxCountSoFar
			maxLetter = maxLetterSoFar
		}
	}

	return Result{C: maxLetter, L: maxCount}
}
