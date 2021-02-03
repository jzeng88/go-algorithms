// https://leetcode.com/problems/palindrome-number/
package palindrome_number

import "strconv"

// input = 121
// output = true

// Input: x = -121
// Output: false

// Input: x = 10
// Output: false

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)

	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
