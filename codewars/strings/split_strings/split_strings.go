// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
package split_strings

func Solution(str string) []string {
	result := []string{}

	for i := 0; i < len(str); i += 2 {
		sliced := ""

		if i+2 > len(str) {
			sliced = str[i:] + "_"
		} else {
			sliced = str[i : i+2]
		}

		result = append(result, sliced)
	}

	return result
}
