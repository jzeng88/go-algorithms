// https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
package vowel_count

func GetCount(str string) (count int) {
	for _, letter := range str {
		switch string(letter) {
		case "a", "e", "i", "o", "u":
			count++
		}
	}

	return count
}
