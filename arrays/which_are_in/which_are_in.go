// https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/go
package which_are_in

import (
	"sort"
	"strings"
)

func InArray(a1 []string, a2 []string) []string {
	var result []string

	for _, substr := range a1 {
		for _, str := range a2 {
			if strings.Contains(str, substr) && !containSubStr(result, substr) {
				result = append(result, substr)
			}
		}
	}

	sort.Strings(result)
	return result
}

func containSubStr(result []string, str string) bool {
	for _, substr := range result {
		if substr == str {
			return true
		}
	}
	return false
}
