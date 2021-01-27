// https://www.codewars.com/kata/59f08f89a5e129c543000069/
package string_array_duplicates

func Dup(arr []string) []string {
	var result []string

	for _, word := range arr {
		str := ""

		for _, letter := range word {
			if len(str) == 0 {
				str += string(letter)
			} else {
				if string(str[len(str)-1]) == string(letter) {
					continue
				} else {
					str += string(letter)
				}
			}
		}
		result = append(result, str)
	}

	return result
}

// Alternative without casting bytes to string
// func Dup(arr []string) []string {
//   for i, s := range arr {
//     var (
//       dupCheck byte
//       fixedStr string
//     )
//     for j, _ := range s {
//       if s[j] != dupCheck {
//         dupCheck = s[j]
//         fixedStr += string(s[j])
//       }
//     }
//     arr[i] = fixedStr
//   }
//   return arr
// }
