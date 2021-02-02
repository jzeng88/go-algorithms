// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package longest_substring

// Sliding window solution
// LR
// d   v   d   f  => {d: true}, maxSoFar = 1
// L   R
// d   v   d   f => {d: true, v: true}, maxSoFar = 2
// L       R
// d   v   d   f => {d: true, v: true}, maxSoFar = 2
// 		 L   R
// d   v   d   f => {v: true, d: true}, maxSoFar = 2
// 		 L       R
// d   v   d   f => {v: true, d: true, f: true}, maxSoFar = 3, 1(l-idx) + 3(maxSoFar) = 4 > len(string) so stop, no need to check further as any substring further is going to have a smaller length

// return 3

// O(N) time, O(N) space
func lengthOfLongestSubstring(s string) int {
	track := make(map[byte]bool)
	max, maxSoFar := 0, 0
	left, right := 0, 0

	for left < len(s) {
		if !track[s[right]] {
			track[s[right]] = true
			maxSoFar++
			right++
		} else {
			track[s[left]] = false
			maxSoFar--
			left++
		}

		if maxSoFar > max {
			max = maxSoFar
		}

		if right >= len(s) || left+max >= len(s) {
			break
		}

	}

	return max
}

// O(N) time, O(1) space
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}

// 	// 256 ASCII code char
// 	var bitSet [256]bool
// 	result, left, right := 0, 0, 0

// 	for left < len(s) {
// 		if bitSet[s[right]] {
// 			bitSet[s[left]] = false
// 			left++
// 		} else {
// 			bitSet[s[right]] = true
// 			right++
// 		}

// 		// Max so far
// 		if result < right-left {
// 			result = right - left
// 		}

// 		// If max + left pointer is greater than length of s, there will no longer be any substring that can be longer
// 		if left+result >= len(s) || right >= len(s) {
// 			break
// 		}
// 	}
// 	return result
// }
