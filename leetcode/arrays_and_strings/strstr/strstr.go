// https://leetcode.com/problems/implement-strstr/
package strstr

import "strings"

// O(n * m) time complexity where n is the length of `haystack` & m is the length of `needle`
// O(1) space
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			// End of needle
			if j == len(needle) {
				return i
			}

			// End of haystack
			if i+j == len(haystack) {
				return -1
			}

			// skip current letter in haystack if it doesn't match the needle
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// Shortcut approach:
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
