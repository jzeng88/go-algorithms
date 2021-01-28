// https://leetcode.com/problems/reverse-words-in-a-string/
package reverse_words_in_a_string

import "strings"

// O(n) time with 2 pass. Split string then reverse.
// O(1) space due to in place swaps
func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

// Utilize same technique as `valid palindrome`
func reverse(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
