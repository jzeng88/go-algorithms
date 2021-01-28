// https://leetcode.com/problems/valid-palindrome/description/
package valid_palindrome

import (
	"strings"
)

// Two pointers approach:
// O(n) time, O(1) space
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1

	for i < j {
		// Move left pointer right if non-alphanumeric
		for i < j && !isChar(s[i]) {
			i++
		}

		// Move right pointer left if non-alphanumeric
		for i < j && !isChar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}
	return true
}

// Check if non-alphanumeric character
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
